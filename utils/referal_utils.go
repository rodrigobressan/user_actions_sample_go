package utils

import (
	"surfe_assignment/models"
)

const (
	ActionReferUserAlias = "REFER_USER"
)

// ComputeReferralIndex calculates the referral index for all users,
// including those who have not referred anyone (index 0)
func ComputeReferralIndex(actions []models.Action) map[int]int {
	// Step 1: Build the referral graphUserReferrals (userID -> list of referred users)
	graphUserReferrals := make(map[int][]int)

	// Track all unique user IDs seen (referrers and referees)
	allUsers := make(map[int]bool)

	for _, action := range actions {
		if action.Type == ActionReferUserAlias {
			referrer := action.UserID
			referred := *action.TargetUser

			graphUserReferrals[referrer] = append(graphUserReferrals[referrer], referred)

			allUsers[referrer] = true
			allUsers[referred] = true
		} else {
			// If it's not REFER_USER, still collect the userID for completeness
			allUsers[action.UserID] = true
		}
	}

	// Step 2: Prepare memoization map to cache referral index per user
	memoUserTotalReferrals := make(map[int]int)

	// Step 3: Define recursive DFS to compute referral index for a user
	var computeReferralIndexDfs func(userID int) int
	computeReferralIndexDfs = func(userID int) int {
		if val, ok := memoUserTotalReferrals[userID]; ok {
			return val
		}

		total := 0
		for _, referredUser := range graphUserReferrals[userID] {
			total += 1 + computeReferralIndexDfs(referredUser)
		}

		memoUserTotalReferrals[userID] = total
		return total
	}

	// Step 4: Compute referral index for all users we encountered
	result := make(map[int]int)
	for userID := range allUsers {
		result[userID] = computeReferralIndexDfs(userID)
	}

	return result
}
