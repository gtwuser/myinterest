package main

import (
	"fmt"
	"regexp"
	"strings"
)

type Value struct {
	Number float64
}


//var s string = `right_to_customer_portal_access_view_information
//
//right_to_administrate_asset_groups_manage_asset_groups
//
//right_to_manage_users_groups
//
//right_to_manage_access_policies
//
//right_to_access_management_assign_roles
//
//right_to_administrate_account_manage_users
//
//right_to_automated_fault_management_opencase_policy
//
//right_to_rapid_problem_resolution_opt_in_out_policy
//
//right_to_compliance_opt_in_out_policy
//
//right_to_services_upgrade_request_upgrade
//
//right_to_services_renewal_request_renewal
//
//right_to_collector_scan_collector_set_up
//
//right_to_collector_scan_schedule_scan
//
//right_to_collector_scan_run_scan
//
//right_to_collector_scan_view_schedule
//
//right_to_manage_pitstop_complete_checklist
//
//right_to_public_communities_view_threads
//
//right_to_public_communities_initiate_thread
//
//right_to_success_track_communities_initiate_thread
//
//right_to_success_track_communities_attend_q_a_session
//
//right_to_success_tips_view_success_tips
//
//right_to_atx_register_atx
//
//right_to_atx_attend_atx
//
//right_to_atx_feedback_atx
//
//right_to_atx_bookmark_atx
//
//right_to_atx_view_on_demand_atx_recordings
//
//right_to_accelerator_request_accelerator
//
//right_to_accelerator_feedback_accelerator
//
//right_to_accelerator_bookmark_accelerator
//
//right_to_customized_group_training_request_training
//
//right_to_elearning_view_information
//
//right_to_certification_prep_remote_labs_view_information
//
//right_to_administrate_asset_groups_assign_asset_to_asset_groups
//
//right_to_view_advisories_sec_adv_field_notices_priority_bugs_
//
//right_to_view_contract_details
//
//right_to_view_licence_details
//
//right_to_run_ad_hoc_diagnostic_scans
//
//right_to_ignore_advisories
//
//right_to_technical_support_view_case
//
//right_to_technical_support_open_case
//
//right_to_hardware_return_open_rma
//
//right_to_hardware_return_view_rma
//
//right_to_view_software_recommendations
//
//right_to_accept_cancel_optimal_software
//
//right_to_change_the_recommendations_interval
//
//right_to_view_crash_risk_and_crashed_score_cards_and_corresponding_grids
//
//right_to_view_similar_and_compare_devices_from_crash_risk_grid
//
//right_to_view_details_from_crashed_devices_grid
//
//right_to_ignore_unignore_devices
//
//right_to_opt_in_opt_out_to_regulatory_compliance
//
//right_to_create_update_policy_profiles
//
//right_to_view_compliance_violations_and_systems_w_violations_pages
//
//right_to_view_violation_details_from_violations_grid
//
//right_to_view_systems_details_from_systems_w_violations_grid
//
//right_to_view_policy_profiles
//
//right_to_view_on_demand_compliance_check_results
//
//right_to_enable_disable_ic_in_fault_catalog
//
//right_to_enable_disable_syslog_ic
//
//right_to_view_fault_and_syslog_dashboard_information
//
//right_to_edit_cco_id_for_tac_case_creation
//
//right_to_create_edit_schedule_maintenance_window
//
//right_to_ignore_unignore_events`
func main() {
	//fields := strings.Fields(s)
	//for _, v := range fields{
	//	if strings.HasSuffix(v, "_"){
	//
	//	}
	//}
	//fmt.Println(fields)// Lots of slashes both forward and back ward
	underscore := replaceAllWithUnderscore(`kapil/////joshi\////kamal/joshi//right_to_view_advisories_sec_adv_field_notices_priority_bugs_
`)//this is like the policy names input
	fmt.Println(underscore)
}

func replaceAllWithUnderscore(va string) string {
	ls := strings.ToLower(va)
	var re = regexp.MustCompile(`[^A-Za-z0-9]`)
	fs := re.ReplaceAllString(ls, `_`)
	re = regexp.MustCompile(`_+`)
	fs = re.ReplaceAllString(fs, `_`)
	return fs
}
