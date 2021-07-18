package disassemble

type register uint32

const (
	REG_NONE register = iota
	REG_W0
	REG_W1
	REG_W2
	REG_W3
	REG_W4
	REG_W5
	REG_W6
	REG_W7
	REG_W8
	REG_W9
	REG_W10
	REG_W11
	REG_W12
	REG_W13
	REG_W14
	REG_W15
	REG_W16
	REG_W17
	REG_W18
	REG_W19
	REG_W20
	REG_W21
	REG_W22
	REG_W23
	REG_W24
	REG_W25
	REG_W26
	REG_W27
	REG_W28
	REG_W29
	REG_W30
	REG_WZR
	REG_WSP
	REG_X0
	REG_X1
	REG_X2
	REG_X3
	REG_X4
	REG_X5
	REG_X6
	REG_X7
	REG_X8
	REG_X9
	REG_X10
	REG_X11
	REG_X12
	REG_X13
	REG_X14
	REG_X15
	REG_X16
	REG_X17
	REG_X18
	REG_X19
	REG_X20
	REG_X21
	REG_X22
	REG_X23
	REG_X24
	REG_X25
	REG_X26
	REG_X27
	REG_X28
	REG_X29
	REG_X30
	REG_XZR
	REG_SP
	REG_V0
	REG_V1
	REG_V2
	REG_V3
	REG_V4
	REG_V5
	REG_V6
	REG_V7
	REG_V8
	REG_V9
	REG_V10
	REG_V11
	REG_V12
	REG_V13
	REG_V14
	REG_V15
	REG_V16
	REG_V17
	REG_V18
	REG_V19
	REG_V20
	REG_V21
	REG_V22
	REG_V23
	REG_V24
	REG_V25
	REG_V26
	REG_V27
	REG_V28
	REG_V29
	REG_V30
	REG_VZR
	REG_V31
	// B vector
	REG_B0
	REG_B1
	REG_B2
	REG_B3
	REG_B4
	REG_B5
	REG_B6
	REG_B7
	REG_B8
	REG_B9
	REG_B10
	REG_B11
	REG_B12
	REG_B13
	REG_B14
	REG_B15
	REG_B16
	REG_B17
	REG_B18
	REG_B19
	REG_B20
	REG_B21
	REG_B22
	REG_B23
	REG_B24
	REG_B25
	REG_B26
	REG_B27
	REG_B28
	REG_B29
	REG_B30
	REG_BZR
	REG_B31
	// H vector
	REG_H0
	REG_H1
	REG_H2
	REG_H3
	REG_H4
	REG_H5
	REG_H6
	REG_H7
	REG_H8
	REG_H9
	REG_H10
	REG_H11
	REG_H12
	REG_H13
	REG_H14
	REG_H15
	REG_H16
	REG_H17
	REG_H18
	REG_H19
	REG_H20
	REG_H21
	REG_H22
	REG_H23
	REG_H24
	REG_H25
	REG_H26
	REG_H27
	REG_H28
	REG_H29
	REG_H30
	REG_HZR
	REG_H31
	REG_S0
	REG_S1
	REG_S2
	REG_S3
	REG_S4
	REG_S5
	REG_S6
	REG_S7
	REG_S8
	REG_S9
	REG_S10
	REG_S11
	REG_S12
	REG_S13
	REG_S14
	REG_S15
	REG_S16
	REG_S17
	REG_S18
	REG_S19
	REG_S20
	REG_S21
	REG_S22
	REG_S23
	REG_S24
	REG_S25
	REG_S26
	REG_S27
	REG_S28
	REG_S29
	REG_S30
	REG_SZR
	REG_S31
	REG_D0
	REG_D1
	REG_D2
	REG_D3
	REG_D4
	REG_D5
	REG_D6
	REG_D7
	REG_D8
	REG_D9
	REG_D10
	REG_D11
	REG_D12
	REG_D13
	REG_D14
	REG_D15
	REG_D16
	REG_D17
	REG_D18
	REG_D19
	REG_D20
	REG_D21
	REG_D22
	REG_D23
	REG_D24
	REG_D25
	REG_D26
	REG_D27
	REG_D28
	REG_D29
	REG_D30
	REG_DZR
	REG_D31
	REG_Q0
	REG_Q1
	REG_Q2
	REG_Q3
	REG_Q4
	REG_Q5
	REG_Q6
	REG_Q7
	REG_Q8
	REG_Q9
	REG_Q10
	REG_Q11
	REG_Q12
	REG_Q13
	REG_Q14
	REG_Q15
	REG_Q16
	REG_Q17
	REG_Q18
	REG_Q19
	REG_Q20
	REG_Q21
	REG_Q22
	REG_Q23
	REG_Q24
	REG_Q25
	REG_Q26
	REG_Q27
	REG_Q28
	REG_Q29
	REG_Q30
	REG_QZR
	REG_Q31
	// B vector
	REG_V0_B0
	REG_V0_B1
	REG_V0_B2
	REG_V0_B3
	REG_V0_B4
	REG_V0_B5
	REG_V0_B6
	REG_V0_B7
	REG_V0_B8
	REG_V0_B9
	REG_V0_B10
	REG_V0_B11
	REG_V0_B12
	REG_V0_B13
	REG_V0_B14
	REG_V0_B15
	REG_V1_B0
	REG_V1_B1
	REG_V1_B2
	REG_V1_B3
	REG_V1_B4
	REG_V1_B5
	REG_V1_B6
	REG_V1_B7
	REG_V1_B8
	REG_V1_B9
	REG_V1_B10
	REG_V1_B11
	REG_V1_B12
	REG_V1_B13
	REG_V1_B14
	REG_V1_B15
	REG_V2_B0
	REG_V2_B1
	REG_V2_B2
	REG_V2_B3
	REG_V2_B4
	REG_V2_B5
	REG_V2_B6
	REG_V2_B7
	REG_V2_B8
	REG_V2_B9
	REG_V2_B10
	REG_V2_B11
	REG_V2_B12
	REG_V2_B13
	REG_V2_B14
	REG_V2_B15
	REG_V3_B0
	REG_V3_B1
	REG_V3_B2
	REG_V3_B3
	REG_V3_B4
	REG_V3_B5
	REG_V3_B6
	REG_V3_B7
	REG_V3_B8
	REG_V3_B9
	REG_V3_B10
	REG_V3_B11
	REG_V3_B12
	REG_V3_B13
	REG_V3_B14
	REG_V3_B15
	REG_V4_B0
	REG_V4_B1
	REG_V4_B2
	REG_V4_B3
	REG_V4_B4
	REG_V4_B5
	REG_V4_B6
	REG_V4_B7
	REG_V4_B8
	REG_V4_B9
	REG_V4_B10
	REG_V4_B11
	REG_V4_B12
	REG_V4_B13
	REG_V4_B14
	REG_V4_B15
	REG_V5_B0
	REG_V5_B1
	REG_V5_B2
	REG_V5_B3
	REG_V5_B4
	REG_V5_B5
	REG_V5_B6
	REG_V5_B7
	REG_V5_B8
	REG_V5_B9
	REG_V5_B10
	REG_V5_B11
	REG_V5_B12
	REG_V5_B13
	REG_V5_B14
	REG_V5_B15
	REG_V6_B0
	REG_V6_B1
	REG_V6_B2
	REG_V6_B3
	REG_V6_B4
	REG_V6_B5
	REG_V6_B6
	REG_V6_B7
	REG_V6_B8
	REG_V6_B9
	REG_V6_B10
	REG_V6_B11
	REG_V6_B12
	REG_V6_B13
	REG_V6_B14
	REG_V6_B15
	REG_V7_B0
	REG_V7_B1
	REG_V7_B2
	REG_V7_B3
	REG_V7_B4
	REG_V7_B5
	REG_V7_B6
	REG_V7_B7
	REG_V7_B8
	REG_V7_B9
	REG_V7_B10
	REG_V7_B11
	REG_V7_B12
	REG_V7_B13
	REG_V7_B14
	REG_V7_B15
	REG_V8_B0
	REG_V8_B1
	REG_V8_B2
	REG_V8_B3
	REG_V8_B4
	REG_V8_B5
	REG_V8_B6
	REG_V8_B7
	REG_V8_B8
	REG_V8_B9
	REG_V8_B10
	REG_V8_B11
	REG_V8_B12
	REG_V8_B13
	REG_V8_B14
	REG_V8_B15
	REG_V9_B0
	REG_V9_B1
	REG_V9_B2
	REG_V9_B3
	REG_V9_B4
	REG_V9_B5
	REG_V9_B6
	REG_V9_B7
	REG_V9_B8
	REG_V9_B9
	REG_V9_B10
	REG_V9_B11
	REG_V9_B12
	REG_V9_B13
	REG_V9_B14
	REG_V9_B15
	REG_V10_B0
	REG_V10_B1
	REG_V10_B2
	REG_V10_B3
	REG_V10_B4
	REG_V10_B5
	REG_V10_B6
	REG_V10_B7
	REG_V10_B8
	REG_V10_B9
	REG_V10_B10
	REG_V10_B11
	REG_V10_B12
	REG_V10_B13
	REG_V10_B14
	REG_V10_B15
	REG_V11_B0
	REG_V11_B1
	REG_V11_B2
	REG_V11_B3
	REG_V11_B4
	REG_V11_B5
	REG_V11_B6
	REG_V11_B7
	REG_V11_B8
	REG_V11_B9
	REG_V11_B10
	REG_V11_B11
	REG_V11_B12
	REG_V11_B13
	REG_V11_B14
	REG_V11_B15
	REG_V12_B0
	REG_V12_B1
	REG_V12_B2
	REG_V12_B3
	REG_V12_B4
	REG_V12_B5
	REG_V12_B6
	REG_V12_B7
	REG_V12_B8
	REG_V12_B9
	REG_V12_B10
	REG_V12_B11
	REG_V12_B12
	REG_V12_B13
	REG_V12_B14
	REG_V12_B15
	REG_V13_B0
	REG_V13_B1
	REG_V13_B2
	REG_V13_B3
	REG_V13_B4
	REG_V13_B5
	REG_V13_B6
	REG_V13_B7
	REG_V13_B8
	REG_V13_B9
	REG_V13_B10
	REG_V13_B11
	REG_V13_B12
	REG_V13_B13
	REG_V13_B14
	REG_V13_B15
	REG_V14_B0
	REG_V14_B1
	REG_V14_B2
	REG_V14_B3
	REG_V14_B4
	REG_V14_B5
	REG_V14_B6
	REG_V14_B7
	REG_V14_B8
	REG_V14_B9
	REG_V14_B10
	REG_V14_B11
	REG_V14_B12
	REG_V14_B13
	REG_V14_B14
	REG_V14_B15
	REG_V15_B0
	REG_V15_B1
	REG_V15_B2
	REG_V15_B3
	REG_V15_B4
	REG_V15_B5
	REG_V15_B6
	REG_V15_B7
	REG_V15_B8
	REG_V15_B9
	REG_V15_B10
	REG_V15_B11
	REG_V15_B12
	REG_V15_B13
	REG_V15_B14
	REG_V15_B15
	REG_V16_B0
	REG_V16_B1
	REG_V16_B2
	REG_V16_B3
	REG_V16_B4
	REG_V16_B5
	REG_V16_B6
	REG_V16_B7
	REG_V16_B8
	REG_V16_B9
	REG_V16_B10
	REG_V16_B11
	REG_V16_B12
	REG_V16_B13
	REG_V16_B14
	REG_V16_B15
	REG_V17_B0
	REG_V17_B1
	REG_V17_B2
	REG_V17_B3
	REG_V17_B4
	REG_V17_B5
	REG_V17_B6
	REG_V17_B7
	REG_V17_B8
	REG_V17_B9
	REG_V17_B10
	REG_V17_B11
	REG_V17_B12
	REG_V17_B13
	REG_V17_B14
	REG_V17_B15
	REG_V18_B0
	REG_V18_B1
	REG_V18_B2
	REG_V18_B3
	REG_V18_B4
	REG_V18_B5
	REG_V18_B6
	REG_V18_B7
	REG_V18_B8
	REG_V18_B9
	REG_V18_B10
	REG_V18_B11
	REG_V18_B12
	REG_V18_B13
	REG_V18_B14
	REG_V18_B15
	REG_V19_B0
	REG_V19_B1
	REG_V19_B2
	REG_V19_B3
	REG_V19_B4
	REG_V19_B5
	REG_V19_B6
	REG_V19_B7
	REG_V19_B8
	REG_V19_B9
	REG_V19_B10
	REG_V19_B11
	REG_V19_B12
	REG_V19_B13
	REG_V19_B14
	REG_V19_B15
	REG_V20_B0
	REG_V20_B1
	REG_V20_B2
	REG_V20_B3
	REG_V20_B4
	REG_V20_B5
	REG_V20_B6
	REG_V20_B7
	REG_V20_B8
	REG_V20_B9
	REG_V20_B10
	REG_V20_B11
	REG_V20_B12
	REG_V20_B13
	REG_V20_B14
	REG_V20_B15
	REG_V21_B0
	REG_V21_B1
	REG_V21_B2
	REG_V21_B3
	REG_V21_B4
	REG_V21_B5
	REG_V21_B6
	REG_V21_B7
	REG_V21_B8
	REG_V21_B9
	REG_V21_B10
	REG_V21_B11
	REG_V21_B12
	REG_V21_B13
	REG_V21_B14
	REG_V21_B15
	REG_V22_B0
	REG_V22_B1
	REG_V22_B2
	REG_V22_B3
	REG_V22_B4
	REG_V22_B5
	REG_V22_B6
	REG_V22_B
	REG_V22_B8
	REG_V22_B9
	REG_V22_B10
	REG_V22_B11
	REG_V22_B12
	REG_V22_B13
	REG_V22_B14
	REG_V22_B15
	REG_V23_B0
	REG_V23_B1
	REG_V23_B2
	REG_V23_B3
	REG_V23_B4
	REG_V23_B5
	REG_V23_B6
	REG_V23_B7
	REG_V23_B8
	REG_V23_B9
	REG_V23_B10
	REG_V23_B11
	REG_V23_B12
	REG_V23_B13
	REG_V23_B14
	REG_V23_B15
	REG_V24_B0
	REG_V24_B1
	REG_V24_B2
	REG_V24_B3
	REG_V24_B4
	REG_V24_B5
	REG_V24_B6
	REG_V24_B7
	REG_V24_B8
	REG_V24_B9
	REG_V24_B10
	REG_V24_B11
	REG_V24_B12
	REG_V24_B13
	REG_V24_B14
	REG_V24_B15
	REG_V25_B0
	REG_V25_B1
	REG_V25_B2
	REG_V25_B3
	REG_V25_B4
	REG_V25_B5
	REG_V25_B6
	REG_V25_B7
	REG_V25_B8
	REG_V25_B9
	REG_V25_B10
	REG_V25_B11
	REG_V25_B12
	REG_V25_B13
	REG_V25_B14
	REG_V25_B15
	REG_V26_B0
	REG_V26_B1
	REG_V26_B2
	REG_V26_B3
	REG_V26_B4
	REG_V26_B5
	REG_V26_B6
	REG_V26_B7
	REG_V26_B8
	REG_V26_B9
	REG_V26_B10
	REG_V26_B11
	REG_V26_B12
	REG_V26_B13
	REG_V26_B14
	REG_V26_B15
	REG_V27_B0
	REG_V27_B1
	REG_V27_B2
	REG_V27_B3
	REG_V27_B4
	REG_V27_B5
	REG_V27_B6
	REG_V27_B7
	REG_V27_B8
	REG_V27_B9
	REG_V27_B10
	REG_V27_B11
	REG_V27_B12
	REG_V27_B13
	REG_V27_B14
	REG_V27_B15
	REG_V28_B0
	REG_V28_B1
	REG_V28_B2
	REG_V28_B3
	REG_V28_B4
	REG_V28_B5
	REG_V28_B6
	REG_V28_B7
	REG_V28_B8
	REG_V28_B9
	REG_V28_B10
	REG_V28_B11
	REG_V28_B12
	REG_V28_B13
	REG_V28_B14
	REG_V28_B15
	REG_V29_B0
	REG_V29_B1
	REG_V29_B2
	REG_V29_B3
	REG_V29_B4
	REG_V29_B5
	REG_V29_B6
	REG_V29_B7
	REG_V29_B8
	REG_V29_B9
	REG_V29_B10
	REG_V29_B11
	REG_V29_B12
	REG_V29_B13
	REG_V29_B14
	REG_V29_B15
	REG_V30_B0
	REG_V30_B1
	REG_V30_B2
	REG_V30_B3
	REG_V30_B4
	REG_V30_B5
	REG_V30_B6
	REG_V30_B7
	REG_V30_B8
	REG_V30_B9
	REG_V30_B10
	REG_V30_B11
	REG_V30_B12
	REG_V30_B13
	REG_V30_B14
	REG_V30_B15
	REG_V31_B0
	REG_V31_B1
	REG_V31_B2
	REG_V31_B3
	REG_V31_B4
	REG_V31_B5
	REG_V31_B6
	REG_V31_B7
	REG_V31_B8
	REG_V31_B9
	REG_V31_B10
	REG_V31_B11
	REG_V31_B12
	REG_V31_B13
	REG_V31_B14
	REG_V31_B15
	// H vector
	REG_V0_H0
	REG_V0_H1
	REG_V0_H2
	REG_V0_H3
	REG_V0_H4
	REG_V0_H5
	REG_V0_H6
	REG_V0_H7
	REG_V1_H0
	REG_V1_H1
	REG_V1_H2
	REG_V1_H3
	REG_V1_H4
	REG_V1_H5
	REG_V1_H6
	REG_V1_H7
	REG_V2_H0
	REG_V2_H1
	REG_V2_H2
	REG_V2_H3
	REG_V2_H4
	REG_V2_H5
	REG_V2_H6
	REG_V2_H7
	REG_V3_H0
	REG_V3_H1
	REG_V3_H2
	REG_V3_H3
	REG_V3_H4
	REG_V3_H5
	REG_V3_H6
	REG_V3_H7
	REG_V4_H0
	REG_V4_H1
	REG_V4_H2
	REG_V4_H3
	REG_V4_H4
	REG_V4_H5
	REG_V4_H6
	REG_V4_H7
	REG_V5_H0
	REG_V5_H1
	REG_V5_H2
	REG_V5_H3
	REG_V5_H4
	REG_V5_H5
	REG_V5_H6
	REG_V5_H7
	REG_V6_H0
	REG_V6_H1
	REG_V6_H2
	REG_V6_H3
	REG_V6_H4
	REG_V6_H5
	REG_V6_H6
	REG_V6_H7
	REG_V7_H0
	REG_V7_H1
	REG_V7_H2
	REG_V7_H3
	REG_V7_H4
	REG_V7_H5
	REG_V7_H6
	REG_V7_H7
	REG_V8_H0
	REG_V8_H1
	REG_V8_H2
	REG_V8_H3
	REG_V8_H4
	REG_V8_H5
	REG_V8_H6
	REG_V8_H7
	REG_V9_H0
	REG_V9_H1
	REG_V9_H2
	REG_V9_H3
	REG_V9_H4
	REG_V9_H5
	REG_V9_H6
	REG_V9_H7
	REG_V10_H0
	REG_V10_H1
	REG_V10_H2
	REG_V10_H3
	REG_V10_H4
	REG_V10_H5
	REG_V10_H6
	REG_V10_H7
	REG_V11_H0
	REG_V11_H1
	REG_V11_H2
	REG_V11_H3
	REG_V11_H4
	REG_V11_H5
	REG_V11_H6
	REG_V11_H7
	REG_V12_H0
	REG_V12_H1
	REG_V12_H2
	REG_V12_H3
	REG_V12_H4
	REG_V12_H5
	REG_V12_H6
	REG_V12_H7
	REG_V13_H0
	REG_V13_H1
	REG_V13_H2
	REG_V13_H3
	REG_V13_H4
	REG_V13_H5
	REG_V13_H6
	REG_V13_H7
	REG_V14_H0
	REG_V14_H1
	REG_V14_H2
	REG_V14_H3
	REG_V14_H4
	REG_V14_H5
	REG_V14_H6
	REG_V14_H7
	REG_V15_H0
	REG_V15_H1
	REG_V15_H2
	REG_V15_H3
	REG_V15_H4
	REG_V15_H5
	REG_V15_H6
	REG_V15_H7
	REG_V16_H0
	REG_V16_H1
	REG_V16_H2
	REG_V16_H3
	REG_V16_H4
	REG_V16_H5
	REG_V16_H6
	REG_V16_H7
	REG_V17_H0
	REG_V17_H1
	REG_V17_H2
	REG_V17_H3
	REG_V17_H4
	REG_V17_H5
	REG_V17_H6
	REG_V17_H7
	REG_V18_H0
	REG_V18_H1
	REG_V18_H2
	REG_V18_H3
	REG_V18_H4
	REG_V18_H5
	REG_V18_H6
	REG_V18_H7
	REG_V19_H0
	REG_V19_H1
	REG_V19_H2
	REG_V19_H3
	REG_V19_H4
	REG_V19_H5
	REG_V19_H6
	REG_V19_H7
	REG_V20_H0
	REG_V20_H1
	REG_V20_H2
	REG_V20_H3
	REG_V20_H4
	REG_V20_H5
	REG_V20_H6
	REG_V20_H7
	REG_V21_H0
	REG_V21_H1
	REG_V21_H2
	REG_V21_H3
	REG_V21_H4
	REG_V21_H5
	REG_V21_H6
	REG_V21_H7
	REG_V22_H0
	REG_V22_H1
	REG_V22_H2
	REG_V22_H3
	REG_V22_H4
	REG_V22_H5
	REG_V22_H6
	REG_V22_H7
	REG_V23_H0
	REG_V23_H1
	REG_V23_H2
	REG_V23_H3
	REG_V23_H4
	REG_V23_H5
	REG_V23_H6
	REG_V23_H7
	REG_V24_H0
	REG_V24_H1
	REG_V24_H2
	REG_V24_H3
	REG_V24_H4
	REG_V24_H5
	REG_V24_H6
	REG_V24_H7
	REG_V25_H0
	REG_V25_H1
	REG_V25_H2
	REG_V25_H3
	REG_V25_H4
	REG_V25_H5
	REG_V25_H6
	REG_V25_H7
	REG_V26_H0
	REG_V26_H1
	REG_V26_H2
	REG_V26_H3
	REG_V26_H4
	REG_V26_H5
	REG_V26_H6
	REG_V26_H7
	REG_V27_H0
	REG_V27_H1
	REG_V27_H2
	REG_V27_H3
	REG_V27_H4
	REG_V27_H5
	REG_V27_H6
	REG_V27_H7
	REG_V28_H0
	REG_V28_H1
	REG_V28_H2
	REG_V28_H3
	REG_V28_H4
	REG_V28_H5
	REG_V28_H6
	REG_V28_H7
	REG_V29_H0
	REG_V29_H1
	REG_V29_H2
	REG_V29_H3
	REG_V29_H4
	REG_V29_H5
	REG_V29_H6
	REG_V29_H7
	REG_V30_H0
	REG_V30_H1
	REG_V30_H2
	REG_V30_H3
	REG_V30_H4
	REG_V30_H5
	REG_V30_H6
	REG_V30_H7
	REG_V31_H0
	REG_V31_H1
	REG_V31_H2
	REG_V31_H3
	REG_V31_H4
	REG_V31_H5
	REG_V31_H6
	REG_V31_H7
	// S vector
	REG_V0_S0
	REG_V0_S1
	REG_V0_S2
	REG_V0_S3
	REG_V1_S0
	REG_V1_S1
	REG_V1_S2
	REG_V1_S3
	REG_V2_S0
	REG_V2_S1
	REG_V2_S2
	REG_V2_S3
	REG_V3_S0
	REG_V3_S1
	REG_V3_S2
	REG_V3_S3
	REG_V4_S0
	REG_V4_S1
	REG_V4_S2
	REG_V4_S3
	REG_V5_S0
	REG_V5_S1
	REG_V5_S2
	REG_V5_S3
	REG_V6_S0
	REG_V6_S1
	REG_V6_S2
	REG_V6_S3
	REG_V7_S0
	REG_V7_S1
	REG_V7_S2
	REG_V7_S3
	REG_V8_S0
	REG_V8_S1
	REG_V8_S2
	REG_V8_S3
	REG_V9_S0
	REG_V9_S1
	REG_V9_S2
	REG_V9_S3
	REG_V10_S0
	REG_V10_S1
	REG_V10_S2
	REG_V10_S3
	REG_V11_S0
	REG_V11_S1
	REG_V11_S2
	REG_V11_S3
	REG_V12_S0
	REG_V12_S1
	REG_V12_S2
	REG_V12_S3
	REG_V13_S0
	REG_V13_S1
	REG_V13_S2
	REG_V13_S3
	REG_V14_S0
	REG_V14_S1
	REG_V14_S2
	REG_V14_S3
	REG_V15_S0
	REG_V15_S1
	REG_V15_S2
	REG_V15_S3
	REG_V16_S0
	REG_V16_S1
	REG_V16_S2
	REG_V16_S3
	REG_V17_S0
	REG_V17_S1
	REG_V17_S2
	REG_V17_S3
	REG_V18_S0
	REG_V18_S1
	REG_V18_S2
	REG_V18_S3
	REG_V19_S0
	REG_V19_S1
	REG_V19_S2
	REG_V19_S3
	REG_V20_S0
	REG_V20_S1
	REG_V20_S2
	REG_V20_S3
	REG_V21_S0
	REG_V21_S1
	REG_V21_S2
	REG_V21_S3
	REG_V22_S0
	REG_V22_S1
	REG_V22_S2
	REG_V22_S3
	REG_V23_S0
	REG_V23_S1
	REG_V23_S2
	REG_V23_S3
	REG_V24_S0
	REG_V24_S1
	REG_V24_S2
	REG_V24_S3
	REG_V25_S0
	REG_V25_S1
	REG_V25_S2
	REG_V25_S3
	REG_V26_S0
	REG_V26_S1
	REG_V26_S2
	REG_V26_S3
	REG_V27_S0
	REG_V27_S1
	REG_V27_S2
	REG_V27_S3
	REG_V28_S0
	REG_V28_S1
	REG_V28_S2
	REG_V28_S3
	REG_V29_S0
	REG_V29_S1
	REG_V29_S2
	REG_V29_S3
	REG_V30_S0
	REG_V30_S1
	REG_V30_S2
	REG_V30_S3
	REG_V31_S0
	REG_V31_S1
	REG_V31_S2
	REG_V31_S3
	// D vector
	REG_V0_D0
	REG_V0_D1
	REG_V1_D0
	REG_V1_D1
	REG_V2_D0
	REG_V2_D1
	REG_V3_D0
	REG_V3_D1
	REG_V4_D0
	REG_V4_D1
	REG_V5_D0
	REG_V5_D1
	REG_V6_D0
	REG_V6_D1
	REG_V7_D0
	REG_V7_D1
	REG_V8_D0
	REG_V8_D1
	REG_V9_D0
	REG_V9_D1
	REG_V10_D0
	REG_V10_D1
	REG_V11_D0
	REG_V11_D1
	REG_V12_D0
	REG_V12_D1
	REG_V13_D0
	REG_V13_D1
	REG_V14_D0
	REG_V14_D1
	REG_V15_D0
	REG_V15_D1
	REG_V16_D0
	REG_V16_D1
	REG_V17_D0
	REG_V17_D1
	REG_V18_D0
	REG_V18_D1
	REG_V19_D0
	REG_V19_D1
	REG_V20_D0
	REG_V20_D1
	REG_V21_D0
	REG_V21_D1
	REG_V22_D0
	REG_V22_D1
	REG_V23_D0
	REG_V23_D1
	REG_V24_D0
	REG_V24_D1
	REG_V25_D0
	REG_V25_D1
	REG_V26_D0
	REG_V26_D1
	REG_V27_D0
	REG_V27_D1
	REG_V28_D0
	REG_V28_D1
	REG_V29_D0
	REG_V29_D1
	REG_V30_D0
	REG_V30_D1
	REG_V31_D0
	REG_V31_D1
	// Q vector is already defined REG_V0, REG_V1, ..., REG_V31
	// SVE
	REG_Z0
	REG_Z1
	REG_Z2
	REG_Z3
	REG_Z4
	REG_Z5
	REG_Z6
	REG_Z7
	REG_Z8
	REG_Z9
	REG_Z10
	REG_Z11
	REG_Z12
	REG_Z13
	REG_Z14
	REG_Z15
	REG_Z16
	REG_Z17
	REG_Z18
	REG_Z19
	REG_Z20
	REG_Z21
	REG_Z22
	REG_Z23
	REG_Z24
	REG_Z25
	REG_Z26
	REG_Z27
	REG_Z28
	REG_Z29
	REG_Z30
	REG_ZZR
	REG_Z31
	REG_P0
	REG_P1
	REG_P2
	REG_P3
	REG_P4
	REG_P5
	REG_P6
	REG_P7
	REG_P8
	REG_P9
	REG_P10
	REG_P11
	REG_P12
	REG_P13
	REG_P14
	REG_P15
	REG_P16
	REG_P17
	REG_P18
	REG_P19
	REG_P20
	REG_P21
	REG_P22
	REG_P23
	REG_P24
	REG_P25
	REG_P26
	REG_P27
	REG_P28
	REG_P29
	REG_P30
	REG_P31
	REG_PF0
	REG_PF1
	REG_PF2
	REG_PF3
	REG_PF4
	REG_PF5
	REG_PF6
	REG_PF7
	REG_PF8
	REG_PF9
	REG_PF10
	REG_PF11
	REG_PF12
	REG_PF13
	REG_PF14
	REG_PF15
	REG_PF16
	REG_PF17
	REG_PF18
	REG_PF19
	REG_PF20
	REG_PF21
	REG_PF22
	REG_PF23
	REG_PF24
	REG_PF25
	REG_PF26
	REG_PF27
	REG_PF28
	REG_PF29
	REG_PF30
	REG_PF31

	REG_END
)

