/* GENERATED FILE */
#pragma once
// 100 HasXXX() functions used by decode:
#define HasFeature(feat) (ctx->decode_features[(feat) / 64] & ((uint64_t)1 << ((feat) % 64)))

#define HasAES() HasFeature(ARCH_FEATURE_AES)
#define HasAdvSIMD() HasFeature(ARCH_FEATURE_AdvSIMD)
#define HasBF16() HasFeature(ARCH_FEATURE_BF16)
#define HasBTI() HasFeature(ARCH_FEATURE_BTI)
#define HasCHK() HasFeature(ARCH_FEATURE_CHK)
#define HasCLRBHB() HasFeature(ARCH_FEATURE_CLRBHB)
#define HasCMPBR() HasFeature(ARCH_FEATURE_CMPBR)
#define HasCPA() HasFeature(ARCH_FEATURE_CPA)
#define HasCRC32() HasFeature(ARCH_FEATURE_CRC32)
#define HasCSSC() HasFeature(ARCH_FEATURE_CSSC)
#define HasD128() HasFeature(ARCH_FEATURE_D128)
#define HasDGH() HasFeature(ARCH_FEATURE_DGH)
#define HasDotProd() HasFeature(ARCH_FEATURE_DotProd)
#define HasF32MM() HasFeature(ARCH_FEATURE_F32MM)
#define HasF64MM() HasFeature(ARCH_FEATURE_F64MM)
#define HasF8F16MM() HasFeature(ARCH_FEATURE_F8F16MM)
#define HasF8F32MM() HasFeature(ARCH_FEATURE_F8F32MM)
#define HasFAMINMAX() HasFeature(ARCH_FEATURE_FAMINMAX)
#define HasFCMA() HasFeature(ARCH_FEATURE_FCMA)
#define HasFHM() HasFeature(ARCH_FEATURE_FHM)
#define HasFP() HasFeature(ARCH_FEATURE_FP)
#define HasFP16() HasFeature(ARCH_FEATURE_FP16)
#define HasFP8() HasFeature(ARCH_FEATURE_FP8)
#define HasFP8DOT2() HasFeature(ARCH_FEATURE_FP8DOT2)
#define HasFP8DOT4() HasFeature(ARCH_FEATURE_FP8DOT4)
#define HasFP8FMA() HasFeature(ARCH_FEATURE_FP8FMA)
#define HasFPRCVT() HasFeature(ARCH_FEATURE_FPRCVT)
#define HasFRINTTS() HasFeature(ARCH_FEATURE_FRINTTS)
#define HasFlagM() HasFeature(ARCH_FEATURE_FlagM)
#define HasFlagM2() HasFeature(ARCH_FEATURE_FlagM2)
#define HasGCS() HasFeature(ARCH_FEATURE_GCS)
#define HasHBC() HasFeature(ARCH_FEATURE_HBC)
#define HasI8MM() HasFeature(ARCH_FEATURE_I8MM)
#define HasJSCVT() HasFeature(ARCH_FEATURE_JSCVT)
#define HasLOR() HasFeature(ARCH_FEATURE_LOR)
#define HasLRCPC() HasFeature(ARCH_FEATURE_LRCPC)
#define HasLRCPC2() HasFeature(ARCH_FEATURE_LRCPC2)
#define HasLRCPC3() HasFeature(ARCH_FEATURE_LRCPC3)
#define HasLS64() HasFeature(ARCH_FEATURE_LS64)
#define HasLS64_ACCDATA() HasFeature(ARCH_FEATURE_LS64_ACCDATA)
#define HasLS64_V() HasFeature(ARCH_FEATURE_LS64_V)
#define HasLSE() HasFeature(ARCH_FEATURE_LSE)
#define HasLSE128() HasFeature(ARCH_FEATURE_LSE128)
#define HasLSFE() HasFeature(ARCH_FEATURE_LSFE)
#define HasLSUI() HasFeature(ARCH_FEATURE_LSUI)
#define HasLUT() HasFeature(ARCH_FEATURE_LUT)
#define HasMOPS() HasFeature(ARCH_FEATURE_MOPS)
#define HasMTE() HasFeature(ARCH_FEATURE_MTE)
#define HasMTE2() HasFeature(ARCH_FEATURE_MTE2)
#define HasPAuth() HasFeature(ARCH_FEATURE_PAuth)
#define HasPAuth_LR() HasFeature(ARCH_FEATURE_PAuth_LR)
#define HasPCDPHINT() HasFeature(ARCH_FEATURE_PCDPHINT)
#define HasRAS() HasFeature(ARCH_FEATURE_RAS)
#define HasRDM() HasFeature(ARCH_FEATURE_RDM)
#define HasRPRFM() HasFeature(ARCH_FEATURE_RPRFM)
#define HasSB() HasFeature(ARCH_FEATURE_SB)
#define HasSHA1() HasFeature(ARCH_FEATURE_SHA1)
#define HasSHA256() HasFeature(ARCH_FEATURE_SHA256)
#define HasSHA3() HasFeature(ARCH_FEATURE_SHA3)
#define HasSHA512() HasFeature(ARCH_FEATURE_SHA512)
#define HasSM3() HasFeature(ARCH_FEATURE_SM3)
#define HasSM4() HasFeature(ARCH_FEATURE_SM4)
#define HasSME() HasFeature(ARCH_FEATURE_SME)
#define HasSME2() HasFeature(ARCH_FEATURE_SME2)
#define HasSME2p1() HasFeature(ARCH_FEATURE_SME2p1)
#define HasSME2p2() HasFeature(ARCH_FEATURE_SME2p2)
#define HasSME_B16B16() HasFeature(ARCH_FEATURE_SME_B16B16)
#define HasSME_F16F16() HasFeature(ARCH_FEATURE_SME_F16F16)
#define HasSME_F64F64() HasFeature(ARCH_FEATURE_SME_F64F64)
#define HasSME_F8F16() HasFeature(ARCH_FEATURE_SME_F8F16)
#define HasSME_F8F32() HasFeature(ARCH_FEATURE_SME_F8F32)
#define HasSME_I16I64() HasFeature(ARCH_FEATURE_SME_I16I64)
#define HasSME_LUTv2() HasFeature(ARCH_FEATURE_SME_LUTv2)
#define HasSME_MOP4() HasFeature(ARCH_FEATURE_SME_MOP4)
#define HasSME_TMOP() HasFeature(ARCH_FEATURE_SME_TMOP)
#define HasSPE() HasFeature(ARCH_FEATURE_SPE)
#define HasSSVE_FEXPA() HasFeature(ARCH_FEATURE_SSVE_FEXPA)
#define HasSSVE_FP8DOT2() HasFeature(ARCH_FEATURE_SSVE_FP8DOT2)
#define HasSSVE_FP8DOT4() HasFeature(ARCH_FEATURE_SSVE_FP8DOT4)
#define HasSSVE_FP8FMA() HasFeature(ARCH_FEATURE_SSVE_FP8FMA)
#define HasSVE() HasFeature(ARCH_FEATURE_SVE)
#define HasSVE2() HasFeature(ARCH_FEATURE_SVE2)
#define HasSVE2p1() HasFeature(ARCH_FEATURE_SVE2p1)
#define HasSVE2p2() HasFeature(ARCH_FEATURE_SVE2p2)
#define HasSVE_AES() HasFeature(ARCH_FEATURE_SVE_AES)
#define HasSVE_AES2() HasFeature(ARCH_FEATURE_SVE_AES2)
#define HasSVE_B16B16() HasFeature(ARCH_FEATURE_SVE_B16B16)
#define HasSVE_BFSCALE() HasFeature(ARCH_FEATURE_SVE_BFSCALE)
#define HasSVE_BitPerm() HasFeature(ARCH_FEATURE_SVE_BitPerm)
#define HasSVE_F16F32MM() HasFeature(ARCH_FEATURE_SVE_F16F32MM)
#define HasSVE_PMULL128() HasFeature(ARCH_FEATURE_SVE_PMULL128)
#define HasSVE_SHA3() HasFeature(ARCH_FEATURE_SVE_SHA3)
#define HasSVE_SM4() HasFeature(ARCH_FEATURE_SVE_SM4)
#define HasSYSINSTR128() HasFeature(ARCH_FEATURE_SYSINSTR128)
#define HasSYSREG128() HasFeature(ARCH_FEATURE_SYSREG128)
#define HasTHE() HasFeature(ARCH_FEATURE_THE)
#define HasTME() HasFeature(ARCH_FEATURE_TME)
#define HasTRF() HasFeature(ARCH_FEATURE_TRF)
#define HasWFxT() HasFeature(ARCH_FEATURE_WFxT)
#define HasXS() HasFeature(ARCH_FEATURE_XS)

