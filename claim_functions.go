package main

import (
	"database/sql"
	"github.com/golang/protobuf/proto"
	"github.com/vamitrou/pia-oracle/protobuf"
)

func ExecuteScoreInsert(stmt *sql.Stmt, score map[string]interface{}) error {
	if _, ok := score["SCORE"].(string); ok {
		_, err := stmt.Exec(score["GLB_OE_ID"].(float64),
			score["CLM_BUS_ID"].(string),
			nil,
			nil,
			score["CREATE_DT_TS"].(float64))
		return err
	}
	_, err := stmt.Exec(score["GLB_OE_ID"].(float64),
		score["CLM_BUS_ID"].(string),
		score["SCORE"].(float64),
		score["MODEL"].(string),
		score["CREATE_DT_TS"].(float64))
	return err
}

func ExecuteImpInsert(stmt *sql.Stmt, var_imp map[string]interface{}) error {
	var err error
	if _, ok := var_imp["pred"].(float64); ok {
		_, err = stmt.Exec(var_imp["GLB_OE_ID"].(float64),
			var_imp["CLM_BUS_ID"].(string),
			var_imp["MODEL"].(string),
			var_imp["MODEL_DESC"].(string),
			var_imp["VAR"].(string),
			var_imp["VALUE"].(string),
			var_imp["WEIGHT"].(string),
			var_imp["RANK"].(float64),
			var_imp["pred"].(float64),
			var_imp["CREATE_DT_TS"].(float64))
	} else {
		_, err = stmt.Exec(var_imp["GLB_OE_ID"].(float64),
			var_imp["CLM_BUS_ID"].(string),
			var_imp["MODEL"].(string),
			var_imp["MODEL_DESC"].(string),
			var_imp["VAR"].(string),
			var_imp["VALUE"].(string),
			nil,
			nil,
			nil,
			var_imp["CREATE_DT_TS"].(float64))
	}
	check(err)
	return err
}

func _ExecuteInsert(stmt *sql.Stmt, score *protoclaim.ProtoListScore_ProtoScore) {
	_, err := stmt.Exec(score.GetGLB_OE_ID(), score.GetCLM_BUS_ID(), score.GetSCORE(),
		score.GetMODEL(), score.GetCREATE_DT(), score.GetCREATE_DT_TS())
	check(err)
}