var regString = []string{
	"NONE",
	"w0", "w1", "w2", "w3", "w4", "w5", "w6", "w7",
	"w8", "w9", "w10", "w11", "w12", "w13", "w14", "w15",
	"w16", "w17", "w18", "w19", "w20", "w21", "w22", "w23",
	"w24", "w25", "w26", "w27", "w28", "w29", "w30", "wzr", "wsp",
	"x0", "x1", "x2", "x3", "x4", "x5", "x6", "x7",
	"x8", "x9", "x10", "x11", "x12", "x13", "x14", "x15",
	"x16", "x17", "x18", "x19", "x20", "x21", "x22", "x23",
	"x24", "x25", "x26", "x27", "x28", "x29", "x30", "xzr", "sp",
	"v0", "v1", "v2", "v3", "v4", "v5", "v6", "v7",
	"v8", "v9", "v10", "v11", "v12", "v13", "v14", "v15",
	"v16", "v17", "v18", "v19", "v20", "v21", "v22", "v23",
	"v24", "v25", "v26", "v27", "v28", "v29", "v30", "v31", "v31",
	"b0", "b1", "b2", "b3", "b4", "b5", "b6", "b7",
	"b8", "b9", "b10", "b11", "b12", "b13", "b14", "b15",
	"b16", "b17", "b18", "b19", "b20", "b21", "b22", "b23",
	"b24", "b25", "b26", "b27", "b28", "b29", "b30", "b31", "b31",
	"h0", "h1", "h2", "h3", "h4", "h5", "h6", "h7",
	"h8", "h9", "h10", "h11", "h12", "h13", "h14", "h15",
	"h16", "h17", "h18", "h19", "h20", "h21", "h22", "h23",
	"h24", "h25", "h26", "h27", "h28", "h29", "h30", "h31", "h31",
	"s0", "s1", "s2", "s3", "s4", "s5", "s6", "s7",
	"s8", "s9", "s10", "s11", "s12", "s13", "s14", "s15",
	"s16", "s17", "s18", "s19", "s20", "s21", "s22", "s23",
	"s24", "s25", "s26", "s27", "s28", "s29", "s30", "s31", "s31",
	"d0", "d1", "d2", "d3", "d4", "d5", "d6", "d7",
	"d8", "d9", "d10", "d11", "d12", "d13", "d14", "d15",
	"d16", "d17", "d18", "d19", "d20", "d21", "d22", "d23",
	"d24", "d25", "d26", "d27", "d28", "d29", "d30", "d31", "d31",
	"q0", "q1", "q2", "q3", "q4", "q5", "q6", "q7",
	"q8", "q9", "q10", "q11", "q12", "q13", "q14", "q15",
	"q16", "q17", "q18", "q19", "q20", "q21", "q22", "q23",
	"q24", "q25", "q26", "q27", "q28", "q29", "q30", "q31", "q31",
	// B vectors
	"v0.b[0]", "v0.b[1]", "v0.b[2]", "v0.b[3]", "v0.b[4]", "v0.b[5]", "v0.b[6]", "v0.b[7]",
	"v0.b[8]", "v0.b[9]", "v0.b[10]", "v0.b[11]", "v0.b[12]", "v0.b[13]", "v0.b[14]", "v0.b[15]",
	"v1.b[0]", "v1.b[1]", "v1.b[2]", "v1.b[3]", "v1.b[4]", "v1.b[5]", "v1.b[6]", "v1.b[7]",
	"v1.b[8]", "v1.b[9]", "v1.b[10]", "v1.b[11]", "v1.b[12]", "v1.b[13]", "v1.b[14]", "v1.b[15]",
	"v2.b[0]", "v2.b[1]", "v2.b[2]", "v2.b[3]", "v2.b[4]", "v2.b[5]", "v2.b[6]", "v2.b[7]",
	"v2.b[8]", "v2.b[9]", "v2.b[10]", "v2.b[11]", "v2.b[12]", "v2.b[13]", "v2.b[14]", "v2.b[15]",
	"v3.b[0]", "v3.b[1]", "v3.b[2]", "v3.b[3]", "v3.b[4]", "v3.b[5]", "v3.b[6]", "v3.b[7]",
	"v3.b[8]", "v3.b[9]", "v3.b[10]", "v3.b[11]", "v3.b[12]", "v3.b[13]", "v3.b[14]", "v3.b[15]",
	"v4.b[0]", "v4.b[1]", "v4.b[2]", "v4.b[3]", "v4.b[4]", "v4.b[5]", "v4.b[6]", "v4.b[7]",
	"v4.b[8]", "v4.b[9]", "v4.b[10]", "v4.b[11]", "v4.b[12]", "v4.b[13]", "v4.b[14]", "v4.b[15]",
	"v5.b[0]", "v5.b[1]", "v5.b[2]", "v5.b[3]", "v5.b[4]", "v5.b[5]", "v5.b[6]", "v5.b[7]",
	"v5.b[8]", "v5.b[9]", "v5.b[10]", "v5.b[11]", "v5.b[12]", "v5.b[13]", "v5.b[14]", "v5.b[15]",
	"v6.b[0]", "v6.b[1]", "v6.b[2]", "v6.b[3]", "v6.b[4]", "v6.b[5]", "v6.b[6]", "v6.b[7]",
	"v6.b[8]", "v6.b[9]", "v6.b[10]", "v6.b[11]", "v6.b[12]", "v6.b[13]", "v6.b[14]", "v6.b[15]",
	"v7.b[0]", "v7.b[1]", "v7.b[2]", "v7.b[3]", "v7.b[4]", "v7.b[5]", "v7.b[6]", "v7.b[7]",
	"v7.b[8]", "v7.b[9]", "v7.b[10]", "v7.b[11]", "v7.b[12]", "v7.b[13]", "v7.b[14]", "v7.b[15]",
	"v8.b[0]", "v8.b[1]", "v8.b[2]", "v8.b[3]", "v8.b[4]", "v8.b[5]", "v8.b[6]", "v8.b[7]",
	"v8.b[8]", "v8.b[9]", "v8.b[10]", "v8.b[11]", "v8.b[12]", "v8.b[13]", "v8.b[14]", "v8.b[15]",
	"v9.b[0]", "v9.b[1]", "v9.b[2]", "v9.b[3]", "v9.b[4]", "v9.b[5]", "v9.b[6]", "v9.b[7]",
	"v9.b[8]", "v9.b[9]", "v9.b[10]", "v9.b[11]", "v9.b[12]", "v9.b[13]", "v9.b[14]", "v9.b[15]",
	"v10.b[0]", "v10.b[1]", "v10.b[2]", "v10.b[3]", "v10.b[4]", "v10.b[5]", "v10.b[6]", "v10.b[7]",
	"v10.b[8]", "v10.b[9]", "v10.b[10]", "v10.b[11]", "v10.b[12]", "v10.b[13]", "v10.b[14]", "v10.b[15]",
	"v11.b[0]", "v11.b[1]", "v11.b[2]", "v11.b[3]", "v11.b[4]", "v11.b[5]", "v11.b[6]", "v11.b[7]",
	"v11.b[8]", "v11.b[9]", "v11.b[10]", "v11.b[11]", "v11.b[12]", "v11.b[13]", "v11.b[14]", "v11.b[15]",
	"v12.b[0]", "v12.b[1]", "v12.b[2]", "v12.b[3]", "v12.b[4]", "v12.b[5]", "v12.b[6]", "v12.b[7]",
	"v12.b[8]", "v12.b[9]", "v12.b[10]", "v12.b[11]", "v12.b[12]", "v12.b[13]", "v12.b[14]", "v12.b[15]",
	"v13.b[0]", "v13.b[1]", "v13.b[2]", "v13.b[3]", "v13.b[4]", "v13.b[5]", "v13.b[6]", "v13.b[7]",
	"v13.b[8]", "v13.b[9]", "v13.b[10]", "v13.b[11]", "v13.b[12]", "v13.b[13]", "v13.b[14]", "v13.b[15]",
	"v14.b[0]", "v14.b[1]", "v14.b[2]", "v14.b[3]", "v14.b[4]", "v14.b[5]", "v14.b[6]", "v14.b[7]",
	"v14.b[8]", "v14.b[9]", "v14.b[10]", "v14.b[11]", "v14.b[12]", "v14.b[13]", "v14.b[14]", "v14.b[15]",
	"v15.b[0]", "v15.b[1]", "v15.b[2]", "v15.b[3]", "v15.b[4]", "v15.b[5]", "v15.b[6]", "v15.b[7]",
	"v15.b[8]", "v15.b[9]", "v15.b[10]", "v15.b[11]", "v15.b[12]", "v15.b[13]", "v15.b[14]", "v15.b[15]",
	"v16.b[0]", "v16.b[1]", "v16.b[2]", "v16.b[3]", "v16.b[4]", "v16.b[5]", "v16.b[6]", "v16.b[7]",
	"v16.b[8]", "v16.b[9]", "v16.b[10]", "v16.b[11]", "v16.b[12]", "v16.b[13]", "v16.b[14]", "v16.b[15]",
	"v17.b[0]", "v17.b[1]", "v17.b[2]", "v17.b[3]", "v17.b[4]", "v17.b[5]", "v17.b[6]", "v17.b[7]",
	"v17.b[8]", "v17.b[9]", "v17.b[10]", "v17.b[11]", "v17.b[12]", "v17.b[13]", "v17.b[14]", "v17.b[15]",
	"v18.b[0]", "v18.b[1]", "v18.b[2]", "v18.b[3]", "v18.b[4]", "v18.b[5]", "v18.b[6]", "v18.b[7]",
	"v18.b[8]", "v18.b[9]", "v18.b[10]", "v18.b[11]", "v18.b[12]", "v18.b[13]", "v18.b[14]", "v18.b[15]",
	"v19.b[0]", "v19.b[1]", "v19.b[2]", "v19.b[3]", "v19.b[4]", "v19.b[5]", "v19.b[6]", "v19.b[7]",
	"v19.b[8]", "v19.b[9]", "v19.b[10]", "v19.b[11]", "v19.b[12]", "v19.b[13]", "v19.b[14]", "v19.b[15]",
	"v20.b[0]", "v20.b[1]", "v20.b[2]", "v20.b[3]", "v20.b[4]", "v20.b[5]", "v20.b[6]", "v20.b[7]",
	"v20.b[8]", "v20.b[9]", "v20.b[10]", "v20.b[11]", "v20.b[12]", "v20.b[13]", "v20.b[14]", "v20.b[15]",
	"v21.b[0]", "v21.b[1]", "v21.b[2]", "v21.b[3]", "v21.b[4]", "v21.b[5]", "v21.b[6]", "v21.b[7]",
	"v21.b[8]", "v21.b[9]", "v21.b[10]", "v21.b[11]", "v21.b[12]", "v21.b[13]", "v21.b[14]", "v21.b[15]",
	"v22.b[0]", "v22.b[1]", "v22.b[2]", "v22.b[3]", "v22.b[4]", "v22.b[5]", "v22.b[6]", "v22.b[7]",
	"v22.b[8]", "v22.b[9]", "v22.b[10]", "v22.b[11]", "v22.b[12]", "v22.b[13]", "v22.b[14]", "v22.b[15]",
	"v23.b[0]", "v23.b[1]", "v23.b[2]", "v23.b[3]", "v23.b[4]", "v23.b[5]", "v23.b[6]", "v23.b[7]",
	"v23.b[8]", "v23.b[9]", "v23.b[10]", "v23.b[11]", "v23.b[12]", "v23.b[13]", "v23.b[14]", "v23.b[15]",
	"v24.b[0]", "v24.b[1]", "v24.b[2]", "v24.b[3]", "v24.b[4]", "v24.b[5]", "v24.b[6]", "v24.b[7]",
	"v24.b[8]", "v24.b[9]", "v24.b[10]", "v24.b[11]", "v24.b[12]", "v24.b[13]", "v24.b[14]", "v24.b[15]",
	"v25.b[0]", "v25.b[1]", "v25.b[2]", "v25.b[3]", "v25.b[4]", "v25.b[5]", "v25.b[6]", "v25.b[7]",
	"v25.b[8]", "v25.b[9]", "v25.b[10]", "v25.b[11]", "v25.b[12]", "v25.b[13]", "v25.b[14]", "v25.b[15]",
	"v26.b[0]", "v26.b[1]", "v26.b[2]", "v26.b[3]", "v26.b[4]", "v26.b[5]", "v26.b[6]", "v26.b[7]",
	"v26.b[8]", "v26.b[9]", "v26.b[10]", "v26.b[11]", "v26.b[12]", "v26.b[13]", "v26.b[14]", "v26.b[15]",
	"v27.b[0]", "v27.b[1]", "v27.b[2]", "v27.b[3]", "v27.b[4]", "v27.b[5]", "v27.b[6]", "v27.b[7]",
	"v27.b[8]", "v27.b[9]", "v27.b[10]", "v27.b[11]", "v27.b[12]", "v27.b[13]", "v27.b[14]", "v27.b[15]",
	"v28.b[0]", "v28.b[1]", "v28.b[2]", "v28.b[3]", "v28.b[4]", "v28.b[5]", "v28.b[6]", "v28.b[7]",
	"v28.b[8]", "v28.b[9]", "v28.b[10]", "v28.b[11]", "v28.b[12]", "v28.b[13]", "v28.b[14]", "v28.b[15]",
	"v29.b[0]", "v29.b[1]", "v29.b[2]", "v29.b[3]", "v29.b[4]", "v29.b[5]", "v29.b[6]", "v29.b[7]",
	"v29.b[8]", "v29.b[9]", "v29.b[10]", "v29.b[11]", "v29.b[12]", "v29.b[13]", "v29.b[14]", "v29.b[15]",
	"v30.b[0]", "v30.b[1]", "v30.b[2]", "v30.b[3]", "v30.b[4]", "v30.b[5]", "v30.b[6]", "v30.b[7]",
	"v30.b[8]", "v30.b[9]", "v30.b[10]", "v30.b[11]", "v30.b[12]", "v30.b[13]", "v30.b[14]", "v30.b[15]",
	"v31.b[0]", "v31.b[1]", "v31.b[2]", "v31.b[3]", "v31.b[4]", "v31.b[5]", "v31.b[6]", "v31.b[7]",
	"v31.b[8]", "v31.b[9]", "v31.b[10]", "v31.b[11]", "v31.b[12]", "v31.b[13]", "v31.b[14]", "v31.b[15]",
	// H vectors
	"v0.h[0]", "v0.h[1]", "v0.h[2]", "v0.h[3]", "v0.h[4]", "v0.h[5]", "v0.h[6]", "v0.h[7]",
	"v1.h[0]", "v1.h[1]", "v1.h[2]", "v1.h[3]", "v1.h[4]", "v1.h[5]", "v1.h[6]", "v1.h[7]",
	"v2.h[0]", "v2.h[1]", "v2.h[2]", "v2.h[3]", "v2.h[4]", "v2.h[5]", "v2.h[6]", "v2.h[7]",
	"v3.h[0]", "v3.h[1]", "v3.h[2]", "v3.h[3]", "v3.h[4]", "v3.h[5]", "v3.h[6]", "v3.h[7]",
	"v4.h[0]", "v4.h[1]", "v4.h[2]", "v4.h[3]", "v4.h[4]", "v4.h[5]", "v4.h[6]", "v4.h[7]",
	"v5.h[0]", "v5.h[1]", "v5.h[2]", "v5.h[3]", "v5.h[4]", "v5.h[5]", "v5.h[6]", "v5.h[7]",
	"v6.h[0]", "v6.h[1]", "v6.h[2]", "v6.h[3]", "v6.h[4]", "v6.h[5]", "v6.h[6]", "v6.h[7]",
	"v7.h[0]", "v7.h[1]", "v7.h[2]", "v7.h[3]", "v7.h[4]", "v7.h[5]", "v7.h[6]", "v7.h[7]",
	"v8.h[0]", "v8.h[1]", "v8.h[2]", "v8.h[3]", "v8.h[4]", "v8.h[5]", "v8.h[6]", "v8.h[7]",
	"v9.h[0]", "v9.h[1]", "v9.h[2]", "v9.h[3]", "v9.h[4]", "v9.h[5]", "v9.h[6]", "v9.h[7]",
	"v10.h[0]", "v10.h[1]", "v10.h[2]", "v10.h[3]", "v10.h[4]", "v10.h[5]", "v10.h[6]", "v10.h[7]",
	"v11.h[0]", "v11.h[1]", "v11.h[2]", "v11.h[3]", "v11.h[4]", "v11.h[5]", "v11.h[6]", "v11.h[7]",
	"v12.h[0]", "v12.h[1]", "v12.h[2]", "v12.h[3]", "v12.h[4]", "v12.h[5]", "v12.h[6]", "v12.h[7]",
	"v13.h[0]", "v13.h[1]", "v13.h[2]", "v13.h[3]", "v13.h[4]", "v13.h[5]", "v13.h[6]", "v13.h[7]",
	"v14.h[0]", "v14.h[1]", "v14.h[2]", "v14.h[3]", "v14.h[4]", "v14.h[5]", "v14.h[6]", "v14.h[7]",
	"v15.h[0]", "v15.h[1]", "v15.h[2]", "v15.h[3]", "v15.h[4]", "v15.h[5]", "v15.h[6]", "v15.h[7]",
	"v16.h[0]", "v16.h[1]", "v16.h[2]", "v16.h[3]", "v16.h[4]", "v16.h[5]", "v16.h[6]", "v16.h[7]",
	"v17.h[0]", "v17.h[1]", "v17.h[2]", "v17.h[3]", "v17.h[4]", "v17.h[5]", "v17.h[6]", "v17.h[7]",
	"v18.h[0]", "v18.h[1]", "v18.h[2]", "v18.h[3]", "v18.h[4]", "v18.h[5]", "v18.h[6]", "v18.h[7]",
	"v19.h[0]", "v19.h[1]", "v19.h[2]", "v19.h[3]", "v19.h[4]", "v19.h[5]", "v19.h[6]", "v19.h[7]",
	"v20.h[0]", "v20.h[1]", "v20.h[2]", "v20.h[3]", "v20.h[4]", "v20.h[5]", "v20.h[6]", "v20.h[7]",
	"v21.h[0]", "v21.h[1]", "v21.h[2]", "v21.h[3]", "v21.h[4]", "v21.h[5]", "v21.h[6]", "v21.h[7]",
	"v22.h[0]", "v22.h[1]", "v22.h[2]", "v22.h[3]", "v22.h[4]", "v22.h[5]", "v22.h[6]", "v22.h[7]",
	"v23.h[0]", "v23.h[1]", "v23.h[2]", "v23.h[3]", "v23.h[4]", "v23.h[5]", "v23.h[6]", "v23.h[7]",
	"v24.h[0]", "v24.h[1]", "v24.h[2]", "v24.h[3]", "v24.h[4]", "v24.h[5]", "v24.h[6]", "v24.h[7]",
	"v25.h[0]", "v25.h[1]", "v25.h[2]", "v25.h[3]", "v25.h[4]", "v25.h[5]", "v25.h[6]", "v25.h[7]",
	"v26.h[0]", "v26.h[1]", "v26.h[2]", "v26.h[3]", "v26.h[4]", "v26.h[5]", "v26.h[6]", "v26.h[7]",
	"v27.h[0]", "v27.h[1]", "v27.h[2]", "v27.h[3]", "v27.h[4]", "v27.h[5]", "v27.h[6]", "v27.h[7]",
	"v28.h[0]", "v28.h[1]", "v28.h[2]", "v28.h[3]", "v28.h[4]", "v28.h[5]", "v28.h[6]", "v28.h[7]",
	"v29.h[0]", "v29.h[1]", "v29.h[2]", "v29.h[3]", "v29.h[4]", "v29.h[5]", "v29.h[6]", "v29.h[7]",
	"v30.h[0]", "v30.h[1]", "v30.h[2]", "v30.h[3]", "v30.h[4]", "v30.h[5]", "v30.h[6]", "v30.h[7]",
	"v31.h[0]", "v31.h[1]", "v31.h[2]", "v31.h[3]", "v31.h[4]", "v31.h[5]", "v31.h[6]", "v31.h[7]",
	// S vectors
	"v0.s[0]", "v0.s[1]", "v0.s[2]", "v0.s[3]", "v1.s[0]", "v1.s[1]", "v1.s[2]", "v1.s[3]",
	"v2.s[0]", "v2.s[1]", "v2.s[2]", "v2.s[3]", "v3.s[0]", "v3.s[1]", "v3.s[2]", "v3.s[3]",
	"v4.s[0]", "v4.s[1]", "v4.s[2]", "v4.s[3]", "v5.s[0]", "v5.s[1]", "v5.s[2]", "v5.s[3]",
	"v6.s[0]", "v6.s[1]", "v6.s[2]", "v6.s[3]", "v7.s[0]", "v7.s[1]", "v7.s[2]", "v7.s[3]",
	"v8.s[0]", "v8.s[1]", "v8.s[2]", "v8.s[3]", "v9.s[0]", "v9.s[1]", "v9.s[2]", "v9.s[3]",
	"v10.s[0]", "v10.s[1]", "v10.s[2]", "v10.s[3]", "v11.s[0]", "v11.s[1]", "v11.s[2]", "v11.s[3]",
	"v12.s[0]", "v12.s[1]", "v12.s[2]", "v12.s[3]", "v13.s[0]", "v13.s[1]", "v13.s[2]", "v13.s[3]",
	"v14.s[0]", "v14.s[1]", "v14.s[2]", "v14.s[3]", "v15.s[0]", "v15.s[1]", "v15.s[2]", "v15.s[3]",
	"v16.s[0]", "v16.s[1]", "v16.s[2]", "v16.s[3]", "v17.s[0]", "v17.s[1]", "v17.s[2]", "v17.s[3]",
	"v18.s[0]", "v18.s[1]", "v18.s[2]", "v18.s[3]", "v19.s[0]", "v19.s[1]", "v19.s[2]", "v19.s[3]",
	"v20.s[0]", "v20.s[1]", "v20.s[2]", "v20.s[3]", "v21.s[0]", "v21.s[1]", "v21.s[2]", "v21.s[3]",
	"v22.s[0]", "v22.s[1]", "v22.s[2]", "v22.s[3]", "v23.s[0]", "v23.s[1]", "v23.s[2]", "v23.s[3]",
	"v24.s[0]", "v24.s[1]", "v24.s[2]", "v24.s[3]", "v25.s[0]", "v25.s[1]", "v25.s[2]", "v25.s[3]",
	"v26.s[0]", "v26.s[1]", "v26.s[2]", "v26.s[3]", "v27.s[0]", "v27.s[1]", "v27.s[2]", "v27.s[3]",
	"v28.s[0]", "v28.s[1]", "v28.s[2]", "v28.s[3]", "v29.s[0]", "v29.s[1]", "v29.s[2]", "v29.s[3]",
	"v30.s[0]", "v30.s[1]", "v30.s[2]", "v30.s[3]", "v31.s[0]", "v31.s[1]", "v31.s[2]", "v31.s[3]",
	// D vectors
	"v0.d[0]", "v0.d[1]", "v1.d[0]", "v1.d[1]", "v2.d[0]", "v2.d[1]", "v3.d[0]", "v3.d[1]",
	"v4.d[0]", "v4.d[1]", "v5.d[0]", "v5.d[1]", "v6.d[0]", "v6.d[1]", "v7.d[0]", "v7.d[1]",
	"v8.d[0]", "v8.d[1]", "v9.d[0]", "v9.d[1]", "v10.d[0]", "v10.d[1]", "v11.d[0]", "v11.d[1]",
	"v12.d[0]", "v12.d[1]", "v13.d[0]", "v13.d[1]", "v14.d[0]", "v14.d[1]", "v15.d[0]", "v15.d[1]",
	"v16.d[0]", "v16.d[1]", "v17.d[0]", "v17.d[1]", "v18.d[0]", "v18.d[1]", "v19.d[0]", "v19.d[1]",
	"v20.d[0]", "v20.d[1]", "v21.d[0]", "v21.d[1]", "v22.d[0]", "v22.d[1]", "v23.d[0]", "v23.d[1]",
	"v24.d[0]", "v24.d[1]", "v25.d[0]", "v25.d[1]", "v26.d[0]", "v26.d[1]", "v27.d[0]", "v27.d[1]",
	"v28.d[0]", "v28.d[1]", "v29.d[0]", "v29.d[1]", "v30.d[0]", "v30.d[1]", "v31.d[0]", "v31.d[1]",
	// SVE
	"z0", "z1", "z2", "z3", "z4", "z5", "z6", "z7",
	"z8", "z9", "z10", "z11", "z12", "z13", "z14", "z15",
	"z16", "z17", "z18", "z19", "z20", "z21", "z22", "z23",
	"z24", "z25", "z26", "z27", "z28", "z29", "z30", "z31", "z31",
	/* scalable predicate registers */
	"p0", "p1", "p2", "p3", "p4", "p5", "p6", "p7",
	"p8", "p9", "p10", "p11", "p12", "p13", "p14", "p15",
	"p16", "p17", "p18", "p19", "p20", "p21", "p22", "p23",
	"p24", "p25", "p26", "p27", "p28", "p29", "p30", "p31",
	/* prefetch operations (TODO: remove these as registers) */
	"pldl1keep", "pldl1strm", "pldl2keep", "pldl2strm",
	"pldl3keep", "pldl3strm", "#0x6", "#0x7",
	"plil1keep", "plil1strm", "plil2keep", "plil2strm",
	"plil3keep", "plil3strm", "#0xe", "#0xf",
	"pstl1keep", "pstl1strm", "pstl2keep", "pstl2strm",
	"pstl3keep", "pstl3strm", "#0x16", "#0x17",
	"#0x18", "#0x19", "#0x1a", "#0x1b",
	"#0x1c", "#0x1d", "#0x1e", "#0x1f",
	"END",
}

func (r register) String() string {
	return regString[r]
}

func (r register) Size() int {
	//Comparison done in order of likelyhood to occur
	if (r >= REG_X0 && r <= REG_SP) || (r >= REG_D0 && r <= REG_D31) {
		return 8
	} else if (r >= REG_W0 && r <= REG_WSP) || (r >= REG_S0 && r <= REG_S31) {
		return 4
	} else if r >= REG_B0 && r <= REG_B31 {
		return 1
	} else if r >= REG_H0 && r <= REG_H31 {
		return 2
	} else if (r >= REG_Q0 && r <= REG_Q31) || (r >= REG_V0 && r <= REG_V31) {
		return 16
	} else if r >= REG_V0_B0 && r <= REG_V31_B15 {
		return 1
	} else if r >= REG_V0_H0 && r <= REG_V31_H7 {
		return 2
	} else if r >= REG_V0_S0 && r <= REG_V31_S3 {
		return 4
	} else if r >= REG_V0_D0 && r <= REG_V31_D1 {
		return 8
	}
	return 0
}