// 108 HaveXXX()/IsImplemented(FEAT_XXX) functions used by pcode:
#define HaveFeature(feat) (ctx->pcode_features[(feat) / 64] & ((uint64_t)1 << ((feat) % 64)))
#define HaveAES() HaveFeature(ARCH_FEATURE_AES)
#define HaveAdvSIMD() HaveFeature(ARCH_FEATURE_AdvSIMD)
#define HaveBF16() HaveFeature(ARCH_FEATURE_BF16)
#define HaveBTI() HaveFeature(ARCH_FEATURE_BTI)
#define HaveCHK() HaveFeature(ARCH_FEATURE_CHK)
#define HaveCLRBHB() HaveFeature(ARCH_FEATURE_CLRBHB)
#define HaveCMPBR() HaveFeature(ARCH_FEATURE_CMPBR)
#define HaveCPA() HaveFeature(ARCH_FEATURE_CPA)
#define HaveCRC32() HaveFeature(ARCH_FEATURE_CRC32)
#define HaveCSSC() HaveFeature(ARCH_FEATURE_CSSC)
#define HaveD128() HaveFeature(ARCH_FEATURE_D128)
#define HaveDGH() HaveFeature(ARCH_FEATURE_DGH)
#define HaveDIT() HaveFeature(ARCH_FEATURE_DIT)
#define HaveDotProd() HaveFeature(ARCH_FEATURE_DotProd)
#define HaveEBEP() HaveFeature(ARCH_FEATURE_EBEP)
#define HaveF32MM() HaveFeature(ARCH_FEATURE_F32MM)
#define HaveF64MM() HaveFeature(ARCH_FEATURE_F64MM)
#define HaveF8F16MM() HaveFeature(ARCH_FEATURE_F8F16MM)
#define HaveF8F32MM() HaveFeature(ARCH_FEATURE_F8F32MM)
#define HaveFAMINMAX() HaveFeature(ARCH_FEATURE_FAMINMAX)
#define HaveFCMA() HaveFeature(ARCH_FEATURE_FCMA)
#define HaveFHM() HaveFeature(ARCH_FEATURE_FHM)
#define HaveFP() HaveFeature(ARCH_FEATURE_FP)
#define HaveFP16() HaveFeature(ARCH_FEATURE_FP16)
#define HaveFP8() HaveFeature(ARCH_FEATURE_FP8)
#define HaveFP8DOT2() HaveFeature(ARCH_FEATURE_FP8DOT2)
#define HaveFP8DOT4() HaveFeature(ARCH_FEATURE_FP8DOT4)
#define HaveFP8FMA() HaveFeature(ARCH_FEATURE_FP8FMA)
#define HaveFPRCVT() HaveFeature(ARCH_FEATURE_FPRCVT)
#define HaveFRINTTS() HaveFeature(ARCH_FEATURE_FRINTTS)
#define HaveFlagM() HaveFeature(ARCH_FEATURE_FlagM)
#define HaveFlagM2() HaveFeature(ARCH_FEATURE_FlagM2)
#define HaveGCS() HaveFeature(ARCH_FEATURE_GCS)
#define HaveHBC() HaveFeature(ARCH_FEATURE_HBC)
#define HaveI8MM() HaveFeature(ARCH_FEATURE_I8MM)
#define HaveJSCVT() HaveFeature(ARCH_FEATURE_JSCVT)
#define HaveLOR() HaveFeature(ARCH_FEATURE_LOR)
#define HaveLRCPC() HaveFeature(ARCH_FEATURE_LRCPC)
#define HaveLRCPC2() HaveFeature(ARCH_FEATURE_LRCPC2)
#define HaveLRCPC3() HaveFeature(ARCH_FEATURE_LRCPC3)
#define HaveLS64() HaveFeature(ARCH_FEATURE_LS64)
#define HaveLS64_ACCDATA() HaveFeature(ARCH_FEATURE_LS64_ACCDATA)
#define HaveLS64_V() HaveFeature(ARCH_FEATURE_LS64_V)
#define HaveLSE() HaveFeature(ARCH_FEATURE_LSE)
#define HaveLSE128() HaveFeature(ARCH_FEATURE_LSE128)
#define HaveLSFE() HaveFeature(ARCH_FEATURE_LSFE)
#define HaveLSUI() HaveFeature(ARCH_FEATURE_LSUI)
#define HaveLUT() HaveFeature(ARCH_FEATURE_LUT)
#define HaveMOPS() HaveFeature(ARCH_FEATURE_MOPS)
#define HaveMTE() HaveFeature(ARCH_FEATURE_MTE)
#define HaveMTE2() HaveFeature(ARCH_FEATURE_MTE2)
#define HaveNMI() HaveFeature(ARCH_FEATURE_NMI)
#define HavePAN() HaveFeature(ARCH_FEATURE_PAN)
#define HavePAuth() HaveFeature(ARCH_FEATURE_PAuth)
#define HavePAuth_LR() HaveFeature(ARCH_FEATURE_PAuth_LR)
#define HavePCDPHINT() HaveFeature(ARCH_FEATURE_PCDPHINT)
#define HavePMULL() HaveFeature(ARCH_FEATURE_PMULL)
#define HaveRAS() HaveFeature(ARCH_FEATURE_RAS)
#define HaveRDM() HaveFeature(ARCH_FEATURE_RDM)
#define HaveRPRFM() HaveFeature(ARCH_FEATURE_RPRFM)
#define HaveSB() HaveFeature(ARCH_FEATURE_SB)
#define HaveSHA1() HaveFeature(ARCH_FEATURE_SHA1)
#define HaveSHA256() HaveFeature(ARCH_FEATURE_SHA256)
#define HaveSHA3() HaveFeature(ARCH_FEATURE_SHA3)
#define HaveSHA512() HaveFeature(ARCH_FEATURE_SHA512)
#define HaveSM3() HaveFeature(ARCH_FEATURE_SM3)
#define HaveSM4() HaveFeature(ARCH_FEATURE_SM4)
#define HaveSME() HaveFeature(ARCH_FEATURE_SME)
#define HaveSME2() HaveFeature(ARCH_FEATURE_SME2)
#define HaveSME2p1() HaveFeature(ARCH_FEATURE_SME2p1)
#define HaveSME2p2() HaveFeature(ARCH_FEATURE_SME2p2)
#define HaveSME_B16B16() HaveFeature(ARCH_FEATURE_SME_B16B16)
#define HaveSME_F16F16() HaveFeature(ARCH_FEATURE_SME_F16F16)
#define HaveSME_F64F64() HaveFeature(ARCH_FEATURE_SME_F64F64)
#define HaveSME_F8F16() HaveFeature(ARCH_FEATURE_SME_F8F16)
#define HaveSME_F8F32() HaveFeature(ARCH_FEATURE_SME_F8F32)
#define HaveSME_I16I64() HaveFeature(ARCH_FEATURE_SME_I16I64)
#define HaveSME_LUTv2() HaveFeature(ARCH_FEATURE_SME_LUTv2)
#define HaveSME_MOP4() HaveFeature(ARCH_FEATURE_SME_MOP4)
#define HaveSME_TMOP() HaveFeature(ARCH_FEATURE_SME_TMOP)
#define HaveSPE() HaveFeature(ARCH_FEATURE_SPE)
#define HaveSSBS() HaveFeature(ARCH_FEATURE_SSBS)
#define HaveSSVE_FEXPA() HaveFeature(ARCH_FEATURE_SSVE_FEXPA)
#define HaveSVE() HaveFeature(ARCH_FEATURE_SVE)
#define HaveSVE2() HaveFeature(ARCH_FEATURE_SVE2)
#define HaveSVE2FP8DOT2() HaveFeature(ARCH_FEATURE_SVE2FP8DOT2)
#define HaveSVE2FP8DOT4() HaveFeature(ARCH_FEATURE_SVE2FP8DOT4)
#define HaveSVE2FP8FMA() HaveFeature(ARCH_FEATURE_SVE2FP8FMA)
#define HaveSVE2p1() HaveFeature(ARCH_FEATURE_SVE2p1)
#define HaveSVE2p2() HaveFeature(ARCH_FEATURE_SVE2p2)
#define HaveSVE_AES() HaveFeature(ARCH_FEATURE_SVE_AES)
#define HaveSVE_AES2() HaveFeature(ARCH_FEATURE_SVE_AES2)
#define HaveSVE_B16B16() HaveFeature(ARCH_FEATURE_SVE_B16B16)
#define HaveSVE_BFSCALE() HaveFeature(ARCH_FEATURE_SVE_BFSCALE)
#define HaveSVE_BitPerm() HaveFeature(ARCH_FEATURE_SVE_BitPerm)
#define HaveSVE_F16F32MM() HaveFeature(ARCH_FEATURE_SVE_F16F32MM)
#define HaveSVE_PMULL128() HaveFeature(ARCH_FEATURE_SVE_PMULL128)
#define HaveSVE_SHA3() HaveFeature(ARCH_FEATURE_SVE_SHA3)
#define HaveSVE_SM4() HaveFeature(ARCH_FEATURE_SVE_SM4)
#define HaveSYSINSTR128() HaveFeature(ARCH_FEATURE_SYSINSTR128)
#define HaveSYSREG128() HaveFeature(ARCH_FEATURE_SYSREG128)
#define HaveTHE() HaveFeature(ARCH_FEATURE_THE)
#define HaveTME() HaveFeature(ARCH_FEATURE_TME)
#define HaveTRF() HaveFeature(ARCH_FEATURE_TRF)
#define HaveUAO() HaveFeature(ARCH_FEATURE_UAO)
#define HaveVHE() HaveFeature(ARCH_FEATURE_VHE)
#define HaveWFxT() HaveFeature(ARCH_FEATURE_WFxT)
#define HaveXS() HaveFeature(ARCH_FEATURE_XS)