func ClaimForMap(cmap map[string]interface{}) *protoclaim.ProtoListClaim_ProtoClaim {
	claim := &protoclaim.ProtoListClaim_ProtoClaim{
		GLB_OE_ID:                   proto.Float64(AssertFloat(cmap["GLB_OE_ID"])),
		CLM_RK:                      proto.Float64(AssertFloat(cmap["CLM_RK"])),
		CLM_BUS_ID:                  proto.String(AssertString(cmap["CLM_BUS_ID"])),
		ABT_DT_ZERO:                 proto.Int64(AssertTime(cmap["ABT_DT_ZERO"])),
		CLM_STTS_CD:                 proto.String(AssertString(cmap["CLM_STTS_CD"])),
		CLM_TYP_CD:                  proto.String(AssertString(cmap["CLM_TYP_CD"])),
		CLM_CVRG_TYP:                proto.String(AssertString(cmap["CLM_CVRG_TYP"])),
		CLM_REASON_CD:               proto.String(AssertString(cmap["CLM_REASON_CD"])),
		CLM_DMG_DTL_CD:              proto.String(AssertString(cmap["CLM_DMG_DTL_CD"])),
		CLM_INDMNTY_AMT:             proto.Float64(AssertFloat(cmap["CLM_INDMNTY_AMT"])),
		CLM_RSRV_AMT:                proto.Float64(AssertFloat(cmap["CLM_RSRV_AMT"])),
		CLM_INITL_RSRV_AMT:          proto.Float64(AssertFloat(cmap["CLM_INITL_RSRV_AMT"])),
		DIFF_INTL_RSRV_PYMNT:        proto.Float64(AssertFloat(cmap["DIFF_INTL_RSRV_PYMNT"])),
		CLM_FEE_AMT:                 proto.Float64(AssertFloat(cmap["CLM_FEE_AMT"])),
		TOT_CLM_AMT:                 proto.Float64(AssertFloat(cmap["TOT_CLM_AMT"])),
		INCDNT_DTTM:                 proto.Int64(AssertTime(cmap["INCDNT_DTTM"])),
		INCDNT_HOUR_OF_DAY:          proto.Float64(AssertFloat(cmap["INCDNT_HOUR_OF_DAY"])),
		INCDNT_DAY_OF_WEEK:          proto.Float64(AssertFloat(cmap["INCDNT_DAY_OF_WEEK"])),
		CLM_RGSTR_DTTM:              proto.Int64(AssertTime(cmap["CLM_RGSTR_DTTM"])),
		DAYS_INCDNT_RGSTR:           proto.Float64(AssertFloat(cmap["DAYS_INCDNT_RGSTR"])),
		CLM_LAST_CLSD_DTTM:          proto.Int64(AssertTime(cmap["CLM_LAST_CLSD_DTTM"])),
		CLM_LAST_STTLM_DTTM:         proto.Int64(AssertTime(cmap["CLM_LAST_STTLM_DTTM"])),
		CLM_LGL_FLG:                 proto.Float64(AssertFloat(cmap["CLM_LGL_FLG"])),
		CLM_NOTIF_CHNL_CD:           proto.String(AssertString(cmap["CLM_NOTIF_CHNL_CD"])),
		CLM_NOTIF_SRC_CD:            proto.String(AssertString(cmap["CLM_NOTIF_SRC_CD"])),
		CLM_NOTIF_TO_CLM_HNDL_DTTM:  proto.Int64(AssertTime(cmap["CLM_NOTIF_TO_CLM_HNDL_DTTM"])),
		DAYS_NOTIF_CHNL_HNDL:        proto.Float64(AssertFloat(cmap["DAYS_NOTIF_CHNL_HNDL"])),
		CLM_RCVR_PAID_AMT:           proto.Float64(AssertFloat(cmap["CLM_RCVR_PAID_AMT"])),
		CLM_RSRV_STTLM_TRTY_AMT:     proto.Float64(AssertFloat(cmap["CLM_RSRV_STTLM_TRTY_AMT"])),
		INCDNT_DESC:                 proto.String(AssertString(cmap["INCDNT_DESC"])),
		INCDNT_LOC_DESC:             proto.String(AssertString(cmap["INCDNT_LOC_DESC"])),
		INCDNT_LOC_GPS_CRDNT:        proto.String(AssertString(cmap["INCDNT_LOC_GPS_CRDNT"])),
		INCDNT_MULTI_VEH_INCDNT_FLG: proto.Float64(AssertFloat(cmap["INCDNT_MULTI_VEH_INCDNT_FLG"])),
		INCDNT_OTH_OBJ_FLG:          proto.Float64(AssertFloat(cmap["INCDNT_OTH_OBJ_FLG"])),
		INCDNT_PLCE_AT_SCENE_FLG:    proto.Float64(AssertFloat(cmap["INCDNT_PLCE_AT_SCENE_FLG"])),
		INCDNT_TOTAL_VEH_NO:         proto.Float64(AssertFloat(cmap["INCDNT_TOTAL_VEH_NO"])),
		INCDNT_CNTRY_CD:             proto.String(AssertString(cmap["INCDNT_CNTRY_CD"])),
		INCDNT_ABROAD_FLG:           proto.Float64(AssertFloat(cmap["INCDNT_ABROAD_FLG"])),
		INCDNT_ADDR_LN_1:            proto.String(AssertString(cmap["INCDNT_ADDR_LN_1"])),
		INCDNT_CITY_NM:              proto.String(AssertString(cmap["INCDNT_CITY_NM"])),
		INCDNT_POSTAL_CD:            proto.String(AssertString(cmap["INCDNT_POSTAL_CD"])),
		INCDNT_DISTRICT_NM:          proto.String(AssertString(cmap["INCDNT_DISTRICT_NM"])),
		CNTRCT_BUS_ID:               proto.String(AssertString(cmap["CNTRCT_BUS_ID"])),
		CNTRCT_ISSUE_DTTM:           proto.Int64(AssertTime(cmap["CNTRCT_ISSUE_DTTM"])),
		DAYS_INCDNT_CNTRCT_ISSUE:    proto.Float64(AssertFloat(cmap["DAYS_INCDNT_CNTRCT_ISSUE"])),
		DAYS_RGSTR_CNTRCT_ISSUE:     proto.Float64(AssertFloat(cmap["DAYS_RGSTR_CNTRCT_ISSUE"])),
		CNTRCT_NEXT_RNWL_DTTM:       proto.Int64(AssertTime(cmap["CNTRCT_NEXT_RNWL_DTTM"])),
		CNTRCT_EXPRY_DTTM:           proto.Int64(AssertTime(cmap["CNTRCT_EXPRY_DTTM"])),
		DAYS_INCDNT_CNTRCT_EXPRY:    proto.Float64(AssertFloat(cmap["DAYS_INCDNT_CNTRCT_EXPRY"])),
		PYMT_METHOD_CD:              proto.String(AssertString(cmap["PYMT_METHOD_CD"])),
		PYMT_MODE_CD:                proto.String(AssertString(cmap["PYMT_MODE_CD"])),
		GWP_ANN_AMT:                 proto.Float64(AssertFloat(cmap["GWP_ANN_AMT"])),
		TSI_AMT:                     proto.Float64(AssertFloat(cmap["TSI_AMT"])),
		RATIO_CLM_AMT_TSI_AMT:       proto.Float64(AssertFloat(cmap["RATIO_CLM_AMT_TSI_AMT"])),
		PLCY_BUS_ID:                 proto.String(AssertString(cmap["PLCY_BUS_ID"])),
		PLCY_CNTRCT_CTGRY_CD:        proto.String(AssertString(cmap["PLCY_CNTRCT_CTGRY_CD"])),
		PLCY_CNTRCT_STTS_CD:         proto.String(AssertString(cmap["PLCY_CNTRCT_STTS_CD"])),
		PLCY_CNTRCT_TYP_CD:          proto.String(AssertString(cmap["PLCY_CNTRCT_TYP_CD"])),
		PLCY_PRM_ACCT_STTS_CD:       proto.String(AssertString(cmap["PLCY_PRM_ACCT_STTS_CD"])),
		PLCY_MNTH_PYMT_DUE_DAY_NO:   proto.Float64(AssertFloat(cmap["PLCY_MNTH_PYMT_DUE_DAY_NO"])),
		PLCY_PRM_CVR_FLG:            proto.Float64(AssertFloat(cmap["PLCY_PRM_CVR_FLG"])),
		PLCY_PRM_CVR_FLG_UPDT_DT:    proto.Int64(AssertTime(cmap["PLCY_PRM_CVR_FLG_UPDT_DT"])),
		PLCY_PRM_PAID_UNTIL_DT:      proto.Int64(AssertTime(cmap["PLCY_PRM_PAID_UNTIL_DT"])),
		PLCY_ISSUE_DTTM:             proto.Int64(AssertTime(cmap["PLCY_ISSUE_DTTM"])),
		DAYS_INCDNT_PLCY_ISSUE:      proto.Float64(AssertFloat(cmap["DAYS_INCDNT_PLCY_ISSUE"])),
		DAYS_RGSTR_PLCY_ISSUE:       proto.Float64(AssertFloat(cmap["DAYS_RGSTR_PLCY_ISSUE"])),
		PLCY_NEXT_RNWL_DTTM:         proto.Int64(AssertTime(cmap["PLCY_NEXT_RNWL_DTTM"])),
		PLCY_EXPRY_DTTM:             proto.Int64(AssertTime(cmap["PLCY_EXPRY_DTTM"])),
		DAYS_INCDNT_PLCY_EXPRY:      proto.Float64(AssertFloat(cmap["DAYS_INCDNT_PLCY_EXPRY"])),
		CLM_PLCY_HLD_RSPN_PCT:       proto.Float64(AssertFloat(cmap["CLM_PLCY_HLD_RSPN_PCT"])),
		FRAUD_FLG:                   proto.String(AssertString(cmap["FRAUD_FLG"])),
		INVSTGT_FLG:                 proto.Float64(AssertFloat(cmap["INVSTGT_FLG"])),
		INVSTGT_STRT_DT:             proto.Int64(AssertTime(cmap["INVSTGT_STRT_DT"])),
		INVSTGT_END_DT:              proto.Int64(AssertTime(cmap["INVSTGT_END_DT"])),
		BI_FLG:                      proto.Float64(AssertFloat(cmap["BI_FLG"])),
		OBJ_TTL_LOSS_FLG:            proto.Float64(AssertFloat(cmap["OBJ_TTL_LOSS_FLG"])),
		CLM_OBJ_ROLE_CD:             proto.String(AssertString(cmap["CLM_OBJ_ROLE_CD"])),
		OBJ_TYP_CD:                  proto.String(AssertString(cmap["OBJ_TYP_CD"])),
		VEH_TYP_CD:                  proto.String(AssertString(cmap["VEH_TYP_CD"])),
		VEH_USE_TYP_CD:              proto.String(AssertString(cmap["VEH_USE_TYP_CD"])),
		OBJ_BUS_ID:                  proto.String(AssertString(cmap["OBJ_BUS_ID"])),
		VEH_VIN:                     proto.String(AssertString(cmap["VEH_VIN"])),
		VEH_MNFCTR_DT:               proto.String(AssertString(cmap["VEH_MNFCTR_DT"])),
		VEH_AGE_YRS:                 proto.Float64(AssertFloat(cmap["VEH_AGE_YRS"])),
		VEH_COLOR_TXT:               proto.String(AssertString(cmap["VEH_COLOR_TXT"])),
		VEH_MNFCTR_NM:               proto.String(AssertString(cmap["VEH_MNFCTR_NM"])),
		VEH_MODEL_TXT:               proto.String(AssertString(cmap["VEH_MODEL_TXT"])),
		VEH_MODEL_YR_NO:             proto.String(AssertString(cmap["VEH_MODEL_YR_NO"])),
		VEH_PRCH_AMT:                proto.Float64(AssertFloat(cmap["VEH_PRCH_AMT"])),
		VEH_LEASED_FLG:              proto.Float64(AssertFloat(cmap["VEH_LEASED_FLG"])),
		VEH_LCNS_PLATE_NO_TXT:       proto.String(AssertString(cmap["VEH_LCNS_PLATE_NO_TXT"])),
		VEH_PRCH_DT:                 proto.Int64(AssertTime(cmap["VEH_PRCH_DT"])),
		RATIO_CLM_AMT_VEH_PRCH_AMT:  proto.Float64(AssertFloat(cmap["RATIO_CLM_AMT_VEH_PRCH_AMT"])),
		TP_VEH_AGE_YRS:              proto.Float64(AssertFloat(cmap["TP_VEH_AGE_YRS"])),
		PLCE_RPRT_FLG:               proto.Float64(AssertFloat(cmap["PLCE_RPRT_FLG"])),
		PRTY_TYP_CD:                 proto.String(AssertString(cmap["PRTY_TYP_CD"])),
		LE_BUS_TYP_CD:               proto.String(AssertString(cmap["LE_BUS_TYP_CD"])),
		PE_GENDER_TYP_CD:            proto.String(AssertString(cmap["PE_GENDER_TYP_CD"])),
		PE_MRTL_STTS_CD:             proto.String(AssertString(cmap["PE_MRTL_STTS_CD"])),
		PE_OCC_TYP_CD:               proto.String(AssertString(cmap["PE_OCC_TYP_CD"])),
		PRTY_TAX_NO_TXT:             proto.String(AssertString(cmap["PRTY_TAX_NO_TXT"])),
		PRTY_WATCH_LIST_FLG:         proto.Float64(AssertFloat(cmap["PRTY_WATCH_LIST_FLG"])),
		PLCY_HLD_NM:                 proto.String(AssertString(cmap["PLCY_HLD_NM"])),
		PE_BIRTH_DT:                 proto.Int64(AssertTime(cmap["PE_BIRTH_DT"])),
		PLCY_HLD_AGE:                proto.Float64(AssertFloat(cmap["PLCY_HLD_AGE"])),
		PRTY_CNTRY_CD:               proto.String(AssertString(cmap["PRTY_CNTRY_CD"])),
		PRTY_ADDR_LN_1:              proto.String(AssertString(cmap["PRTY_ADDR_LN_1"])),
		PRTY_CITY_NM:                proto.String(AssertString(cmap["PRTY_CITY_NM"])),
		PRTY_POSTAL_CD:              proto.String(AssertString(cmap["PRTY_POSTAL_CD"])),
		PRTY_DISTRICT_NM:            proto.String(AssertString(cmap["PRTY_DISTRICT_NM"])),
		BR002_I001:                  proto.Float64(AssertFloat(cmap["BR002_I001"])),
		BR002_I002:                  proto.Float64(AssertFloat(cmap["BR002_I002"])),
		BR002_I003:                  proto.Float64(AssertFloat(cmap["BR002_I003"])),
		BR002_I004:                  proto.Float64(AssertFloat(cmap["BR002_I004"])),
		BR002_I005:                  proto.Float64(AssertFloat(cmap["BR002_I005"])),
		BR002_I006:                  proto.Float64(AssertFloat(cmap["BR002_I006"])),
		BR002_I007:                  proto.Float64(AssertFloat(cmap["BR002_I007"])),
		BR002_I008:                  proto.Float64(AssertFloat(cmap["BR002_I008"])),
		BR002_I009:                  proto.Float64(AssertFloat(cmap["BR002_I009"])),
		BR002_I010:                  proto.Float64(AssertFloat(cmap["BR002_I010"])),
		BR002_I011:                  proto.Float64(AssertFloat(cmap["BR002_I011"])),
		BR002_I012:                  proto.Float64(AssertFloat(cmap["BR002_I012"])),
		BR002_I013:                  proto.Float64(AssertFloat(cmap["BR002_I013"])),
		BR002_I014:                  proto.Float64(AssertFloat(cmap["BR002_I014"])),
		BR002_I015:                  proto.Float64(AssertFloat(cmap["BR002_I015"])),
		BR002_I016:                  proto.Float64(AssertFloat(cmap["BR002_I016"])),
		BR002_I017:                  proto.Float64(AssertFloat(cmap["BR002_I017"])),
		BR002_I018:                  proto.Float64(AssertFloat(cmap["BR002_I018"])),
		BR002_I019:                  proto.Float64(AssertFloat(cmap["BR002_I019"])),
		BR002_I020:                  proto.Float64(AssertFloat(cmap["BR002_I020"])),
		BR002_I021:                  proto.Float64(AssertFloat(cmap["BR002_I021"])),
		BR002_I022:                  proto.Float64(AssertFloat(cmap["BR002_I022"])),
		BR002_I023:                  proto.Float64(AssertFloat(cmap["BR002_I023"])),
		BR002_I024:                  proto.Float64(AssertFloat(cmap["BR002_I024"])),
		BR002_I025:                  proto.Float64(AssertFloat(cmap["BR002_I025"])),
		BR002_I026:                  proto.Float64(AssertFloat(cmap["BR002_I026"])),
		BR002_I027:                  proto.Float64(AssertFloat(cmap["BR002_I027"])),
		BR002_I028:                  proto.Float64(AssertFloat(cmap["BR002_I028"])),
		BR002_I029:                  proto.Float64(AssertFloat(cmap["BR002_I029"])),
		BR002_I030:                  proto.Float64(AssertFloat(cmap["BR002_I030"])),
		BR003_I001:                  proto.Float64(AssertFloat(cmap["BR003_I001"])),
		BR003_I002:                  proto.Float64(AssertFloat(cmap["BR003_I002"])),
		BR003_I003:                  proto.Float64(AssertFloat(cmap["BR003_I003"])),
		BR003_I004:                  proto.Float64(AssertFloat(cmap["BR003_I004"])),
		BR003_I005:                  proto.Float64(AssertFloat(cmap["BR003_I005"])),
		BR003_I006:                  proto.Float64(AssertFloat(cmap["BR003_I006"])),
		BR003_I007:                  proto.Float64(AssertFloat(cmap["BR003_I007"])),
		BR003_I008:                  proto.Float64(AssertFloat(cmap["BR003_I008"])),
		BR003_I009:                  proto.Float64(AssertFloat(cmap["BR003_I009"])),
		BR003_I010:                  proto.Float64(AssertFloat(cmap["BR003_I010"])),
		BR003_I011:                  proto.Float64(AssertFloat(cmap["BR003_I011"])),
		BR003_I012:                  proto.Float64(AssertFloat(cmap["BR003_I012"])),
		BR003_I013:                  proto.Float64(AssertFloat(cmap["BR003_I013"])),
		BR003_I014:                  proto.Float64(AssertFloat(cmap["BR003_I014"])),
		BR003_I015:                  proto.Float64(AssertFloat(cmap["BR003_I015"])),
		BR003_I016:                  proto.Float64(AssertFloat(cmap["BR003_I016"])),
		BR003_I017:                  proto.Float64(AssertFloat(cmap["BR003_I017"])),
		BR003_I018:                  proto.Float64(AssertFloat(cmap["BR003_I018"])),
		BR003_I019:                  proto.Float64(AssertFloat(cmap["BR003_I019"])),
		BR003_I020:                  proto.Float64(AssertFloat(cmap["BR003_I020"])),
		BR003_I021:                  proto.Float64(AssertFloat(cmap["BR003_I021"])),
		BR003_I022:                  proto.Float64(AssertFloat(cmap["BR003_I022"])),
		BR003_I023:                  proto.Float64(AssertFloat(cmap["BR003_I023"])),
		BR003_I024:                  proto.Float64(AssertFloat(cmap["BR003_I024"])),
		BR003_I025:                  proto.Float64(AssertFloat(cmap["BR003_I025"])),
		BR003_I026:                  proto.Float64(AssertFloat(cmap["BR003_I026"])),
		BR003_I027:                  proto.Float64(AssertFloat(cmap["BR003_I027"])),
		BR003_I028:                  proto.Float64(AssertFloat(cmap["BR003_I028"])),
		BR003_I029:                  proto.Float64(AssertFloat(cmap["BR003_I029"])),
		BR003_I030:                  proto.Float64(AssertFloat(cmap["BR003_I030"])),
		BR003_I031:                  proto.Float64(AssertFloat(cmap["BR003_I031"])),
		BR003_I032:                  proto.Float64(AssertFloat(cmap["BR003_I032"])),
		BR003_I033:                  proto.Float64(AssertFloat(cmap["BR003_I033"])),
		BR003_I034:                  proto.Float64(AssertFloat(cmap["BR003_I034"])),
		BR003_I035:                  proto.Float64(AssertFloat(cmap["BR003_I035"])),
		BR006_I001:                  proto.Float64(AssertFloat(cmap["BR006_I001"])),
		BR006_I002:                  proto.Float64(AssertFloat(cmap["BR006_I002"])),
		BR006_I003:                  proto.Float64(AssertFloat(cmap["BR006_I003"])),
		BR006_I004:                  proto.Float64(AssertFloat(cmap["BR006_I004"])),
		BR006_I005:                  proto.Float64(AssertFloat(cmap["BR006_I005"])),
		BR006_I006:                  proto.Float64(AssertFloat(cmap["BR006_I006"])),
		BR006_I007:                  proto.Float64(AssertFloat(cmap["BR006_I007"])),
		BR012_I001:                  proto.Float64(AssertFloat(cmap["BR012_I001"])),
		BR012_I002:                  proto.Float64(AssertFloat(cmap["BR012_I002"])),
		BR012_I003:                  proto.Float64(AssertFloat(cmap["BR012_I003"])),
		BR012_I004:                  proto.Float64(AssertFloat(cmap["BR012_I004"])),
		BR012_I005:                  proto.Float64(AssertFloat(cmap["BR012_I005"])),
		BR031_I001:                  proto.Float64(AssertFloat(cmap["BR031_I001"])),
		BR031_I002:                  proto.Float64(AssertFloat(cmap["BR031_I002"])),
		BR031_I003:                  proto.Float64(AssertFloat(cmap["BR031_I003"])),
		BR031_I004:                  proto.Float64(AssertFloat(cmap["BR031_I004"])),
		BR031_I005:                  proto.Float64(AssertFloat(cmap["BR031_I005"])),
		BR031_I006:                  proto.Float64(AssertFloat(cmap["BR031_I006"])),
		BR031_I007:                  proto.Float64(AssertFloat(cmap["BR031_I007"])),
		BR038_I001:                  proto.Float64(AssertFloat(cmap["BR038_I001"])),
		BR038_I002:                  proto.Float64(AssertFloat(cmap["BR038_I002"])),
		BR038_I003:                  proto.Float64(AssertFloat(cmap["BR038_I003"])),
		BR038_I004:                  proto.Float64(AssertFloat(cmap["BR038_I004"])),
		BR038_I005:                  proto.Float64(AssertFloat(cmap["BR038_I005"])),
		BR038_I006:                  proto.Float64(AssertFloat(cmap["BR038_I006"])),
		BR055_I001:                  proto.Float64(AssertFloat(cmap["BR055_I001"])),
		BR055_I002:                  proto.Float64(AssertFloat(cmap["BR055_I002"])),
		BR055_I003:                  proto.Float64(AssertFloat(cmap["BR055_I003"])),
		BR055_I004:                  proto.Float64(AssertFloat(cmap["BR055_I004"])),
		BR060_I001:                  proto.Float64(AssertFloat(cmap["BR060_I001"])),
		BR060_I002:                  proto.Float64(AssertFloat(cmap["BR060_I002"])),
		BR060_I003:                  proto.Float64(AssertFloat(cmap["BR060_I003"])),
		BR060_I004:                  proto.Float64(AssertFloat(cmap["BR060_I004"])),
		BR060_I005:                  proto.Float64(AssertFloat(cmap["BR060_I005"])),
		BR060_I006:                  proto.Float64(AssertFloat(cmap["BR060_I006"])),
		BR063_I001:                  proto.Float64(AssertFloat(cmap["BR063_I001"])),
		BR063_I002:                  proto.Float64(AssertFloat(cmap["BR063_I002"])),
		BR063_I003:                  proto.Float64(AssertFloat(cmap["BR063_I003"])),
		BR063_I004:                  proto.Float64(AssertFloat(cmap["BR063_I004"])),
		BR063_I005:                  proto.Float64(AssertFloat(cmap["BR063_I005"])),
		BR076_I001:                  proto.Float64(AssertFloat(cmap["BR076_I001"])),
		BR095_I001:                  proto.Float64(AssertFloat(cmap["BR095_I001"])),
		BR095_I002:                  proto.Float64(AssertFloat(cmap["BR095_I002"])),
		BR095_I003:                  proto.Float64(AssertFloat(cmap["BR095_I003"])),
		BR095_I004:                  proto.Float64(AssertFloat(cmap["BR095_I004"])),
		BR095_I005:                  proto.Float64(AssertFloat(cmap["BR095_I005"])),
		BR096_I001:                  proto.Float64(AssertFloat(cmap["BR096_I001"])),
		BR096_I002:                  proto.Float64(AssertFloat(cmap["BR096_I002"])),
		BR096_I003:                  proto.Float64(AssertFloat(cmap["BR096_I003"])),
		BR096_I004:                  proto.Float64(AssertFloat(cmap["BR096_I004"])),
		BR096_I005:                  proto.Float64(AssertFloat(cmap["BR096_I005"])),
		BR102_I001:                  proto.Float64(AssertFloat(cmap["BR102_I001"])),
		BR103_I001:                  proto.Float64(AssertFloat(cmap["BR103_I001"])),
		BR103_I002:                  proto.Float64(AssertFloat(cmap["BR103_I002"])),
		BR108_I001:                  proto.Float64(AssertFloat(cmap["BR108_I001"])),
		BR108_I002:                  proto.Float64(AssertFloat(cmap["BR108_I002"])),
		BR108_I003:                  proto.Float64(AssertFloat(cmap["BR108_I003"])),
		BR109_I001:                  proto.Float64(AssertFloat(cmap["BR109_I001"])),
		BR109_I002:                  proto.Float64(AssertFloat(cmap["BR109_I002"])),
		BR109_I003:                  proto.Float64(AssertFloat(cmap["BR109_I003"])),
		BR109_I004:                  proto.Float64(AssertFloat(cmap["BR109_I004"])),
		BR109_I005:                  proto.Float64(AssertFloat(cmap["BR109_I005"])),
		BR109_I006:                  proto.Float64(AssertFloat(cmap["BR109_I006"])),
		BR109_I007:                  proto.Float64(AssertFloat(cmap["BR109_I007"])),
		BR109_I008:                  proto.Float64(AssertFloat(cmap["BR109_I008"])),
		BR109_I009:                  proto.Float64(AssertFloat(cmap["BR109_I009"])),
		BR109_I010:                  proto.Float64(AssertFloat(cmap["BR109_I010"])),
		BR110_I001:                  proto.Float64(AssertFloat(cmap["BR110_I001"])),
		BR111_I001:                  proto.Float64(AssertFloat(cmap["BR111_I001"])),
		BR112_I001:                  proto.Float64(AssertFloat(cmap["BR112_I001"])),
		BR113_I001:                  proto.Float64(AssertFloat(cmap["BR113_I001"])),
		BR114_I001:                  proto.Float64(AssertFloat(cmap["BR114_I001"])),
		BR114_I002:                  proto.Float64(AssertFloat(cmap["BR114_I002"])),
		BR114_I003:                  proto.Float64(AssertFloat(cmap["BR114_I003"])),
		BR114_I004:                  proto.Float64(AssertFloat(cmap["BR114_I004"])),
		BR115_I001:                  proto.Float64(AssertFloat(cmap["BR115_I001"])),
		BR117_I001:                  proto.Float64(AssertFloat(cmap["BR117_I001"])),
		BR118_I001:                  proto.Float64(AssertFloat(cmap["BR118_I001"])),
		BR119_I001:                  proto.Float64(AssertFloat(cmap["BR119_I001"])),
		BR121_I001:                  proto.Float64(AssertFloat(cmap["BR121_I001"])),
		BR122_I001:                  proto.Float64(AssertFloat(cmap["BR122_I001"])),
		BR123_I001:                  proto.Float64(AssertFloat(cmap["BR123_I001"])),
		BR125_I001:                  proto.Float64(AssertFloat(cmap["BR125_I001"])),
		NEW_INVSTGT_FLG:             proto.Float64(AssertFloat(cmap["NEW_INVSTGT_FLG"])),
		LOAD_DATE:                   proto.Int64(AssertTime(cmap["LOAD_DATE"])),
	}

	return claim
}
