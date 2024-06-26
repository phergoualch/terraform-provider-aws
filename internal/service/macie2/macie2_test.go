// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package macie2_test

import (
	"testing"

	"github.com/hashicorp/terraform-provider-aws/internal/acctest"
	"github.com/hashicorp/terraform-provider-aws/names"
)

func TestAccMacie2_serial(t *testing.T) {
	t.Parallel()

	testCases := map[string]map[string]func(t *testing.T){
		"Account": {
			"basic":                        testAccAccount_basic,
			"finding_publishing_frequency": testAccAccount_FindingPublishingFrequency,
			names.AttrStatus:               testAccAccount_WithStatus,
			"finding_and_status":           testAccAccount_WithFindingAndStatus,
			"disappears":                   testAccAccount_disappears,
		},
		"ClassificationExportConfiguration": {
			"basic": testAccClassificationExportConfiguration_basic,
		},
		"ClassificationJob": {
			"basic":              testAccClassificationJob_basic,
			"name_generated":     testAccClassificationJob_Name_Generated,
			names.AttrNamePrefix: testAccClassificationJob_NamePrefix,
			"disappears":         testAccClassificationJob_disappears,
			names.AttrStatus:     testAccClassificationJob_Status,
			"complete":           testAccClassificationJob_complete,
			names.AttrTags:       testAccClassificationJob_WithTags,
			"bucket_criteria":    testAccClassificationJob_BucketCriteria,
		},
		"CustomDataIdentifier": {
			"basic":              testAccCustomDataIdentifier_basic,
			"name_generated":     testAccCustomDataIdentifier_Name_Generated,
			names.AttrNamePrefix: testAccCustomDataIdentifier_disappears,
			"disappears":         testAccCustomDataIdentifier_NamePrefix,
			"classification_job": testAccCustomDataIdentifier_WithClassificationJob,
			names.AttrTags:       testAccCustomDataIdentifier_WithTags,
		},
		"FindingsFilter": {
			"basic":              testAccFindingsFilter_basic,
			"name_generated":     testAccFindingsFilter_Name_Generated,
			names.AttrNamePrefix: testAccFindingsFilter_NamePrefix,
			"disappears":         testAccFindingsFilter_disappears,
			"complete":           testAccFindingsFilter_complete,
			"date":               testAccFindingsFilter_WithDate,
			"number":             testAccFindingsFilter_WithNumber,
			names.AttrTags:       testAccFindingsFilter_withTags,
		},
		"OrganizationAdminAccount": {
			"basic":      testAccOrganizationAdminAccount_basic,
			"disappears": testAccOrganizationAdminAccount_disappears,
		},
		"Member": {
			"basic":                                 testAccMember_basic,
			"disappears":                            testAccMember_disappears,
			names.AttrTags:                          testAccMember_withTags,
			"invitation_disable_email_notification": testAccMember_invitationDisableEmailNotification,
			"invite":                                testAccMember_invite,
			"invite_removed":                        testAccMember_inviteRemoved,
			names.AttrStatus:                        testAccMember_status,
		},
		"InvitationAccepter": {
			"basic": testAccInvitationAccepter_basic,
		},
	}

	acctest.RunSerialTests2Levels(t, testCases, 0)
}
