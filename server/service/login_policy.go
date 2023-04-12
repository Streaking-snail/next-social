package service

import (
	"context"
	"errors"
	"net"
	"next-social/server/model"
	"next-social/server/repository"
	"next-social/server/utils"
	"strings"
	"time"
)

var LoginPolicyService = new(loginPolicyService)

type loginPolicyService struct {
	baseService
}

func (s loginPolicyService) Check(userId, clientIp string) error {
	ctx := context.Background()
	// 按照优先级倒排进行查询
	policies, err := repository.LoginPolicyRepository.FindByUserId(ctx, userId)
	if err != nil {
		return err
	}
	if len(policies) == 0 {
		return nil
	}

	if err := s.checkClientIp(policies, clientIp); err != nil {
		return err
	}

	if err := s.checkWeekDay(policies); err != nil {
		return err
	}
	return nil
}

func (s loginPolicyService) checkClientIp(policies []model.LoginPolicy, clientIp string) error {
	var pass = true
	// 优先级低的先进行判断
	for _, policy := range policies {
		if !policy.Enabled {
			continue
		}
		ipGroups := strings.Split(policy.IpGroup, ",")
		for _, group := range ipGroups {
			if strings.Contains(group, "/") {
				// CIDR
				_, ipNet, err := net.ParseCIDR(group)
				if err != nil {
					continue
				}
				if !ipNet.Contains(net.ParseIP(clientIp)) {
					continue
				}
			} else if strings.Contains(group, "-") {
				// 范围段
				split := strings.Split(group, "-")
				if len(split) < 2 {
					continue
				}
				start := split[0]
				end := split[1]
				intReqIP := utils.IpToInt(clientIp)
				if intReqIP < utils.IpToInt(start) || intReqIP > utils.IpToInt(end) {
					continue
				}
			} else {
				// IP
				if group != clientIp {
					continue
				}
			}
			pass = policy.Rule == "allow"
		}
	}

	if !pass {
		return errors.New("非常抱歉，您当前使用的IP地址不允许进行登录。")
	}
	return nil
}

func (s loginPolicyService) checkWeekDay(policies []model.LoginPolicy) error {
	// 获取当前日期是星期几
	now := time.Now()
	weekday := int(now.Weekday())
	hwc := now.Format("15:04")

	var timePass = true

	// 优先级低的先进行判断
	for _, policy := range policies {
		if !policy.Enabled {
			continue
		}
		timePeriods, err := repository.TimePeriodRepository.FindByLoginPolicyId(context.Background(), policy.ID)
		if err != nil {
			return err
		}

		for _, period := range timePeriods {
			if weekday != period.Key {
				continue
			}
			if period.Value == "" {
				continue
			}
			// 只处理对应天的数据
			times := strings.Split(period.Value, "、")
			for _, t := range times {
				timeRange := strings.Split(t, "~")
				start := timeRange[0]
				end := timeRange[1]
				if (start == "00:00" && end == "00:00") || (start <= hwc && hwc <= end) {
					timePass = policy.Rule == "allow"
				}
			}
		}
	}

	if !timePass {
		return errors.New("非常抱歉，当前时段不允许您进行登录。")
	}

	return nil
}
