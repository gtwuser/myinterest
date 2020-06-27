package cxc.customer.portal

    default right_to_customer_portal_access_view_information = false
    right_to_customer_portal_access_view_information {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customer_portal_access_lifecycle_view_information = false
    right_to_customer_portal_access_lifecycle_view_information {
    	input.method == GET
    	portal_account.roles[_] == [ ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customer_portal_access_assets_view_information = false
    right_to_customer_portal_access_assets_view_information {
    	input.method == GET
    	portal_account.roles[_] == [ ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customer_portal_access_advisories_view_information = false
    right_to_customer_portal_access_advisories_view_information {
    	input.method == GET
    	portal_account.roles[_] == [ ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customer_portal_access_cases_view_information = false
    right_to_customer_portal_access_cases_view_information {
    	input.method == GET
    	portal_account.roles[_] == [ ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customer_portal_access_insights_view_information = false
    right_to_customer_portal_access_insights_view_information {
    	input.method == GET
    	portal_account.roles[_] == [ ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_administrate_asset_groups_manage_asset_groups = false
    right_to_administrate_asset_groups_manage_asset_groups {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_users_groups = false
    right_to_manage_users_groups {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_access_policies = false
    right_to_manage_access_policies {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_have_resource_access_assign_smart_account = false
    right_to_have_resource_access_assign_smart_account {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_access_management_assign_roles = false
    right_to_access_management_assign_roles {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_administrate_account_manage_users = false
    right_to_administrate_account_manage_users {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_automated_fault_management_opencase_policy = false
    right_to_automated_fault_management_opencase_policy {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_rapid_problem_resolution_opt_in_out_policy = false
    right_to_rapid_problem_resolution_opt_in_out_policy {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_compliance_opt_in_out_policy = false
    right_to_compliance_opt_in_out_policy {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_services_upgrade_request_upgrade = false
    right_to_services_upgrade_request_upgrade {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_services_renewal_request_renewal = false
    right_to_services_renewal_request_renewal {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_collector_scan_collector_set_up = false
    right_to_collector_scan_collector_set_up {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_collector_scan_schedule_scan = false
    right_to_collector_scan_schedule_scan {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_collector_scan_run_scan = false
    right_to_collector_scan_run_scan {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_collector_scan_view_schedule = false
    right_to_collector_scan_view_schedule {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_pitstop_complete_checklist = false
    right_to_manage_pitstop_complete_checklist {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_communities_initiate_thread = false
    right_to_communities_initiate_thread {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_success_track_communities_initiate_thread = false
    right_to_success_track_communities_initiate_thread {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_atx_register_atx = false
    right_to_atx_register_atx {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_atx_attend_atx = false
    right_to_atx_attend_atx {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_accelerator_request_accelerator = false
    right_to_accelerator_request_accelerator {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_customized_group_training_request_training = false
    right_to_customized_group_training_request_training {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_elearning_view_information = false
    right_to_elearning_view_information {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_certification_prep_remote_labs_view_information = false
    right_to_certification_prep_remote_labs_view_information {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_service_manage_service_entitlement = false
    right_to_manage_service_manage_service_entitlement {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_device_manage_device = false
    right_to_manage_device_manage_device {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_assets_from_dnac = false
    right_to_view_assets_from_dnac {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_purchased_assets = false
    right_to_view_purchased_assets {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_advisories_sec_adv_field_notices_priority_bugs = false
    right_to_view_advisories_sec_adv_field_notices_priority_bugs {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_contract_details = false
    right_to_view_contract_details {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_licence_details = false
    right_to_view_licence_details {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_run_ad_hoc_diagnostic_scans = false
    right_to_run_ad_hoc_diagnostic_scans {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_edit_create_tags = false
    right_to_edit_create_tags {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_use_assign_reassign_tags = false
    right_to_use_assign_reassign_tags {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_ignore_advisories = false
    right_to_ignore_advisories {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_license_manage_license_entitlement = false
    right_to_manage_license_manage_license_entitlement {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_license_move_license_entitlement = false
    right_to_manage_license_move_license_entitlement {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_manage_license_use_license = false
    right_to_manage_license_use_license {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_software_version_upgrade_upgrade_license_entitlement = false
    right_to_software_version_upgrade_upgrade_license_entitlement {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_technical_support_view_case = false
    right_to_technical_support_view_case {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_technical_support_open_case = false
    right_to_technical_support_open_case {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_hardware_return_open_rma = false
    right_to_hardware_return_open_rma {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_hardware_return_view_rma = false
    right_to_hardware_return_view_rma {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_software_download_download_software_image = false
    right_to_software_download_download_software_image {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_order_order_download = false
    right_to_order_order_download {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_software_recommendations = false
    right_to_view_software_recommendations {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_accept_cancel_optimal_software = false
    right_to_accept_cancel_optimal_software {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_mark_optimal_software_as_golden = false
    right_to_mark_optimal_software_as_golden {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_change_the_recommendations_interval = false
    right_to_change_the_recommendations_interval {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_crash_risk_and_crashed_score_cards_and_corresponding_grids = false
    right_to_view_crash_risk_and_crashed_score_cards_and_corresponding_grids {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_similar_and_compare_devices_from_crash_risk_grid = false
    right_to_view_similar_and_compare_devices_from_crash_risk_grid {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_details_from_crashed_devices_grid = false
    right_to_view_details_from_crashed_devices_grid {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_ignore_unignore_devices = false
    right_to_ignore_unignore_devices {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_opt_in_opt_out_to_regulatory_compliance = false
    right_to_opt_in_opt_out_to_regulatory_compliance {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_create_update_policy_profiles = false
    right_to_create_update_policy_profiles {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_create_update_cancel_a_violation_waiver = false
    right_to_create_update_cancel_a_violation_waiver {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_initiate_an_on_demand_compliance_check = false
    right_to_initiate_an_on_demand_compliance_check {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","internalFullUser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_compliance_violations_and_systems_w_violations_pages = false
    right_to_view_compliance_violations_and_systems_w_violations_pages {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_violation_details_from_violations_grid = false
    right_to_view_violation_details_from_violations_grid {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_systems_details_from_systems_w_violations_grid = false
    right_to_view_systems_details_from_systems_w_violations_grid {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_policy_profiles = false
    right_to_view_policy_profiles {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_on_demand_compliance_check_results = false
    right_to_view_on_demand_compliance_check_results {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_enable_disable_ic_in_fault_catalog = false
    right_to_enable_disable_ic_in_fault_catalog {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_enable_disable_syslog_ic = false
    right_to_enable_disable_syslog_ic {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_view_fault_and_syslog_dashboard_information = false
    right_to_view_fault_and_syslog_dashboard_information {
    	input.method == GET
    	portal_account.roles[_] == ["superAdmin","admin","fulluser","viewer","internalFullUser", "internalviewer"]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_edit_cco_id_for_tac_case_creation = false
    right_to_edit_cco_id_for_tac_case_creation {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_create_edit_schedule_maintenance_window = false
    right_to_create_edit_schedule_maintenance_window {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_ignore_unignore_events = false
    right_to_ignore_unignore_events {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin","fulluser", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }

    default right_to_on_demand_expert_request_on_demand_expert = false
    right_to_on_demand_expert_request_on_demand_expert {
    	input.method == ["GET", "POST", "UPDATE", "DELETE"]
    	portal_account.roles[_] == ["superAdmin","admin", ]
    	[header, payload, sig] := io.jwt.decode(input.token)
    }