#define ARCH_FEATURE_AES                0  // Referenced by decode and pcode
#define ARCH_FEATURE_AdvSIMD            1  // Referenced by decode and pcode
#define ARCH_FEATURE_BF16               2  // Referenced by decode and pcode
#define ARCH_FEATURE_BTI                3  // Referenced by decode and pcode
#define ARCH_FEATURE_CHK                4  // Referenced by decode and pcode
#define ARCH_FEATURE_CLRBHB             5  // Referenced by decode and pcode
#define ARCH_FEATURE_CMPBR              6  // Referenced by decode and pcode
#define ARCH_FEATURE_CPA                7  // Referenced by decode and pcode
#define ARCH_FEATURE_CRC32              8  // Referenced by decode and pcode
#define ARCH_FEATURE_CSSC               9  // Referenced by decode and pcode
#define ARCH_FEATURE_D128              10  // Referenced by decode and pcode
#define ARCH_FEATURE_DGH               11  // Referenced by decode and pcode
#define ARCH_FEATURE_DIT               12  // Referenced by pcode
#define ARCH_FEATURE_DotProd           13  // Referenced by decode and pcode
#define ARCH_FEATURE_EBEP              14  // Referenced by pcode
#define ARCH_FEATURE_F32MM             15  // Referenced by decode and pcode
#define ARCH_FEATURE_F64MM             16  // Referenced by decode and pcode
#define ARCH_FEATURE_F8F16MM           17  // Referenced by decode and pcode
#define ARCH_FEATURE_F8F32MM           18  // Referenced by decode and pcode
#define ARCH_FEATURE_FAMINMAX          19  // Referenced by decode and pcode
#define ARCH_FEATURE_FCMA              20  // Referenced by decode and pcode
#define ARCH_FEATURE_FHM               21  // Referenced by decode and pcode
#define ARCH_FEATURE_FP                22  // Referenced by decode and pcode
#define ARCH_FEATURE_FP16              23  // Referenced by decode and pcode
#define ARCH_FEATURE_FP8               24  // Referenced by decode and pcode
#define ARCH_FEATURE_FP8DOT2           25  // Referenced by decode and pcode
#define ARCH_FEATURE_FP8DOT4           26  // Referenced by decode and pcode
#define ARCH_FEATURE_FP8FMA            27  // Referenced by decode and pcode
#define ARCH_FEATURE_FPRCVT            28  // Referenced by decode and pcode
#define ARCH_FEATURE_FRINTTS           29  // Referenced by decode and pcode
#define ARCH_FEATURE_FlagM             30  // Referenced by decode and pcode
#define ARCH_FEATURE_FlagM2            31  // Referenced by decode and pcode
#define ARCH_FEATURE_GCS               32  // Referenced by decode and pcode
#define ARCH_FEATURE_HBC               33  // Referenced by decode and pcode
#define ARCH_FEATURE_I8MM              34  // Referenced by decode and pcode
#define ARCH_FEATURE_JSCVT             35  // Referenced by decode and pcode
#define ARCH_FEATURE_LOR               36  // Referenced by decode and pcode
#define ARCH_FEATURE_LRCPC             37  // Referenced by decode and pcode
#define ARCH_FEATURE_LRCPC2            38  // Referenced by decode and pcode
#define ARCH_FEATURE_LRCPC3            39  // Referenced by decode and pcode
#define ARCH_FEATURE_LS64              40  // Referenced by decode and pcode
#define ARCH_FEATURE_LS64_ACCDATA      41  // Referenced by decode and pcode
#define ARCH_FEATURE_LS64_V            42  // Referenced by decode and pcode
#define ARCH_FEATURE_LSE               43  // Referenced by decode and pcode
#define ARCH_FEATURE_LSE128            44  // Referenced by decode and pcode
#define ARCH_FEATURE_LSFE              45  // Referenced by decode and pcode
#define ARCH_FEATURE_LSUI              46  // Referenced by decode and pcode
#define ARCH_FEATURE_LUT               47  // Referenced by decode and pcode
#define ARCH_FEATURE_MOPS              48  // Referenced by decode and pcode
#define ARCH_FEATURE_MTE               49  // Referenced by decode and pcode
#define ARCH_FEATURE_MTE2              50  // Referenced by decode and pcode
#define ARCH_FEATURE_NMI               51  // Referenced by pcode
#define ARCH_FEATURE_PAN               52  // Referenced by pcode
#define ARCH_FEATURE_PAuth             53  // Referenced by decode and pcode
#define ARCH_FEATURE_PAuth_LR          54  // Referenced by decode and pcode
#define ARCH_FEATURE_PCDPHINT          55  // Referenced by decode and pcode
#define ARCH_FEATURE_PMULL             56  // Referenced by pcode
#define ARCH_FEATURE_RAS               57  // Referenced by decode and pcode
#define ARCH_FEATURE_RDM               58  // Referenced by decode and pcode
#define ARCH_FEATURE_RPRFM             59  // Referenced by decode and pcode
#define ARCH_FEATURE_SB                60  // Referenced by decode and pcode
#define ARCH_FEATURE_SHA1              61  // Referenced by decode and pcode
#define ARCH_FEATURE_SHA256            62  // Referenced by decode and pcode
#define ARCH_FEATURE_SHA3              63  // Referenced by decode and pcode
#define ARCH_FEATURE_SHA512            64  // Referenced by decode and pcode
#define ARCH_FEATURE_SM3               65  // Referenced by decode and pcode
#define ARCH_FEATURE_SM4               66  // Referenced by decode and pcode
#define ARCH_FEATURE_SME               67  // Referenced by decode and pcode
#define ARCH_FEATURE_SME2              68  // Referenced by decode and pcode
#define ARCH_FEATURE_SME2p1            69  // Referenced by decode and pcode
#define ARCH_FEATURE_SME2p2            70  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_B16B16        71  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_F16F16        72  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_F64F64        73  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_F8F16         74  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_F8F32         75  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_I16I64        76  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_LUTv2         77  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_MOP4          78  // Referenced by decode and pcode
#define ARCH_FEATURE_SME_TMOP          79  // Referenced by decode and pcode
#define ARCH_FEATURE_SPE               80  // Referenced by decode and pcode
#define ARCH_FEATURE_SSBS              81  // Referenced by pcode
#define ARCH_FEATURE_SSVE_FEXPA        82  // Referenced by decode and pcode
#define ARCH_FEATURE_SSVE_FP8DOT2      83  // Referenced by decode
#define ARCH_FEATURE_SSVE_FP8DOT4      84  // Referenced by decode
#define ARCH_FEATURE_SSVE_FP8FMA       85  // Referenced by decode
#define ARCH_FEATURE_SVE               86  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE2              87  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE2FP8DOT2       88  // Referenced by pcode
#define ARCH_FEATURE_SVE2FP8DOT4       89  // Referenced by pcode
#define ARCH_FEATURE_SVE2FP8FMA        90  // Referenced by pcode
#define ARCH_FEATURE_SVE2p1            91  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE2p2            92  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_AES           93  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_AES2          94  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_B16B16        95  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_BFSCALE       96  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_BitPerm       97  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_F16F32MM      98  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_PMULL128      99  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_SHA3         100  // Referenced by decode and pcode
#define ARCH_FEATURE_SVE_SM4          101  // Referenced by decode and pcode
#define ARCH_FEATURE_SYSINSTR128      102  // Referenced by decode and pcode
#define ARCH_FEATURE_SYSREG128        103  // Referenced by decode and pcode
#define ARCH_FEATURE_THE              104  // Referenced by decode and pcode
#define ARCH_FEATURE_TME              105  // Referenced by decode and pcode
#define ARCH_FEATURE_TRF              106  // Referenced by decode and pcode
#define ARCH_FEATURE_UAO              107  // Referenced by pcode
#define ARCH_FEATURE_VHE              108  // Referenced by pcode
#define ARCH_FEATURE_WFxT             109  // Referenced by decode and pcode
#define ARCH_FEATURE_XS               110  // Referenced by decode and pcode

#define ARCH_FEATURES_ENABLE_ALL(feature_field)  do { \
	    _Static_assert(sizeof(feature_field) * 8 > ARCH_FEATURE_XS, "Feature field too small for largest feature value"); \
	    memset((feature_field), 0xFF, sizeof(feature_field)); \
	} while(0)
	   
