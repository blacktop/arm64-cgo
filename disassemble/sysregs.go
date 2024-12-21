package disassemble

type systemReg uint32

const (
	SYSREG_NONE          systemReg = 32769
	REG_OSDTRRX_EL1      systemReg = 32770
	REG_DBGBVR0_EL1      systemReg = 32772
	REG_DBGBCR0_EL1      systemReg = 32773
	REG_DBGWVR0_EL1      systemReg = 32774
	REG_DBGWCR0_EL1      systemReg = 32775
	REG_DBGBVR1_EL1      systemReg = 32780
	REG_DBGBCR1_EL1      systemReg = 32781
	REG_DBGWVR1_EL1      systemReg = 32782
	REG_DBGWCR1_EL1      systemReg = 32783
	REG_MDCCINT_EL1      systemReg = 32784
	REG_MDSCR_EL1        systemReg = 32786
	REG_DBGBVR2_EL1      systemReg = 32788
	REG_DBGBCR2_EL1      systemReg = 32789
	REG_DBGWVR2_EL1      systemReg = 32790
	REG_DBGWCR2_EL1      systemReg = 32791
	REG_OSDTRTX_EL1      systemReg = 32794
	REG_DBGBVR3_EL1      systemReg = 32796
	REG_DBGBCR3_EL1      systemReg = 32797
	REG_DBGWVR3_EL1      systemReg = 32798
	REG_DBGWCR3_EL1      systemReg = 32799
	REG_DBGBVR4_EL1      systemReg = 32804
	REG_DBGBCR4_EL1      systemReg = 32805
	REG_DBGWVR4_EL1      systemReg = 32806
	REG_DBGWCR4_EL1      systemReg = 32807
	REG_DBGBVR5_EL1      systemReg = 32812
	REG_DBGBCR5_EL1      systemReg = 32813
	REG_DBGWVR5_EL1      systemReg = 32814
	REG_DBGWCR5_EL1      systemReg = 32815
	REG_OSECCR_EL1       systemReg = 32818
	REG_DBGBVR6_EL1      systemReg = 32820
	REG_DBGBCR6_EL1      systemReg = 32821
	REG_DBGWVR6_EL1      systemReg = 32822
	REG_DBGWCR6_EL1      systemReg = 32823
	REG_DBGBVR7_EL1      systemReg = 32828
	REG_DBGBCR7_EL1      systemReg = 32829
	REG_DBGWVR7_EL1      systemReg = 32830
	REG_DBGWCR7_EL1      systemReg = 32831
	REG_DBGBVR8_EL1      systemReg = 32836
	REG_DBGBCR8_EL1      systemReg = 32837
	REG_DBGWVR8_EL1      systemReg = 32838
	REG_DBGWCR8_EL1      systemReg = 32839
	REG_DBGBVR9_EL1      systemReg = 32844
	REG_DBGBCR9_EL1      systemReg = 32845
	REG_DBGWVR9_EL1      systemReg = 32846
	REG_DBGWCR9_EL1      systemReg = 32847
	REG_DBGBVR10_EL1     systemReg = 32852
	REG_DBGBCR10_EL1     systemReg = 32853
	REG_DBGWVR10_EL1     systemReg = 32854
	REG_DBGWCR10_EL1     systemReg = 32855
	REG_DBGBVR11_EL1     systemReg = 32860
	REG_DBGBCR11_EL1     systemReg = 32861
	REG_DBGWVR11_EL1     systemReg = 32862
	REG_DBGWCR11_EL1     systemReg = 32863
	REG_DBGBVR12_EL1     systemReg = 32868
	REG_DBGBCR12_EL1     systemReg = 32869
	REG_DBGWVR12_EL1     systemReg = 32870
	REG_DBGWCR12_EL1     systemReg = 32871
	REG_DBGBVR13_EL1     systemReg = 32876
	REG_DBGBCR13_EL1     systemReg = 32877
	REG_DBGWVR13_EL1     systemReg = 32878
	REG_DBGWCR13_EL1     systemReg = 32879
	REG_DBGBVR14_EL1     systemReg = 32884
	REG_DBGBCR14_EL1     systemReg = 32885
	REG_DBGWVR14_EL1     systemReg = 32886
	REG_DBGWCR14_EL1     systemReg = 32887
	REG_DBGBVR15_EL1     systemReg = 32892
	REG_DBGBCR15_EL1     systemReg = 32893
	REG_DBGWVR15_EL1     systemReg = 32894
	REG_DBGWCR15_EL1     systemReg = 32895
	REG_OSLAR_EL1        systemReg = 32900
	REG_OSDLR_EL1        systemReg = 32924
	REG_DBGPRCR_EL1      systemReg = 32932
	REG_DBGCLAIMSET_EL1  systemReg = 33734
	REG_DBGCLAIMCLR_EL1  systemReg = 33742
	REG_TRCTRACEIDR      systemReg = 34817
	REG_TRCVICTLR        systemReg = 34818
	REG_TRCSEQEVR0       systemReg = 34820
	REG_TRCCNTRLDVR0     systemReg = 34821
	REG_TRCIMSPEC0       systemReg = 34823
	REG_TRCPRGCTLR       systemReg = 34824
	REG_TRCQCTLR         systemReg = 34825
	REG_TRCVIIECTLR      systemReg = 34826
	REG_TRCSEQEVR1       systemReg = 34828
	REG_TRCCNTRLDVR1     systemReg = 34829
	REG_TRCIMSPEC1       systemReg = 34831
	REG_TRCPROCSELR      systemReg = 34832
	REG_TRCVISSCTLR      systemReg = 34834
	REG_TRCSEQEVR2       systemReg = 34836
	REG_TRCCNTRLDVR2     systemReg = 34837
	REG_TRCIMSPEC2       systemReg = 34839
	REG_TRCVIPCSSCTLR    systemReg = 34842
	REG_TRCCNTRLDVR3     systemReg = 34845
	REG_TRCIMSPEC3       systemReg = 34847
	REG_TRCCONFIGR       systemReg = 34848
	REG_TRCCNTCTLR0      systemReg = 34853
	REG_TRCIMSPEC4       systemReg = 34855
	REG_TRCCNTCTLR1      systemReg = 34861
	REG_TRCIMSPEC5       systemReg = 34863
	REG_TRCAUXCTLR       systemReg = 34864
	REG_TRCSEQRSTEVR     systemReg = 34868
	REG_TRCCNTCTLR2      systemReg = 34869
	REG_TRCIMSPEC6       systemReg = 34871
	REG_TRCSEQSTR        systemReg = 34876
	REG_TRCCNTCTLR3      systemReg = 34877
	REG_TRCIMSPEC7       systemReg = 34879
	REG_TRCEVENTCTL0R    systemReg = 34880
	REG_TRCVDCTLR        systemReg = 34882
	REG_TRCEXTINSELR     systemReg = 34884
	REG_TRCCNTVR0        systemReg = 34885
	REG_TRCEVENTCTL1R    systemReg = 34888
	REG_TRCVDSACCTLR     systemReg = 34890
	REG_TRCEXTINSELR1    systemReg = 34892
	REG_TRCCNTVR1        systemReg = 34893
	REG_TRCRSR           systemReg = 34896
	REG_TRCVDARCCTLR     systemReg = 34898
	REG_TRCEXTINSELR2    systemReg = 34900
	REG_TRCCNTVR2        systemReg = 34901
	REG_TRCSTALLCTLR     systemReg = 34904
	REG_TRCEXTINSELR3    systemReg = 34908
	REG_TRCCNTVR3        systemReg = 34909
	REG_TRCTSCTLR        systemReg = 34912
	REG_TRCSYNCPR        systemReg = 34920
	REG_TRCCCCTLR        systemReg = 34928
	REG_TRCBBCTLR        systemReg = 34936
	REG_TRCRSCTLR16      systemReg = 34945
	REG_TRCSSCCR0        systemReg = 34946
	REG_TRCSSPCICR0      systemReg = 34947
	REG_TRCOSLAR         systemReg = 34948
	REG_TRCRSCTLR17      systemReg = 34953
	REG_TRCSSCCR1        systemReg = 34954
	REG_TRCSSPCICR1      systemReg = 34955
	REG_TRCRSCTLR2       systemReg = 34960
	REG_TRCRSCTLR18      systemReg = 34961
	REG_TRCSSCCR2        systemReg = 34962
	REG_TRCSSPCICR2      systemReg = 34963
	REG_TRCRSCTLR3       systemReg = 34968
	REG_TRCRSCTLR19      systemReg = 34969
	REG_TRCSSCCR3        systemReg = 34970
	REG_TRCSSPCICR3      systemReg = 34971
	REG_TRCRSCTLR4       systemReg = 34976
	REG_TRCRSCTLR20      systemReg = 34977
	REG_TRCSSCCR4        systemReg = 34978
	REG_TRCSSPCICR4      systemReg = 34979
	REG_TRCPDCR          systemReg = 34980
	REG_TRCRSCTLR5       systemReg = 34984
	REG_TRCRSCTLR21      systemReg = 34985
	REG_TRCSSCCR5        systemReg = 34986
	REG_TRCSSPCICR5      systemReg = 34987
	REG_TRCRSCTLR6       systemReg = 34992
	REG_TRCRSCTLR22      systemReg = 34993
	REG_TRCSSCCR6        systemReg = 34994
	REG_TRCSSPCICR6      systemReg = 34995
	REG_TRCRSCTLR7       systemReg = 35000
	REG_TRCRSCTLR23      systemReg = 35001
	REG_TRCSSCCR7        systemReg = 35002
	REG_TRCSSPCICR7      systemReg = 35003
	REG_TRCRSCTLR8       systemReg = 35008
	REG_TRCRSCTLR24      systemReg = 35009
	REG_TRCSSCSR0        systemReg = 35010
	REG_TRCRSCTLR9       systemReg = 35016
	REG_TRCRSCTLR25      systemReg = 35017
	REG_TRCSSCSR1        systemReg = 35018
	REG_TRCRSCTLR10      systemReg = 35024
	REG_TRCRSCTLR26      systemReg = 35025
	REG_TRCSSCSR2        systemReg = 35026
	REG_TRCRSCTLR11      systemReg = 35032
	REG_TRCRSCTLR27      systemReg = 35033
	REG_TRCSSCSR3        systemReg = 35034
	REG_TRCRSCTLR12      systemReg = 35040
	REG_TRCRSCTLR28      systemReg = 35041
	REG_TRCSSCSR4        systemReg = 35042
	REG_TRCRSCTLR13      systemReg = 35048
	REG_TRCRSCTLR29      systemReg = 35049
	REG_TRCSSCSR5        systemReg = 35050
	REG_TRCRSCTLR14      systemReg = 35056
	REG_TRCRSCTLR30      systemReg = 35057
	REG_TRCSSCSR6        systemReg = 35058
	REG_TRCRSCTLR15      systemReg = 35064
	REG_TRCRSCTLR31      systemReg = 35065
	REG_TRCSSCSR7        systemReg = 35066
	REG_TRCACVR0         systemReg = 35072
	REG_TRCACVR8         systemReg = 35073
	REG_TRCACATR0        systemReg = 35074
	REG_TRCACATR8        systemReg = 35075
	REG_TRCDVCVR0        systemReg = 35076
	REG_TRCDVCVR4        systemReg = 35077
	REG_TRCDVCMR0        systemReg = 35078
	REG_TRCDVCMR4        systemReg = 35079
	REG_TRCACVR1         systemReg = 35088
	REG_TRCACVR9         systemReg = 35089
	REG_TRCACATR1        systemReg = 35090
	REG_TRCACATR9        systemReg = 35091
	REG_TRCACVR2         systemReg = 35104
	REG_TRCACVR10        systemReg = 35105
	REG_TRCACATR2        systemReg = 35106
	REG_TRCACATR10       systemReg = 35107
	REG_TRCDVCVR1        systemReg = 35108
	REG_TRCDVCVR5        systemReg = 35109
	REG_TRCDVCMR1        systemReg = 35110
	REG_TRCDVCMR5        systemReg = 35111
	REG_TRCACVR3         systemReg = 35120
	REG_TRCACVR11        systemReg = 35121
	REG_TRCACATR3        systemReg = 35122
	REG_TRCACATR11       systemReg = 35123
	REG_TRCACVR4         systemReg = 35136
	REG_TRCACVR12        systemReg = 35137
	REG_TRCACATR4        systemReg = 35138
	REG_TRCACATR12       systemReg = 35139
	REG_TRCDVCVR2        systemReg = 35140
	REG_TRCDVCVR6        systemReg = 35141
	REG_TRCDVCMR2        systemReg = 35142
	REG_TRCDVCMR6        systemReg = 35143
	REG_TRCACVR5         systemReg = 35152
	REG_TRCACVR13        systemReg = 35153
	REG_TRCACATR5        systemReg = 35154
	REG_TRCACATR13       systemReg = 35155
	REG_TRCACVR6         systemReg = 35168
	REG_TRCACVR14        systemReg = 35169
	REG_TRCACATR6        systemReg = 35170
	REG_TRCACATR14       systemReg = 35171
	REG_TRCDVCVR3        systemReg = 35172
	REG_TRCDVCVR7        systemReg = 35173
	REG_TRCDVCMR3        systemReg = 35174
	REG_TRCDVCMR7        systemReg = 35175
	REG_TRCACVR7         systemReg = 35184
	REG_TRCACVR15        systemReg = 35185
	REG_TRCACATR7        systemReg = 35186
	REG_TRCACATR15       systemReg = 35187
	REG_TRCCIDCVR0       systemReg = 35200
	REG_TRCVMIDCVR0      systemReg = 35201
	REG_TRCCIDCCTLR0     systemReg = 35202
	REG_TRCCIDCCTLR1     systemReg = 35210
	REG_TRCCIDCVR1       systemReg = 35216
	REG_TRCVMIDCVR1      systemReg = 35217
	REG_TRCVMIDCCTLR0    systemReg = 35218
	REG_TRCVMIDCCTLR1    systemReg = 35226
	REG_TRCCIDCVR2       systemReg = 35232
	REG_TRCVMIDCVR2      systemReg = 35233
	REG_TRCCIDCVR3       systemReg = 35248
	REG_TRCVMIDCVR3      systemReg = 35249
	REG_TRCCIDCVR4       systemReg = 35264
	REG_TRCVMIDCVR4      systemReg = 35265
	REG_TRCCIDCVR5       systemReg = 35280
	REG_TRCVMIDCVR5      systemReg = 35281
	REG_TRCCIDCVR6       systemReg = 35296
	REG_TRCVMIDCVR6      systemReg = 35297
	REG_TRCCIDCVR7       systemReg = 35312
	REG_TRCVMIDCVR7      systemReg = 35313
	REG_TRCITCTRL        systemReg = 35716
	REG_TRCCLAIMSET      systemReg = 35782
	REG_TRCCLAIMCLR      systemReg = 35790
	REG_TRCLAR           systemReg = 35814
	REG_TEECR32_EL1      systemReg = 36864
	REG_TEEHBR32_EL1     systemReg = 36992
	REG_DBGDTR_EL0       systemReg = 38944
	REG_DBGDTRTX_EL0     systemReg = 38952
	REG_DBGVCR32_EL2     systemReg = 41016
	REG_SCTLR_EL1        systemReg = 49280
	REG_ACTLR_EL1        systemReg = 49281
	REG_CPACR_EL1        systemReg = 49282
	REG_RGSR_EL1         systemReg = 49285
	REG_GCR_EL1          systemReg = 49286
	REG_TRFCR_EL1        systemReg = 49297
	REG_TTBR0_EL1        systemReg = 49408
	REG_TTBR1_EL1        systemReg = 49409
	REG_TCR_EL1          systemReg = 49410
	REG_APIAKEYLO_EL1    systemReg = 49416
	REG_APIAKEYHI_EL1    systemReg = 49417
	REG_APIBKEYLO_EL1    systemReg = 49418
	REG_APIBKEYHI_EL1    systemReg = 49419
	REG_APDAKEYLO_EL1    systemReg = 49424
	REG_APDAKEYHI_EL1    systemReg = 49425
	REG_APDBKEYLO_EL1    systemReg = 49426
	REG_APDBKEYHI_EL1    systemReg = 49427
	REG_APGAKEYLO_EL1    systemReg = 49432
	REG_APGAKEYHI_EL1    systemReg = 49433
	REG_SPSR_EL1         systemReg = 49664
	REG_ELR_EL1          systemReg = 49665
	REG_SP_EL0           systemReg = 49672
	REG_SPSEL            systemReg = 49680
	REG_CURRENTEL        systemReg = 49682
	REG_PAN              systemReg = 49683
	REG_UAO              systemReg = 49684
	REG_ICC_PMR_EL1      systemReg = 49712
	REG_AFSR0_EL1        systemReg = 49800
	REG_AFSR1_EL1        systemReg = 49801
	REG_ESR_EL1          systemReg = 49808
	REG_ERRSELR_EL1      systemReg = 49817
	REG_ERXCTLR_EL1      systemReg = 49825
	REG_ERXSTATUS_EL1    systemReg = 49826
	REG_ERXADDR_EL1      systemReg = 49827
	REG_ERXPFGCTL_EL1    systemReg = 49829
	REG_ERXPFGCDN_EL1    systemReg = 49830
	REG_ERXMISC0_EL1     systemReg = 49832
	REG_ERXMISC1_EL1     systemReg = 49833
	REG_ERXMISC2_EL1     systemReg = 49834
	REG_ERXMISC3_EL1     systemReg = 49835
	REG_ERXTS_EL1        systemReg = 49839
	REG_TFSR_EL1         systemReg = 49840
	REG_TFSRE0_EL1       systemReg = 49841
	REG_FAR_EL1          systemReg = 49920
	REG_PAR_EL1          systemReg = 50080
	REG_PMSCR_EL1        systemReg = 50376
	REG_PMSICR_EL1       systemReg = 50378
	REG_PMSIRR_EL1       systemReg = 50379
	REG_PMSFCR_EL1       systemReg = 50380
	REG_PMSEVFR_EL1      systemReg = 50381
	REG_PMSLATFR_EL1     systemReg = 50382
	REG_PMSIDR_EL1       systemReg = 50383
	REG_PMBLIMITR_EL1    systemReg = 50384
	REG_PMBPTR_EL1       systemReg = 50385
	REG_PMBSR_EL1        systemReg = 50387
	REG_PMBIDR_EL1       systemReg = 50391
	REG_TRBLIMITR_EL1    systemReg = 50392
	REG_TRBPTR_EL1       systemReg = 50393
	REG_TRBBASER_EL1     systemReg = 50394
	REG_TRBSR_EL1        systemReg = 50395
	REG_TRBMAR_EL1       systemReg = 50396
	REG_TRBTRG_EL1       systemReg = 50398
	REG_PMINTENSET_EL1   systemReg = 50417
	REG_PMINTENCLR_EL1   systemReg = 50418
	REG_PMMIR_EL1        systemReg = 50422
	REG_MAIR_EL1         systemReg = 50448
	REG_AMAIR_EL1        systemReg = 50456
	REG_LORSA_EL1        systemReg = 50464
	REG_LOREA_EL1        systemReg = 50465
	REG_LORN_EL1         systemReg = 50466
	REG_LORC_EL1         systemReg = 50467
	REG_MPAM1_EL1        systemReg = 50472
	REG_MPAM0_EL1        systemReg = 50473
	REG_VBAR_EL1         systemReg = 50688
	REG_RMR_EL1          systemReg = 50690
	REG_DISR_EL1         systemReg = 50697
	REG_ICC_EOIR0_EL1    systemReg = 50753
	REG_ICC_BPR0_EL1     systemReg = 50755
	REG_ICC_AP0R0_EL1    systemReg = 50756
	REG_ICC_AP0R1_EL1    systemReg = 50757
	REG_ICC_AP0R2_EL1    systemReg = 50758
	REG_ICC_AP0R3_EL1    systemReg = 50759
	REG_ICC_AP1R0_EL1    systemReg = 50760
	REG_ICC_AP1R1_EL1    systemReg = 50761
	REG_ICC_AP1R2_EL1    systemReg = 50762
	REG_ICC_AP1R3_EL1    systemReg = 50763
	REG_ICC_DIR_EL1      systemReg = 50777
	REG_ICC_SGI1R_EL1    systemReg = 50781
	REG_ICC_ASGI1R_EL1   systemReg = 50782
	REG_ICC_SGI0R_EL1    systemReg = 50783
	REG_ICC_EOIR1_EL1    systemReg = 50785
	REG_ICC_BPR1_EL1     systemReg = 50787
	REG_ICC_CTLR_EL1     systemReg = 50788
	REG_ICC_SRE_EL1      systemReg = 50789
	REG_ICC_IGRPEN0_EL1  systemReg = 50790
	REG_ICC_IGRPEN1_EL1  systemReg = 50791
	REG_ICC_SEIEN_EL1    systemReg = 50792
	REG_CONTEXTIDR_EL1   systemReg = 50817
	REG_TPIDR_EL1        systemReg = 50820
	REG_SCXTNUM_EL1      systemReg = 50823
	REG_CNTKCTL_EL1      systemReg = 50952
	REG_CSSELR_EL1       systemReg = 53248
	REG_NZCV             systemReg = 55824
	REG_DAIFSET          systemReg = 55825
	REG_DIT              systemReg = 55829
	REG_SSBS             systemReg = 55830
	REG_TCO              systemReg = 55831
	REG_FPCR             systemReg = 55840
	REG_FPSR             systemReg = 55841
	REG_DSPSR_EL0        systemReg = 55848
	REG_DLR_EL0          systemReg = 55849
	REG_PMCR_EL0         systemReg = 56544
	REG_PMCNTENSET_EL0   systemReg = 56545
	REG_PMCNTENCLR_EL0   systemReg = 56546
	REG_PMOVSCLR_EL0     systemReg = 56547
	REG_PMSWINC_EL0      systemReg = 56548
	REG_PMSELR_EL0       systemReg = 56549
	REG_PMCCNTR_EL0      systemReg = 56552
	REG_PMXEVTYPER_EL0   systemReg = 56553
	REG_PMXEVCNTR_EL0    systemReg = 56554
	REG_DAIFCLR          systemReg = 56557
	REG_PMUSERENR_EL0    systemReg = 56560
	REG_PMOVSSET_EL0     systemReg = 56563
	REG_TPIDR_EL0        systemReg = 56962
	REG_TPIDRRO_EL0      systemReg = 56963
	REG_SCXTNUM_EL0      systemReg = 56967
	REG_AMCR_EL0         systemReg = 56976
	REG_AMUSERENR_EL0    systemReg = 56979
	REG_AMCNTENCLR0_EL0  systemReg = 56980
	REG_AMCNTENSET0_EL0  systemReg = 56981
	REG_AMCNTENCLR1_EL0  systemReg = 56984
	REG_AMCNTENSET1_EL0  systemReg = 56985
	REG_AMEVCNTR00_EL0   systemReg = 56992
	REG_AMEVCNTR01_EL0   systemReg = 56993
	REG_AMEVCNTR02_EL0   systemReg = 56994
	REG_AMEVCNTR03_EL0   systemReg = 56995
	REG_AMEVCNTR10_EL0   systemReg = 57056
	REG_AMEVCNTR11_EL0   systemReg = 57057
	REG_AMEVCNTR12_EL0   systemReg = 57058
	REG_AMEVCNTR13_EL0   systemReg = 57059
	REG_AMEVCNTR14_EL0   systemReg = 57060
	REG_AMEVCNTR15_EL0   systemReg = 57061
	REG_AMEVCNTR16_EL0   systemReg = 57062
	REG_AMEVCNTR17_EL0   systemReg = 57063
	REG_AMEVCNTR18_EL0   systemReg = 57064
	REG_AMEVCNTR19_EL0   systemReg = 57065
	REG_AMEVCNTR110_EL0  systemReg = 57066
	REG_AMEVCNTR111_EL0  systemReg = 57067
	REG_AMEVCNTR112_EL0  systemReg = 57068
	REG_AMEVCNTR113_EL0  systemReg = 57069
	REG_AMEVCNTR114_EL0  systemReg = 57070
	REG_AMEVCNTR115_EL0  systemReg = 57071
	REG_AMEVTYPER10_EL0  systemReg = 57072
	REG_AMEVTYPER11_EL0  systemReg = 57073
	REG_AMEVTYPER12_EL0  systemReg = 57074
	REG_AMEVTYPER13_EL0  systemReg = 57075
	REG_AMEVTYPER14_EL0  systemReg = 57076
	REG_AMEVTYPER15_EL0  systemReg = 57077
	REG_AMEVTYPER16_EL0  systemReg = 57078
	REG_AMEVTYPER17_EL0  systemReg = 57079
	REG_AMEVTYPER18_EL0  systemReg = 57080
	REG_AMEVTYPER19_EL0  systemReg = 57081
	REG_AMEVTYPER110_EL0 systemReg = 57082
	REG_AMEVTYPER111_EL0 systemReg = 57083
	REG_AMEVTYPER112_EL0 systemReg = 57084
	REG_AMEVTYPER113_EL0 systemReg = 57085
	REG_AMEVTYPER114_EL0 systemReg = 57086
	REG_AMEVTYPER115_EL0 systemReg = 57087
	REG_CNTFRQ_EL0       systemReg = 57088
	REG_CNTP_TVAL_EL0    systemReg = 57104
	REG_CNTP_CTL_EL0     systemReg = 57105
	REG_CNTP_CVAL_EL0    systemReg = 57106
	REG_CNTV_TVAL_EL0    systemReg = 57112
	REG_CNTV_CTL_EL0     systemReg = 57113
	REG_CNTV_CVAL_EL0    systemReg = 57114
	REG_PMEVCNTR0_EL0    systemReg = 57152
	REG_PMEVCNTR1_EL0    systemReg = 57153
	REG_PMEVCNTR2_EL0    systemReg = 57154
	REG_PMEVCNTR3_EL0    systemReg = 57155
	REG_PMEVCNTR4_EL0    systemReg = 57156
	REG_PMEVCNTR5_EL0    systemReg = 57157
	REG_PMEVCNTR6_EL0    systemReg = 57158
	REG_PMEVCNTR7_EL0    systemReg = 57159
	REG_PMEVCNTR8_EL0    systemReg = 57160
	REG_PMEVCNTR9_EL0    systemReg = 57161
	REG_PMEVCNTR10_EL0   systemReg = 57162
	REG_PMEVCNTR11_EL0   systemReg = 57163
	REG_PMEVCNTR12_EL0   systemReg = 57164
	REG_PMEVCNTR13_EL0   systemReg = 57165
	REG_PMEVCNTR14_EL0   systemReg = 57166
	REG_PMEVCNTR15_EL0   systemReg = 57167
	REG_PMEVCNTR16_EL0   systemReg = 57168
	REG_PMEVCNTR17_EL0   systemReg = 57169
	REG_PMEVCNTR18_EL0   systemReg = 57170
	REG_PMEVCNTR19_EL0   systemReg = 57171
	REG_PMEVCNTR20_EL0   systemReg = 57172
	REG_PMEVCNTR21_EL0   systemReg = 57173
	REG_PMEVCNTR22_EL0   systemReg = 57174
	REG_PMEVCNTR23_EL0   systemReg = 57175
	REG_PMEVCNTR24_EL0   systemReg = 57176
	REG_PMEVCNTR25_EL0   systemReg = 57177
	REG_PMEVCNTR26_EL0   systemReg = 57178
	REG_PMEVCNTR27_EL0   systemReg = 57179
	REG_PMEVCNTR28_EL0   systemReg = 57180
	REG_PMEVCNTR29_EL0   systemReg = 57181
	REG_PMEVCNTR30_EL0   systemReg = 57182
	REG_PMEVTYPER0_EL0   systemReg = 57184
	REG_PMEVTYPER1_EL0   systemReg = 57185
	REG_PMEVTYPER2_EL0   systemReg = 57186
	REG_PMEVTYPER3_EL0   systemReg = 57187
	REG_PMEVTYPER4_EL0   systemReg = 57188
	REG_PMEVTYPER5_EL0   systemReg = 57189
	REG_PMEVTYPER6_EL0   systemReg = 57190
	REG_PMEVTYPER7_EL0   systemReg = 57191
	REG_PMEVTYPER8_EL0   systemReg = 57192
	REG_PMEVTYPER9_EL0   systemReg = 57193
	REG_PMEVTYPER10_EL0  systemReg = 57194
	REG_PMEVTYPER11_EL0  systemReg = 57195
	REG_PMEVTYPER12_EL0  systemReg = 57196
	REG_PMEVTYPER13_EL0  systemReg = 57197
	REG_PMEVTYPER14_EL0  systemReg = 57198
	REG_PMEVTYPER15_EL0  systemReg = 57199
	REG_PMEVTYPER16_EL0  systemReg = 57200
	REG_PMEVTYPER17_EL0  systemReg = 57201
	REG_PMEVTYPER18_EL0  systemReg = 57202
	REG_PMEVTYPER19_EL0  systemReg = 57203
	REG_PMEVTYPER20_EL0  systemReg = 57204
	REG_PMEVTYPER21_EL0  systemReg = 57205
	REG_PMEVTYPER22_EL0  systemReg = 57206
	REG_PMEVTYPER23_EL0  systemReg = 57207
	REG_PMEVTYPER24_EL0  systemReg = 57208
	REG_PMEVTYPER25_EL0  systemReg = 57209
	REG_PMEVTYPER26_EL0  systemReg = 57210
	REG_PMEVTYPER27_EL0  systemReg = 57211
	REG_PMEVTYPER28_EL0  systemReg = 57212
	REG_PMEVTYPER29_EL0  systemReg = 57213
	REG_PMEVTYPER30_EL0  systemReg = 57214
	REG_PMCCFILTR_EL0    systemReg = 57215
	REG_VPIDR_EL2        systemReg = 57344
	REG_VMPIDR_EL2       systemReg = 57349
	REG_SCTLR_EL2        systemReg = 57472
	REG_ACTLR_EL2        systemReg = 57473
	REG_HCR_EL2          systemReg = 57480
	REG_MDCR_EL2         systemReg = 57481
	REG_CPTR_EL2         systemReg = 57482
	REG_HSTR_EL2         systemReg = 57483
	REG_HACR_EL2         systemReg = 57487
	REG_TRFCR_EL2        systemReg = 57489
	REG_SDER32_EL2       systemReg = 57497
	REG_TTBR0_EL2        systemReg = 57600
	REG_TTBR1_EL2        systemReg = 57601
	REG_TCR_EL2          systemReg = 57602
	REG_VTTBR_EL2        systemReg = 57608
	REG_VTCR_EL2         systemReg = 57610
	REG_VNCR_EL2         systemReg = 57616
	REG_VSTTBR_EL2       systemReg = 57648
	REG_VSTCR_EL2        systemReg = 57650
	REG_DACR32_EL2       systemReg = 57728
	REG_SPSR_EL2         systemReg = 57856
	REG_ELR_EL2          systemReg = 57857
	REG_SP_EL1           systemReg = 57864
	REG_SPSR_IRQ         systemReg = 57880
	REG_SPSR_ABT         systemReg = 57881
	REG_SPSR_UND         systemReg = 57882
	REG_SPSR_FIQ         systemReg = 57883
	REG_IFSR32_EL2       systemReg = 57985
	REG_AFSR0_EL2        systemReg = 57992
	REG_AFSR1_EL2        systemReg = 57993
	REG_ESR_EL2          systemReg = 58000
	REG_VSESR_EL2        systemReg = 58003
	REG_FPEXC32_EL2      systemReg = 58008
	REG_TFSR_EL2         systemReg = 58032
	REG_FAR_EL2          systemReg = 58112
	REG_HPFAR_EL2        systemReg = 58116
	REG_PMSCR_EL2        systemReg = 58568
	REG_MAIR_EL2         systemReg = 58640
	REG_AMAIR_EL2        systemReg = 58648
	REG_MPAMHCR_EL2      systemReg = 58656
	REG_MPAMVPMV_EL2     systemReg = 58657
	REG_MPAM2_EL2        systemReg = 58664
	REG_MPAMVPM0_EL2     systemReg = 58672
	REG_MPAMVPM1_EL2     systemReg = 58673
	REG_MPAMVPM2_EL2     systemReg = 58674
	REG_MPAMVPM3_EL2     systemReg = 58675
	REG_MPAMVPM4_EL2     systemReg = 58676
	REG_MPAMVPM5_EL2     systemReg = 58677
	REG_MPAMVPM6_EL2     systemReg = 58678
	REG_MPAMVPM7_EL2     systemReg = 58679
	REG_VBAR_EL2         systemReg = 58880
	REG_RMR_EL2          systemReg = 58882
	REG_VDISR_EL2        systemReg = 58889
	REG_ICH_AP0R0_EL2    systemReg = 58944
	REG_ICH_AP0R1_EL2    systemReg = 58945
	REG_ICH_AP0R2_EL2    systemReg = 58946
	REG_ICH_AP0R3_EL2    systemReg = 58947
	REG_ICH_AP1R0_EL2    systemReg = 58952
	REG_ICH_AP1R1_EL2    systemReg = 58953
	REG_ICH_AP1R2_EL2    systemReg = 58954
	REG_ICH_AP1R3_EL2    systemReg = 58955
	REG_ICH_VSEIR_EL2    systemReg = 58956
	REG_ICC_SRE_EL2      systemReg = 58957
	REG_ICH_HCR_EL2      systemReg = 58968
	REG_ICH_MISR_EL2     systemReg = 58970
	REG_ICH_VMCR_EL2     systemReg = 58975
	REG_ICH_LR0_EL2      systemReg = 58976
	REG_ICH_LR1_EL2      systemReg = 58977
	REG_ICH_LR2_EL2      systemReg = 58978
	REG_ICH_LR3_EL2      systemReg = 58979
	REG_ICH_LR4_EL2      systemReg = 58980
	REG_ICH_LR5_EL2      systemReg = 58981
	REG_ICH_LR6_EL2      systemReg = 58982
	REG_ICH_LR7_EL2      systemReg = 58983
	REG_ICH_LR8_EL2      systemReg = 58984
	REG_ICH_LR9_EL2      systemReg = 58985
	REG_ICH_LR10_EL2     systemReg = 58986
	REG_ICH_LR11_EL2     systemReg = 58987
	REG_ICH_LR12_EL2     systemReg = 58988
	REG_ICH_LR13_EL2     systemReg = 58989
	REG_ICH_LR14_EL2     systemReg = 58990
	REG_ICH_LR15_EL2     systemReg = 58991
	REG_CONTEXTIDR_EL2   systemReg = 59009
	REG_TPIDR_EL2        systemReg = 59010
	REG_SCXTNUM_EL2      systemReg = 59015
	REG_CNTVOFF_EL2      systemReg = 59139
	REG_CNTHCTL_EL2      systemReg = 59144
	REG_CNTHP_TVAL_EL2   systemReg = 59152
	REG_CNTHP_CTL_EL2    systemReg = 59153
	REG_CNTHP_CVAL_EL2   systemReg = 59154
	REG_CNTHV_TVAL_EL2   systemReg = 59160
	REG_CNTHV_CTL_EL2    systemReg = 59161
	REG_CNTHV_CVAL_EL2   systemReg = 59162
	REG_CNTHVS_TVAL_EL2  systemReg = 59168
	REG_CNTHVS_CTL_EL2   systemReg = 59169
	REG_CNTHVS_CVAL_EL2  systemReg = 59170
	REG_CNTHPS_TVAL_EL2  systemReg = 59176
	REG_CNTHPS_CTL_EL2   systemReg = 59177
	REG_CNTHPS_CVAL_EL2  systemReg = 59178
	REG_SCTLR_EL12       systemReg = 59520
	REG_CPACR_EL12       systemReg = 59522
	REG_TRFCR_EL12       systemReg = 59537
	REG_TTBR0_EL12       systemReg = 59648
	REG_TTBR1_EL12       systemReg = 59649
	REG_TCR_EL12         systemReg = 59650
	REG_SPSR_EL12        systemReg = 59904
	REG_ELR_EL12         systemReg = 59905
	REG_AFSR0_EL12       systemReg = 60040
	REG_AFSR1_EL12       systemReg = 60041
	REG_ESR_EL12         systemReg = 60048
	REG_TFSR_EL12        systemReg = 60080
	REG_FAR_EL12         systemReg = 60160
	REG_PMSCR_EL12       systemReg = 60616
	REG_MAIR_EL12        systemReg = 60688
	REG_AMAIR_EL12       systemReg = 60696
	REG_MPAM1_EL12       systemReg = 60712
	REG_VBAR_EL12        systemReg = 60928
	REG_CONTEXTIDR_EL12  systemReg = 61057
	REG_SCXTNUM_EL12     systemReg = 61063
	REG_CNTKCTL_EL12     systemReg = 61192
	REG_CNTP_TVAL_EL02   systemReg = 61200
	REG_CNTP_CTL_EL02    systemReg = 61201
	REG_CNTP_CVAL_EL02   systemReg = 61202
	REG_CNTV_TVAL_EL02   systemReg = 61208
	REG_CNTV_CTL_EL02    systemReg = 61209
	REG_CNTV_CVAL_EL02   systemReg = 61210
	REG_SCTLR_EL3        systemReg = 61568
	REG_ACTLR_EL3        systemReg = 61569
	REG_SCR_EL3          systemReg = 61576
	REG_SDER32_EL3       systemReg = 61577
	REG_CPTR_EL3         systemReg = 61578
	REG_MDCR_EL3         systemReg = 61593
	REG_TTBR0_EL3        systemReg = 61696
	REG_TCR_EL3          systemReg = 61698
	REG_SPSR_EL3         systemReg = 61952
	REG_ELR_EL3          systemReg = 61953
	REG_SP_EL2           systemReg = 61960
	REG_AFSR0_EL3        systemReg = 62088
	REG_AFSR1_EL3        systemReg = 62089
	REG_ESR_EL3          systemReg = 62096
	REG_TFSR_EL3         systemReg = 62128
	REG_FAR_EL3          systemReg = 62208
	REG_MAIR_EL3         systemReg = 62736
	REG_AMAIR_EL3        systemReg = 62744
	REG_MPAM3_EL3        systemReg = 62760
	REG_VBAR_EL3         systemReg = 62976
	REG_RMR_EL3          systemReg = 62978
	REG_ICC_CTLR_EL3     systemReg = 63076
	REG_ICC_SRE_EL3      systemReg = 63077
	REG_ICC_IGRPEN1_EL3  systemReg = 63079
	REG_TPIDR_EL3        systemReg = 63106
	REG_SCXTNUM_EL3      systemReg = 63111
	REG_CNTPS_TVAL_EL1   systemReg = 65296
	REG_CNTPS_CTL_EL1    systemReg = 65297
	REG_CNTPS_CVAL_EL1   systemReg = 65298

	/* From https://github.com/AsahiLinux/m1n1 */

	REG_MDRAR_EL1         systemReg = 32896
	REG_OSLSR_EL1         systemReg = 32908
	REG_DBGAUTHSTATUS_EL1 systemReg = 33782
	REG_MDCCSR_EL0        systemReg = 38920
	// REG_DBGDTRRX_EL0               systemReg = 38952 TODO: how to wire these up
	REG_MIDR_EL1         systemReg = 49152
	REG_MPIDR_EL1        systemReg = 49157
	REG_REVIDR_EL1       systemReg = 49158
	REG_ID_PFR0_EL1      systemReg = 49160
	REG_ID_PFR1_EL1      systemReg = 49161
	REG_ID_DFR0_EL1      systemReg = 49162
	REG_ID_AFR0_EL1      systemReg = 49163
	REG_ID_MMFR0_EL1     systemReg = 49164
	REG_ID_MMFR1_EL1     systemReg = 49165
	REG_ID_MMFR2_EL1     systemReg = 49166
	REG_ID_MMFR3_EL1     systemReg = 49167
	REG_ID_ISAR0_EL1     systemReg = 49168
	REG_ID_ISAR1_EL1     systemReg = 49169
	REG_ID_ISAR2_EL1     systemReg = 49170
	REG_ID_ISAR3_EL1     systemReg = 49171
	REG_ID_ISAR4_EL1     systemReg = 49172
	REG_ID_ISAR5_EL1     systemReg = 49173
	REG_ID_MMFR4_EL1     systemReg = 49174
	REG_ID_ISAR6_EL1     systemReg = 49175
	REG_MVFR0_EL1        systemReg = 49176
	REG_MVFR1_EL1        systemReg = 49177
	REG_MVFR2_EL1        systemReg = 49178
	REG_ID_PFR2_EL1      systemReg = 49180
	REG_ID_DFR1_EL1      systemReg = 49181
	REG_ID_MMFR5_EL1     systemReg = 49182
	REG_ID_AA64PFR0_EL1  systemReg = 49184
	REG_ID_AA64PFR1_EL1  systemReg = 49185
	REG_ID_AA64ZFR0_EL1  systemReg = 49188
	REG_ID_AA64DFR0_EL1  systemReg = 49192
	REG_ID_AA64DFR1_EL1  systemReg = 49193
	REG_ID_AA64AFR0_EL1  systemReg = 49196
	REG_ID_AA64AFR1_EL1  systemReg = 49197
	REG_ID_AA64ISAR0_EL1 systemReg = 49200
	REG_ID_AA64ISAR1_EL1 systemReg = 49201
	REG_ID_AA64ISAR2_EL1 systemReg = 49202
	REG_ID_AA64MMFR0_EL1 systemReg = 49208
	REG_ID_AA64MMFR1_EL1 systemReg = 49209
	REG_ID_AA64MMFR2_EL1 systemReg = 49210
	REG_ZCR_EL1          systemReg = 49296
	REG_ERRIDR_EL1       systemReg = 49816
	REG_ERXFR_EL1        systemReg = 49824
	REG_ERXPFGF_EL1      systemReg = 49828
	REG_PMSNEVFR_EL1     systemReg = 50377
	REG_MPAMIDR_EL1      systemReg = 50468
	REG_LORID_EL1        systemReg = 50471
	REG_RVBAR_EL1        systemReg = 50689
	REG_ISR_EL1          systemReg = 50696
	REG_ICC_IAR0_EL1     systemReg = 50752
	REG_ICC_HPPIR0_EL1   systemReg = 50754
	REG_ICC_RPR_EL1      systemReg = 50779
	REG_ICC_IAR1_EL1     systemReg = 50784
	REG_ICC_HPPIR1_EL1   systemReg = 50786
	REG_ACCDATA_EL1      systemReg = 50821
	REG_HID0_EL1         systemReg = 51072
	REG_EHID0_EL1        systemReg = 51073
	REG_HID1_EL1         systemReg = 51080
	REG_EHID1_EL1        systemReg = 51081
	REG_EHID20_EL1       systemReg = 51082
	REG_HID21_EL1        systemReg = 51083
	REG_HID2_EL1         systemReg = 51088
	REG_EHID2_EL1        systemReg = 51089
	REG_HID3_EL1         systemReg = 51096
	REG_EHID3_EL1        systemReg = 51097
	REG_HID4_EL1         systemReg = 51104
	REG_EHID4_EL1        systemReg = 51105
	REG_HID5_EL1         systemReg = 51112
	REG_EHID5_EL1        systemReg = 51113
	REG_HID6_EL1         systemReg = 51120
	REG_HID7_EL1         systemReg = 51128
	REG_EHID7_EL1        systemReg = 51129
	REG_HID8_EL1         systemReg = 51136
	REG_HID9_EL1         systemReg = 51144
	REG_EHID9_EL1        systemReg = 51145
	REG_HID10_EL1        systemReg = 51152
	REG_EHID10_EL1       systemReg = 51153
	REG_HID11_EL1        systemReg = 51160
	REG_EHID11_EL1       systemReg = 51161
	REG_HID18_EL1        systemReg = 51162
	REG_HID13_EL1        systemReg = 51184
	REG_HID14_EL1        systemReg = 51192
	REG_HID16_EL1        systemReg = 51194
	REG_HID17_EL1        systemReg = 51197
	REG_CCSIDR_EL1       systemReg = 51200
	REG_CLIDR_EL1        systemReg = 51201
	REG_CCSIDR2_EL1      systemReg = 51202
	REG_GMID_EL1         systemReg = 51204
	REG_AIDR_EL1         systemReg = 51207
	REG_PMCR0_EL1        systemReg = 53120
	REG_PMCR1_EL1        systemReg = 53128
	REG_PMCR2_EL1        systemReg = 53136
	REG_PMCR3_EL1        systemReg = 53144
	REG_PMCR4_EL1        systemReg = 53152
	REG_PMESR0_EL1       systemReg = 53160
	REG_PMESR1_EL1       systemReg = 53168
	REG_PMSR_EL1         systemReg = 53224
	REG_PMC0_EL1         systemReg = 55168
	REG_PMC1_EL1         systemReg = 55176
	REG_PMC2_EL1         systemReg = 55184
	REG_PMC3_EL1         systemReg = 55192
	REG_PMC4_EL1         systemReg = 55200
	REG_PMC5_EL1         systemReg = 55208
	REG_PMC6_EL1         systemReg = 55216
	REG_PMC7_EL1         systemReg = 55224
	REG_PMC8_EL1         systemReg = 55240
	REG_PMC9_EL1         systemReg = 55248
	REG_CTR_EL0          systemReg = 55297
	REG_DCZID_EL0        systemReg = 55303
	REG_RNDR             systemReg = 55584
	REG_RNDRRS           systemReg = 55585
	REG_PMCEID0_EL0      systemReg = 56550
	REG_PMCEID1_EL0      systemReg = 56551
	REG_AMCFGR_EL0       systemReg = 56977
	REG_AMCGCR_EL0       systemReg = 56978
	REG_AMCG1IDR_EL0     systemReg = 56982
	REG_AMEVTYPER00_EL0  systemReg = 57008
	// REG_AMEVTYPER010_EL0 systemReg = 57008 TODO: how to wire these up
	// REG_AMEVTYPER011_EL0 systemReg = 57008
	// REG_AMEVTYPER012_EL0 systemReg = 57008
	// REG_AMEVTYPER013_EL0 systemReg = 57008
	// REG_AMEVTYPER014_EL0 systemReg = 57008
	// REG_AMEVTYPER015_EL0 systemReg = 57008
	// REG_AMEVTYPER01_EL0  systemReg = 57008
	// REG_AMEVTYPER02_EL0  systemReg = 57008
	// REG_AMEVTYPER03_EL0  systemReg = 57008
	// REG_AMEVTYPER04_EL0  systemReg = 57008
	// REG_AMEVTYPER05_EL0  systemReg = 57008
	// REG_AMEVTYPER06_EL0  systemReg = 57008
	// REG_AMEVTYPER07_EL0  systemReg = 57008
	// REG_AMEVTYPER08_EL0  systemReg = 57008
	// REG_AMEVTYPER09_EL0  systemReg = 57008
	REG_CNTPCT_EL0        systemReg = 57089
	REG_CNTVCT_EL0        systemReg = 57090
	REG_CNTPCTSS_EL0      systemReg = 57093
	REG_CNTVCTSS_EL0      systemReg = 57094
	REG_LSU_ERR_STS_EL1   systemReg = 57216
	REG_LSU_ERR_CTL_EL1   systemReg = 57224
	REG_E_LSU_ERR_STS_EL1 systemReg = 57232
	REG_L2C_ERR_STS_EL1   systemReg = 57280
	REG_L2C_ERR_ADR_EL1   systemReg = 57288
	REG_L2C_ERR_INF_EL1   systemReg = 57296
	REG_HFGRTR_EL2        systemReg = 57484
	REG_HFGWTR_EL2        systemReg = 57485
	REG_HFGITR_EL2        systemReg = 57486
	REG_ZCR_EL2           systemReg = 57488
	REG_HCRX_EL2          systemReg = 57490
	REG_HDFGRTR_EL2       systemReg = 57740
	REG_HDFGWTR_EL2       systemReg = 57741
	REG_HAFGRTR_EL2       systemReg = 57742
	REG_RVBAR_EL2         systemReg = 58881
	REG_ICH_VTR_EL2       systemReg = 58969
	REG_ICH_EISR_EL2      systemReg = 58971
	REG_ICH_ELRSR_EL2     systemReg = 58973
	REG_AMEVCNTVOFF00_EL2 systemReg = 59072
	// REG_AMEVCNTVOFF010_EL2         systemReg = 59072 TODO: how to wire these up
	// REG_AMEVCNTVOFF011_EL2         systemReg = 59072
	// REG_AMEVCNTVOFF012_EL2         systemReg = 59072
	// REG_AMEVCNTVOFF013_EL2         systemReg = 59072
	// REG_AMEVCNTVOFF014_EL2         systemReg = 59072
	// REG_AMEVCNTVOFF015_EL2         systemReg = 59072
	// REG_AMEVCNTVOFF01_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF02_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF03_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF04_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF05_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF06_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF07_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF08_EL2          systemReg = 59072
	// REG_AMEVCNTVOFF09_EL2          systemReg = 59072
	REG_AMEVCNTVOFF10_EL2 systemReg = 59088
	// REG_AMEVCNTVOFF110_EL2         systemReg = 59088 TODO: how to wire these up
	// REG_AMEVCNTVOFF111_EL2         systemReg = 59088
	// REG_AMEVCNTVOFF112_EL2         systemReg = 59088
	// REG_AMEVCNTVOFF113_EL2         systemReg = 59088
	// REG_AMEVCNTVOFF114_EL2         systemReg = 59088
	// REG_AMEVCNTVOFF115_EL2         systemReg = 59088
	// REG_AMEVCNTVOFF11_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF12_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF13_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF14_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF15_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF16_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF17_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF18_EL2          systemReg = 59088
	// REG_AMEVCNTVOFF19_EL2          systemReg = 59088
	REG_CNTPOFF_EL2                systemReg = 59142
	REG_FED_ERR_STS_EL1            systemReg = 59264
	REG_E_FED_ERR_STS_EL1          systemReg = 59266
	REG_APCTL_EL1                  systemReg = 59268
	REG_KERNELKEYLO_EL1            systemReg = 59272
	REG_KERNELKEYHI_EL1            systemReg = 59273
	REG_VMSA_LOCK_EL1              systemReg = 59274
	REG_AMX_CTL_EL1                systemReg = 59276
	REG_APRR_EL0                   systemReg = 59280
	REG_APRR_EL1                   systemReg = 59281
	REG_CTRR_LOCK_EL1              systemReg = 59282
	REG_CTRR_A_LWR_EL1             systemReg = 59283
	REG_CTRR_A_UPR_EL1             systemReg = 59284
	REG_CTRR_CTL_EL1               systemReg = 59285
	REG_APRR_JIT_ENABLE_EL2        systemReg = 59286
	REG_APRR_JIT_MASK_EL2          systemReg = 59287
	REG_AMX_CTL_EL12               systemReg = 59302
	REG_AMX_CTL_EL2                systemReg = 59303
	REG_SPRR_PERM_EL20_SILLY_THING systemReg = 59305
	REG_SPRR_PERM_EL02             systemReg = 59306
	REG_SPRR_UMASK0_EL2            systemReg = 59320
	REG_SPRR_UMASK1_EL2            systemReg = 59321
	REG_SPRR_UMASK2_EL2            systemReg = 59322
	REG_SPRR_UMASK3_EL2            systemReg = 59323
	REG_SPRR_UMASK0_EL12           systemReg = 59328
	REG_SPRR_UMASK1_EL12           systemReg = 59329
	REG_SPRR_UMASK2_EL12           systemReg = 59330
	REG_SPRR_UMASK3_EL12           systemReg = 59331
	REG_CNTPCT_ALIAS_EL0           systemReg = 59349
	REG_CNTVCT_ALIAS_EL0           systemReg = 59350
	REG_CTRR_A_LWR_EL2             systemReg = 59352
	REG_CTRR_A_UPR_EL2             systemReg = 59353
	REG_CTRR_CTL_EL2               systemReg = 59356
	REG_CTRR_LOCK_EL2              systemReg = 59357
	REG_ZCR_EL12                   systemReg = 59536
	REG_IPI_RR_LOCAL_EL1           systemReg = 61312
	REG_IPI_RR_GLOBAL_EL1          systemReg = 61313
	REG_DPC_ERR_STS_EL1            systemReg = 61317
	REG_IPI_SR_EL1                 systemReg = 61321
	REG_VM_TMR_LR_EL2              systemReg = 61322
	REG_VM_TMR_FIQ_ENA_EL2         systemReg = 61323
	REG_IPI_CR_EL1                 systemReg = 61337
	REG_ACC_CFG_EL1                systemReg = 61344
	REG_CYC_OVRD_EL1               systemReg = 61352
	REG_ACC_OVRD_EL1               systemReg = 61360
	REG_ACC_EBLK_OVRD_EL1          systemReg = 61361
	REG_ZCR_EL3                    systemReg = 61584
	REG_RVBAR_EL3                  systemReg = 62977
	REG_MMU_ERR_STS_EL1            systemReg = 63360
	REG_AFSR1_GL1                  systemReg = 63361
	REG_AFSR1_GL2                  systemReg = 63362
	REG_AFSR1_GL12                 systemReg = 63363
	REG_SPRR_CONFIG_EL1            systemReg = 63368
	REG_GXF_CONFIG_EL1             systemReg = 63370
	REG_SPRR_UNK1_EL1              systemReg = 63371
	REG_GXF_CONFIG_EL2             systemReg = 63372
	REG_SPRR_PERM_EL0              systemReg = 63373
	REG_SPRR_PERM_EL1              systemReg = 63374
	REG_SPRR_PERM_EL2              systemReg = 63375
	REG_E_MMU_ERR_STS_EL1          systemReg = 63376
	REG_APGAKeyLo_EL12             systemReg = 63377
	REG_APGAKeyHi_EL12             systemReg = 63378
	REG_KERNELKEYLO_EL12           systemReg = 63379
	REG_KERNELKEYHI_EL12           systemReg = 63380
	REG_AFPCR_EL0                  systemReg = 63381
	REG_AIDR2_EL1                  systemReg = 63383
	REG_SPRR_UMASK0_EL1            systemReg = 63384
	REG_SPRR_KMASK0_EL1            systemReg = 63385
	REG_SPRR_KMASK0_EL2            systemReg = 63386
	REG_SPRR_UMASK1_EL1            systemReg = 63387
	REG_SPRR_UMASK2_EL1            systemReg = 63388
	REG_SPRR_UMASK3_EL1            systemReg = 63389
	REG_SPRR_KMASK1_EL1            systemReg = 63394
	REG_SPRR_KMASK2_EL1            systemReg = 63395
	REG_SPRR_KMASK3_EL1            systemReg = 63396
	REG_SPRR_KMASK1_EL2            systemReg = 63401
	REG_SPRR_KMASK2_EL2            systemReg = 63402
	REG_SPRR_KMASK3_EL2            systemReg = 63403
	REG_SPRR_KMASK0_EL12           systemReg = 63408
	REG_SPRR_KMASK1_EL12           systemReg = 63409
	REG_SPRR_KMASK2_EL12           systemReg = 63410
	REG_SPRR_KMASK3_EL12           systemReg = 63411
	REG_APIAKeyLo_EL12             systemReg = 63416
	REG_APIAKeyHi_EL12             systemReg = 63417
	REG_APIBKeyLo_EL12             systemReg = 63418
	REG_APIBKeyHi_EL12             systemReg = 63419
	REG_APDAKeyLo_EL12             systemReg = 63420
	REG_APDAKeyHi_EL12             systemReg = 63421
	REG_APDBKeyLo_EL12             systemReg = 63422
	REG_APDBKeyHi_EL12             systemReg = 63423
	REG_GXF_STATUS_EL1             systemReg = 63424
	REG_GXF_ENTER_EL1              systemReg = 63425
	REG_GXF_ABORT_EL1              systemReg = 63426
	REG_VBAR_GL12                  systemReg = 63434
	REG_SPSR_GL12                  systemReg = 63435
	REG_ASPSR_GL12                 systemReg = 63436
	REG_ESR_GL12                   systemReg = 63437
	REG_ELR_GL12                   systemReg = 63438
	REG_SP_GL12                    systemReg = 63440
	REG_TPIDR_GL1                  systemReg = 63441
	REG_VBAR_GL1                   systemReg = 63442
	REG_SPSR_GL1                   systemReg = 63443
	REG_ASPSR_GL1                  systemReg = 63444
	REG_ESR_GL1                    systemReg = 63445
	REG_ELR_GL1                    systemReg = 63446
	REG_FAR_GL1                    systemReg = 63447
	REG_TPIDR_GL2                  systemReg = 63449
	REG_VBAR_GL2                   systemReg = 63450
	REG_SPSR_GL2                   systemReg = 63451
	REG_ASPSR_GL2                  systemReg = 63452
	REG_ESR_GL2                    systemReg = 63453
	REG_ELR_GL2                    systemReg = 63454
	REG_FAR_GL2                    systemReg = 63455
	REG_GXF_ENTER_EL2              systemReg = 63456
	REG_GXF_ABORT_EL2              systemReg = 63457
	REG_APCTL_EL2                  systemReg = 63458
	REG_APSTS_EL2_MAYBE            systemReg = 63459
	REG_APSTS_EL1                  systemReg = 63460
	REG_SPRR_CONFIG_EL2            systemReg = 63474
	REG_SPRR_UNK1_EL2              systemReg = 63475
	REG_APVMKEYLO_EL2              systemReg = 63476
	REG_APVMKEYHI_EL2              systemReg = 63477
	REG_ACTLR_EL12                 systemReg = 63478
	REG_APSTS_EL12                 systemReg = 63479
	REG_APCTL_EL12                 systemReg = 63480
	REG_GXF_CONFIG_EL12            systemReg = 63481
	REG_GXF_ENTER_EL12             systemReg = 63482
	REG_GXF_ABORT_EL12             systemReg = 63483
	REG_SPRR_CONFIG_EL12           systemReg = 63484
	REG_SPRR_UNK1_EL12             systemReg = 63485
	REG_SPRR_PERM_EL12             systemReg = 63487
	REG_UPMCR0_EL1                 systemReg = 65412
	REG_UPMC8_EL1                  systemReg = 65413
	REG_UPMESR0_EL1                systemReg = 65420
	REG_UPMC9_EL1                  systemReg = 65421
	REG_UPMC10_EL1                 systemReg = 65429
	REG_UPMECM0_EL1                systemReg = 65436
	REG_UPMC11_EL1                 systemReg = 65437
	REG_UPMECM1_EL1                systemReg = 65444
	REG_UPMC12_EL1                 systemReg = 65445
	REG_UPMPCM_EL1                 systemReg = 65452
	REG_UPMC13_EL1                 systemReg = 65453
	REG_UPMSR_EL1                  systemReg = 65460
	REG_UPMC14_EL1                 systemReg = 65461
	REG_UPMC0_EL1                  systemReg = 65468
	REG_UPMC15_EL1                 systemReg = 65469
	REG_UPMC1_EL1                  systemReg = 65476
	REG_UPMECM2_EL1                systemReg = 65477
	REG_UPMC2_EL1                  systemReg = 65484
	REG_UPMECM3_EL1                systemReg = 65485
	REG_UPMC3_EL1                  systemReg = 65492
	REG_UPMC4_EL1                  systemReg = 65500
	REG_UPMESR1_EL1                systemReg = 65501
	REG_UPMC5_EL1                  systemReg = 65508
	REG_UPMC6_EL1                  systemReg = 65516
	REG_UPMC7_EL1                  systemReg = 65524

	/* exceptional system registers */
	REG_PSTATE_SPSEL systemReg = 65299 // (op0op1crncrmop2)=(00495) doesn't map to [SYSREG_NONE+1 SYSREG_END)
	/* end marker needed for other reg defines */
	SYSREG_UNKNOWN systemReg = 65300
	SYSREG_END     systemReg = 65301
)

func (s systemReg) String() string {
	switch s {
	case 32770:
		return "osdtrrx_el1"
	case 32772:
		return "dbgbvr0_el1"
	case 32773:
		return "dbgbcr0_el1"
	case 32774:
		return "dbgwvr0_el1"
	case 32775:
		return "dbgwcr0_el1"
	case 32780:
		return "dbgbvr1_el1"
	case 32781:
		return "dbgbcr1_el1"
	case 32782:
		return "dbgwvr1_el1"
	case 32783:
		return "dbgwcr1_el1"
	case 32784:
		return "mdccint_el1"
	case 32786:
		return "mdscr_el1"
	case 32788:
		return "dbgbvr2_el1"
	case 32789:
		return "dbgbcr2_el1"
	case 32790:
		return "dbgwvr2_el1"
	case 32791:
		return "dbgwcr2_el1"
	case 32794:
		return "osdtrtx_el1"
	case 32796:
		return "dbgbvr3_el1"
	case 32797:
		return "dbgbcr3_el1"
	case 32798:
		return "dbgwvr3_el1"
	case 32799:
		return "dbgwcr3_el1"
	case 32804:
		return "dbgbvr4_el1"
	case 32805:
		return "dbgbcr4_el1"
	case 32806:
		return "dbgwvr4_el1"
	case 32807:
		return "dbgwcr4_el1"
	case 32812:
		return "dbgbvr5_el1"
	case 32813:
		return "dbgbcr5_el1"
	case 32814:
		return "dbgwvr5_el1"
	case 32815:
		return "dbgwcr5_el1"
	case 32818:
		return "oseccr_el1"
	case 32820:
		return "dbgbvr6_el1"
	case 32821:
		return "dbgbcr6_el1"
	case 32822:
		return "dbgwvr6_el1"
	case 32823:
		return "dbgwcr6_el1"
	case 32828:
		return "dbgbvr7_el1"
	case 32829:
		return "dbgbcr7_el1"
	case 32830:
		return "dbgwvr7_el1"
	case 32831:
		return "dbgwcr7_el1"
	case 32836:
		return "dbgbvr8_el1"
	case 32837:
		return "dbgbcr8_el1"
	case 32838:
		return "dbgwvr8_el1"
	case 32839:
		return "dbgwcr8_el1"
	case 32844:
		return "dbgbvr9_el1"
	case 32845:
		return "dbgbcr9_el1"
	case 32846:
		return "dbgwvr9_el1"
	case 32847:
		return "dbgwcr9_el1"
	case 32852:
		return "dbgbvr10_el1"
	case 32853:
		return "dbgbcr10_el1"
	case 32854:
		return "dbgwvr10_el1"
	case 32855:
		return "dbgwcr10_el1"
	case 32860:
		return "dbgbvr11_el1"
	case 32861:
		return "dbgbcr11_el1"
	case 32862:
		return "dbgwvr11_el1"
	case 32863:
		return "dbgwcr11_el1"
	case 32868:
		return "dbgbvr12_el1"
	case 32869:
		return "dbgbcr12_el1"
	case 32870:
		return "dbgwvr12_el1"
	case 32871:
		return "dbgwcr12_el1"
	case 32876:
		return "dbgbvr13_el1"
	case 32877:
		return "dbgbcr13_el1"
	case 32878:
		return "dbgwvr13_el1"
	case 32879:
		return "dbgwcr13_el1"
	case 32884:
		return "dbgbvr14_el1"
	case 32885:
		return "dbgbcr14_el1"
	case 32886:
		return "dbgwvr14_el1"
	case 32887:
		return "dbgwcr14_el1"
	case 32892:
		return "dbgbvr15_el1"
	case 32893:
		return "dbgbcr15_el1"
	case 32894:
		return "dbgwvr15_el1"
	case 32895:
		return "dbgwcr15_el1"
	case 32900:
		return "oslar_el1"
	case 32924:
		return "osdlr_el1"
	case 32932:
		return "dbgprcr_el1"
	case 33734:
		return "dbgclaimset_el1"
	case 33742:
		return "dbgclaimclr_el1"
	case 34817:
		return "trctraceidr"
	case 34818:
		return "trcvictlr"
	case 34820:
		return "trcseqevr0"
	case 34821:
		return "trccntrldvr0"
	case 34823:
		return "trcimspec0"
	case 34824:
		return "trcprgctlr"
	case 34825:
		return "trcqctlr"
	case 34826:
		return "trcviiectlr"
	case 34828:
		return "trcseqevr1"
	case 34829:
		return "trccntrldvr1"
	case 34831:
		return "trcimspec1"
	case 34832:
		return "trcprocselr"
	case 34834:
		return "trcvissctlr"
	case 34836:
		return "trcseqevr2"
	case 34837:
		return "trccntrldvr2"
	case 34839:
		return "trcimspec2"
	case 34842:
		return "trcvipcssctlr"
	case 34845:
		return "trccntrldvr3"
	case 34847:
		return "trcimspec3"
	case 34848:
		return "trcconfigr"
	case 34853:
		return "trccntctlr0"
	case 34855:
		return "trcimspec4"
	case 34861:
		return "trccntctlr1"
	case 34863:
		return "trcimspec5"
	case 34864:
		return "trcauxctlr"
	case 34868:
		return "trcseqrstevr"
	case 34869:
		return "trccntctlr2"
	case 34871:
		return "trcimspec6"
	case 34876:
		return "trcseqstr"
	case 34877:
		return "trccntctlr3"
	case 34879:
		return "trcimspec7"
	case 34880:
		return "trceventctl0r"
	case 34882:
		return "trcvdctlr"
	case 34884:
		return "trcextinselr"
	case 34885:
		return "trccntvr0"
	case 34888:
		return "trceventctl1r"
	case 34890:
		return "trcvdsacctlr"
	case 34892:
		return "trcextinselr1"
	case 34893:
		return "trccntvr1"
	case 34896:
		return "trcrsr"
	case 34898:
		return "trcvdarcctlr"
	case 34900:
		return "trcextinselr2"
	case 34901:
		return "trccntvr2"
	case 34904:
		return "trcstallctlr"
	case 34908:
		return "trcextinselr3"
	case 34909:
		return "trccntvr3"
	case 34912:
		return "trctsctlr"
	case 34920:
		return "trcsyncpr"
	case 34928:
		return "trcccctlr"
	case 34936:
		return "trcbbctlr"
	case 34945:
		return "trcrsctlr16"
	case 34946:
		return "trcssccr0"
	case 34947:
		return "trcsspcicr0"
	case 34948:
		return "trcoslar"
	case 34953:
		return "trcrsctlr17"
	case 34954:
		return "trcssccr1"
	case 34955:
		return "trcsspcicr1"
	case 34960:
		return "trcrsctlr2"
	case 34961:
		return "trcrsctlr18"
	case 34962:
		return "trcssccr2"
	case 34963:
		return "trcsspcicr2"
	case 34968:
		return "trcrsctlr3"
	case 34969:
		return "trcrsctlr19"
	case 34970:
		return "trcssccr3"
	case 34971:
		return "trcsspcicr3"
	case 34976:
		return "trcrsctlr4"
	case 34977:
		return "trcrsctlr20"
	case 34978:
		return "trcssccr4"
	case 34979:
		return "trcsspcicr4"
	case 34980:
		return "trcpdcr"
	case 34984:
		return "trcrsctlr5"
	case 34985:
		return "trcrsctlr21"
	case 34986:
		return "trcssccr5"
	case 34987:
		return "trcsspcicr5"
	case 34992:
		return "trcrsctlr6"
	case 34993:
		return "trcrsctlr22"
	case 34994:
		return "trcssccr6"
	case 34995:
		return "trcsspcicr6"
	case 35000:
		return "trcrsctlr7"
	case 35001:
		return "trcrsctlr23"
	case 35002:
		return "trcssccr7"
	case 35003:
		return "trcsspcicr7"
	case 35008:
		return "trcrsctlr8"
	case 35009:
		return "trcrsctlr24"
	case 35010:
		return "trcsscsr0"
	case 35016:
		return "trcrsctlr9"
	case 35017:
		return "trcrsctlr25"
	case 35018:
		return "trcsscsr1"
	case 35024:
		return "trcrsctlr10"
	case 35025:
		return "trcrsctlr26"
	case 35026:
		return "trcsscsr2"
	case 35032:
		return "trcrsctlr11"
	case 35033:
		return "trcrsctlr27"
	case 35034:
		return "trcsscsr3"
	case 35040:
		return "trcrsctlr12"
	case 35041:
		return "trcrsctlr28"
	case 35042:
		return "trcsscsr4"
	case 35048:
		return "trcrsctlr13"
	case 35049:
		return "trcrsctlr29"
	case 35050:
		return "trcsscsr5"
	case 35056:
		return "trcrsctlr14"
	case 35057:
		return "trcrsctlr30"
	case 35058:
		return "trcsscsr6"
	case 35064:
		return "trcrsctlr15"
	case 35065:
		return "trcrsctlr31"
	case 35066:
		return "trcsscsr7"
	case 35072:
		return "trcacvr0"
	case 35073:
		return "trcacvr8"
	case 35074:
		return "trcacatr0"
	case 35075:
		return "trcacatr8"
	case 35076:
		return "trcdvcvr0"
	case 35077:
		return "trcdvcvr4"
	case 35078:
		return "trcdvcmr0"
	case 35079:
		return "trcdvcmr4"
	case 35088:
		return "trcacvr1"
	case 35089:
		return "trcacvr9"
	case 35090:
		return "trcacatr1"
	case 35091:
		return "trcacatr9"
	case 35104:
		return "trcacvr2"
	case 35105:
		return "trcacvr10"
	case 35106:
		return "trcacatr2"
	case 35107:
		return "trcacatr10"
	case 35108:
		return "trcdvcvr1"
	case 35109:
		return "trcdvcvr5"
	case 35110:
		return "trcdvcmr1"
	case 35111:
		return "trcdvcmr5"
	case 35120:
		return "trcacvr3"
	case 35121:
		return "trcacvr11"
	case 35122:
		return "trcacatr3"
	case 35123:
		return "trcacatr11"
	case 35136:
		return "trcacvr4"
	case 35137:
		return "trcacvr12"
	case 35138:
		return "trcacatr4"
	case 35139:
		return "trcacatr12"
	case 35140:
		return "trcdvcvr2"
	case 35141:
		return "trcdvcvr6"
	case 35142:
		return "trcdvcmr2"
	case 35143:
		return "trcdvcmr6"
	case 35152:
		return "trcacvr5"
	case 35153:
		return "trcacvr13"
	case 35154:
		return "trcacatr5"
	case 35155:
		return "trcacatr13"
	case 35168:
		return "trcacvr6"
	case 35169:
		return "trcacvr14"
	case 35170:
		return "trcacatr6"
	case 35171:
		return "trcacatr14"
	case 35172:
		return "trcdvcvr3"
	case 35173:
		return "trcdvcvr7"
	case 35174:
		return "trcdvcmr3"
	case 35175:
		return "trcdvcmr7"
	case 35184:
		return "trcacvr7"
	case 35185:
		return "trcacvr15"
	case 35186:
		return "trcacatr7"
	case 35187:
		return "trcacatr15"
	case 35200:
		return "trccidcvr0"
	case 35201:
		return "trcvmidcvr0"
	case 35202:
		return "trccidcctlr0"
	case 35210:
		return "trccidcctlr1"
	case 35216:
		return "trccidcvr1"
	case 35217:
		return "trcvmidcvr1"
	case 35218:
		return "trcvmidcctlr0"
	case 35226:
		return "trcvmidcctlr1"
	case 35232:
		return "trccidcvr2"
	case 35233:
		return "trcvmidcvr2"
	case 35248:
		return "trccidcvr3"
	case 35249:
		return "trcvmidcvr3"
	case 35264:
		return "trccidcvr4"
	case 35265:
		return "trcvmidcvr4"
	case 35280:
		return "trccidcvr5"
	case 35281:
		return "trcvmidcvr5"
	case 35296:
		return "trccidcvr6"
	case 35297:
		return "trcvmidcvr6"
	case 35312:
		return "trccidcvr7"
	case 35313:
		return "trcvmidcvr7"
	case 35716:
		return "trcitctrl"
	case 35782:
		return "trcclaimset"
	case 35790:
		return "trcclaimclr"
	case 35814:
		return "trclar"
	case 36864:
		return "teecr32_el1"
	case 36992:
		return "teehbr32_el1"
	case 38944:
		return "dbgdtr_el0"
	case 38952:
		return "dbgdtrtx_el0"
	case 41016:
		return "dbgvcr32_el2"
	case 49280:
		return "sctlr_el1"
	case 49281:
		return "actlr_el1"
	case 49282:
		return "cpacr_el1"
	case 49285:
		return "rgsr_el1"
	case 49286:
		return "gcr_el1"
	case 49297:
		return "trfcr_el1"
	case 49408:
		return "ttbr0_el1"
	case 49409:
		return "ttbr1_el1"
	case 49410:
		return "tcr_el1"
	case 49416:
		return "apiakeylo_el1"
	case 49417:
		return "apiakeyhi_el1"
	case 49418:
		return "apibkeylo_el1"
	case 49419:
		return "apibkeyhi_el1"
	case 49424:
		return "apdakeylo_el1"
	case 49425:
		return "apdakeyhi_el1"
	case 49426:
		return "apdbkeylo_el1"
	case 49427:
		return "apdbkeyhi_el1"
	case 49432:
		return "apgakeylo_el1"
	case 49433:
		return "apgakeyhi_el1"
	case 49664:
		return "spsr_el1"
	case 49665:
		return "elr_el1"
	case 49672:
		return "sp_el0"
	case 49680:
		return "spsel"
	case 49682:
		return "currentel"
	case 49683:
		return "pan"
	case 49684:
		return "uao"
	case 49712:
		return "icc_pmr_el1"
	case 49800:
		return "afsr0_el1"
	case 49801:
		return "afsr1_el1"
	case 49808:
		return "esr_el1"
	case 49817:
		return "errselr_el1"
	case 49825:
		return "erxctlr_el1"
	case 49826:
		return "erxstatus_el1"
	case 49827:
		return "erxaddr_el1"
	case 49829:
		return "erxpfgctl_el1"
	case 49830:
		return "erxpfgcdn_el1"
	case 49832:
		return "erxmisc0_el1"
	case 49833:
		return "erxmisc1_el1"
	case 49834:
		return "erxmisc2_el1"
	case 49835:
		return "erxmisc3_el1"
	case 49839:
		return "erxts_el1"
	case 49840:
		return "tfsr_el1"
	case 49841:
		return "tfsre0_el1"
	case 49920:
		return "far_el1"
	case 50080:
		return "par_el1"
	case 50376:
		return "pmscr_el1"
	case 50378:
		return "pmsicr_el1"
	case 50379:
		return "pmsirr_el1"
	case 50380:
		return "pmsfcr_el1"
	case 50381:
		return "pmsevfr_el1"
	case 50382:
		return "pmslatfr_el1"
	case 50383:
		return "pmsidr_el1"
	case 50384:
		return "pmblimitr_el1"
	case 50385:
		return "pmbptr_el1"
	case 50387:
		return "pmbsr_el1"
	case 50391:
		return "pmbidr_el1"
	case 50392:
		return "trblimitr_el1"
	case 50393:
		return "trbptr_el1"
	case 50394:
		return "trbbaser_el1"
	case 50395:
		return "trbsr_el1"
	case 50396:
		return "trbmar_el1"
	case 50398:
		return "trbtrg_el1"
	case 50417:
		return "pmintenset_el1"
	case 50418:
		return "pmintenclr_el1"
	case 50422:
		return "pmmir_el1"
	case 50448:
		return "mair_el1"
	case 50456:
		return "amair_el1"
	case 50464:
		return "lorsa_el1"
	case 50465:
		return "lorea_el1"
	case 50466:
		return "lorn_el1"
	case 50467:
		return "lorc_el1"
	case 50472:
		return "mpam1_el1"
	case 50473:
		return "mpam0_el1"
	case 50688:
		return "vbar_el1"
	case 50690:
		return "rmr_el1"
	case 50697:
		return "disr_el1"
	case 50753:
		return "icc_eoir0_el1"
	case 50755:
		return "icc_bpr0_el1"
	case 50756:
		return "icc_ap0r0_el1"
	case 50757:
		return "icc_ap0r1_el1"
	case 50758:
		return "icc_ap0r2_el1"
	case 50759:
		return "icc_ap0r3_el1"
	case 50760:
		return "icc_ap1r0_el1"
	case 50761:
		return "icc_ap1r1_el1"
	case 50762:
		return "icc_ap1r2_el1"
	case 50763:
		return "icc_ap1r3_el1"
	case 50777:
		return "icc_dir_el1"
	case 50781:
		return "icc_sgi1r_el1"
	case 50782:
		return "icc_asgi1r_el1"
	case 50783:
		return "icc_sgi0r_el1"
	case 50785:
		return "icc_eoir1_el1"
	case 50787:
		return "icc_bpr1_el1"
	case 50788:
		return "icc_ctlr_el1"
	case 50789:
		return "icc_sre_el1"
	case 50790:
		return "icc_igrpen0_el1"
	case 50791:
		return "icc_igrpen1_el1"
	case 50792:
		return "icc_seien_el1"
	case 50817:
		return "contextidr_el1"
	case 50820:
		return "tpidr_el1"
	case 50823:
		return "scxtnum_el1"
	case 50952:
		return "cntkctl_el1"
	case 53248:
		return "csselr_el1"
	case 55824:
		return "nzcv"
	case 55825:
		return "daifset"
	case 55829:
		return "dit"
	case 55830:
		return "ssbs"
	case 55831:
		return "tco"
	case 55840:
		return "fpcr"
	case 55841:
		return "fpsr"
	case 55848:
		return "dspsr_el0"
	case 55849:
		return "dlr_el0"
	case 56544:
		return "pmcr_el0"
	case 56545:
		return "pmcntenset_el0"
	case 56546:
		return "pmcntenclr_el0"
	case 56547:
		return "pmovsclr_el0"
	case 56548:
		return "pmswinc_el0"
	case 56549:
		return "pmselr_el0"
	case 56552:
		return "pmccntr_el0"
	case 56553:
		return "pmxevtyper_el0"
	case 56554:
		return "pmxevcntr_el0"
	case 56557:
		return "daifclr"
	case 56560:
		return "pmuserenr_el0"
	case 56563:
		return "pmovsset_el0"
	case 56962:
		return "tpidr_el0"
	case 56963:
		return "tpidrro_el0"
	case 56967:
		return "scxtnum_el0"
	case 56976:
		return "amcr_el0"
	case 56979:
		return "amuserenr_el0"
	case 56980:
		return "amcntenclr0_el0"
	case 56981:
		return "amcntenset0_el0"
	case 56984:
		return "amcntenclr1_el0"
	case 56985:
		return "amcntenset1_el0"
	case 56992:
		return "amevcntr00_el0"
	case 56993:
		return "amevcntr01_el0"
	case 56994:
		return "amevcntr02_el0"
	case 56995:
		return "amevcntr03_el0"
	case 57056:
		return "amevcntr10_el0"
	case 57057:
		return "amevcntr11_el0"
	case 57058:
		return "amevcntr12_el0"
	case 57059:
		return "amevcntr13_el0"
	case 57060:
		return "amevcntr14_el0"
	case 57061:
		return "amevcntr15_el0"
	case 57062:
		return "amevcntr16_el0"
	case 57063:
		return "amevcntr17_el0"
	case 57064:
		return "amevcntr18_el0"
	case 57065:
		return "amevcntr19_el0"
	case 57066:
		return "amevcntr110_el0"
	case 57067:
		return "amevcntr111_el0"
	case 57068:
		return "amevcntr112_el0"
	case 57069:
		return "amevcntr113_el0"
	case 57070:
		return "amevcntr114_el0"
	case 57071:
		return "amevcntr115_el0"
	case 57072:
		return "amevtyper10_el0"
	case 57073:
		return "amevtyper11_el0"
	case 57074:
		return "amevtyper12_el0"
	case 57075:
		return "amevtyper13_el0"
	case 57076:
		return "amevtyper14_el0"
	case 57077:
		return "amevtyper15_el0"
	case 57078:
		return "amevtyper16_el0"
	case 57079:
		return "amevtyper17_el0"
	case 57080:
		return "amevtyper18_el0"
	case 57081:
		return "amevtyper19_el0"
	case 57082:
		return "amevtyper110_el0"
	case 57083:
		return "amevtyper111_el0"
	case 57084:
		return "amevtyper112_el0"
	case 57085:
		return "amevtyper113_el0"
	case 57086:
		return "amevtyper114_el0"
	case 57087:
		return "amevtyper115_el0"
	case 57088:
		return "cntfrq_el0"
	case 57104:
		return "cntp_tval_el0"
	case 57105:
		return "cntp_ctl_el0"
	case 57106:
		return "cntp_cval_el0"
	case 57112:
		return "cntv_tval_el0"
	case 57113:
		return "cntv_ctl_el0"
	case 57114:
		return "cntv_cval_el0"
	case 57152:
		return "pmevcntr0_el0"
	case 57153:
		return "pmevcntr1_el0"
	case 57154:
		return "pmevcntr2_el0"
	case 57155:
		return "pmevcntr3_el0"
	case 57156:
		return "pmevcntr4_el0"
	case 57157:
		return "pmevcntr5_el0"
	case 57158:
		return "pmevcntr6_el0"
	case 57159:
		return "pmevcntr7_el0"
	case 57160:
		return "pmevcntr8_el0"
	case 57161:
		return "pmevcntr9_el0"
	case 57162:
		return "pmevcntr10_el0"
	case 57163:
		return "pmevcntr11_el0"
	case 57164:
		return "pmevcntr12_el0"
	case 57165:
		return "pmevcntr13_el0"
	case 57166:
		return "pmevcntr14_el0"
	case 57167:
		return "pmevcntr15_el0"
	case 57168:
		return "pmevcntr16_el0"
	case 57169:
		return "pmevcntr17_el0"
	case 57170:
		return "pmevcntr18_el0"
	case 57171:
		return "pmevcntr19_el0"
	case 57172:
		return "pmevcntr20_el0"
	case 57173:
		return "pmevcntr21_el0"
	case 57174:
		return "pmevcntr22_el0"
	case 57175:
		return "pmevcntr23_el0"
	case 57176:
		return "pmevcntr24_el0"
	case 57177:
		return "pmevcntr25_el0"
	case 57178:
		return "pmevcntr26_el0"
	case 57179:
		return "pmevcntr27_el0"
	case 57180:
		return "pmevcntr28_el0"
	case 57181:
		return "pmevcntr29_el0"
	case 57182:
		return "pmevcntr30_el0"
	case 57184:
		return "pmevtyper0_el0"
	case 57185:
		return "pmevtyper1_el0"
	case 57186:
		return "pmevtyper2_el0"
	case 57187:
		return "pmevtyper3_el0"
	case 57188:
		return "pmevtyper4_el0"
	case 57189:
		return "pmevtyper5_el0"
	case 57190:
		return "pmevtyper6_el0"
	case 57191:
		return "pmevtyper7_el0"
	case 57192:
		return "pmevtyper8_el0"
	case 57193:
		return "pmevtyper9_el0"
	case 57194:
		return "pmevtyper10_el0"
	case 57195:
		return "pmevtyper11_el0"
	case 57196:
		return "pmevtyper12_el0"
	case 57197:
		return "pmevtyper13_el0"
	case 57198:
		return "pmevtyper14_el0"
	case 57199:
		return "pmevtyper15_el0"
	case 57200:
		return "pmevtyper16_el0"
	case 57201:
		return "pmevtyper17_el0"
	case 57202:
		return "pmevtyper18_el0"
	case 57203:
		return "pmevtyper19_el0"
	case 57204:
		return "pmevtyper20_el0"
	case 57205:
		return "pmevtyper21_el0"
	case 57206:
		return "pmevtyper22_el0"
	case 57207:
		return "pmevtyper23_el0"
	case 57208:
		return "pmevtyper24_el0"
	case 57209:
		return "pmevtyper25_el0"
	case 57210:
		return "pmevtyper26_el0"
	case 57211:
		return "pmevtyper27_el0"
	case 57212:
		return "pmevtyper28_el0"
	case 57213:
		return "pmevtyper29_el0"
	case 57214:
		return "pmevtyper30_el0"
	case 57215:
		return "pmccfiltr_el0"
	case 57344:
		return "vpidr_el2"
	case 57349:
		return "vmpidr_el2"
	case 57472:
		return "sctlr_el2"
	case 57473:
		return "actlr_el2"
	case 57480:
		return "hcr_el2"
	case 57481:
		return "mdcr_el2"
	case 57482:
		return "cptr_el2"
	case 57483:
		return "hstr_el2"
	case 57487:
		return "hacr_el2"
	case 57489:
		return "trfcr_el2"
	case 57497:
		return "sder32_el2"
	case 57600:
		return "ttbr0_el2"
	case 57601:
		return "ttbr1_el2"
	case 57602:
		return "tcr_el2"
	case 57608:
		return "vttbr_el2"
	case 57610:
		return "vtcr_el2"
	case 57616:
		return "vncr_el2"
	case 57648:
		return "vsttbr_el2"
	case 57650:
		return "vstcr_el2"
	case 57728:
		return "dacr32_el2"
	case 57856:
		return "spsr_el2"
	case 57857:
		return "elr_el2"
	case 57864:
		return "sp_el1"
	case 57880:
		return "spsr_irq"
	case 57881:
		return "spsr_abt"
	case 57882:
		return "spsr_und"
	case 57883:
		return "spsr_fiq"
	case 57985:
		return "ifsr32_el2"
	case 57992:
		return "afsr0_el2"
	case 57993:
		return "afsr1_el2"
	case 58000:
		return "esr_el2"
	case 58003:
		return "vsesr_el2"
	case 58008:
		return "fpexc32_el2"
	case 58032:
		return "tfsr_el2"
	case 58112:
		return "far_el2"
	case 58116:
		return "hpfar_el2"
	case 58568:
		return "pmscr_el2"
	case 58640:
		return "mair_el2"
	case 58648:
		return "amair_el2"
	case 58656:
		return "mpamhcr_el2"
	case 58657:
		return "mpamvpmv_el2"
	case 58664:
		return "mpam2_el2"
	case 58672:
		return "mpamvpm0_el2"
	case 58673:
		return "mpamvpm1_el2"
	case 58674:
		return "mpamvpm2_el2"
	case 58675:
		return "mpamvpm3_el2"
	case 58676:
		return "mpamvpm4_el2"
	case 58677:
		return "mpamvpm5_el2"
	case 58678:
		return "mpamvpm6_el2"
	case 58679:
		return "mpamvpm7_el2"
	case 58880:
		return "vbar_el2"
	case 58882:
		return "rmr_el2"
	case 58889:
		return "vdisr_el2"
	case 58944:
		return "ich_ap0r0_el2"
	case 58945:
		return "ich_ap0r1_el2"
	case 58946:
		return "ich_ap0r2_el2"
	case 58947:
		return "ich_ap0r3_el2"
	case 58952:
		return "ich_ap1r0_el2"
	case 58953:
		return "ich_ap1r1_el2"
	case 58954:
		return "ich_ap1r2_el2"
	case 58955:
		return "ich_ap1r3_el2"
	case 58956:
		return "ich_vseir_el2"
	case 58957:
		return "icc_sre_el2"
	case 58968:
		return "ich_hcr_el2"
	case 58970:
		return "ich_misr_el2"
	case 58975:
		return "ich_vmcr_el2"
	case 58976:
		return "ich_lr0_el2"
	case 58977:
		return "ich_lr1_el2"
	case 58978:
		return "ich_lr2_el2"
	case 58979:
		return "ich_lr3_el2"
	case 58980:
		return "ich_lr4_el2"
	case 58981:
		return "ich_lr5_el2"
	case 58982:
		return "ich_lr6_el2"
	case 58983:
		return "ich_lr7_el2"
	case 58984:
		return "ich_lr8_el2"
	case 58985:
		return "ich_lr9_el2"
	case 58986:
		return "ich_lr10_el2"
	case 58987:
		return "ich_lr11_el2"
	case 58988:
		return "ich_lr12_el2"
	case 58989:
		return "ich_lr13_el2"
	case 58990:
		return "ich_lr14_el2"
	case 58991:
		return "ich_lr15_el2"
	case 59009:
		return "contextidr_el2"
	case 59010:
		return "tpidr_el2"
	case 59015:
		return "scxtnum_el2"
	case 59139:
		return "cntvoff_el2"
	case 59144:
		return "cnthctl_el2"
	case 59152:
		return "cnthp_tval_el2"
	case 59153:
		return "cnthp_ctl_el2"
	case 59154:
		return "cnthp_cval_el2"
	case 59160:
		return "cnthv_tval_el2"
	case 59161:
		return "cnthv_ctl_el2"
	case 59162:
		return "cnthv_cval_el2"
	case 59168:
		return "cnthvs_tval_el2"
	case 59169:
		return "cnthvs_ctl_el2"
	case 59170:
		return "cnthvs_cval_el2"
	case 59176:
		return "cnthps_tval_el2"
	case 59177:
		return "cnthps_ctl_el2"
	case 59178:
		return "cnthps_cval_el2"
	case 59520:
		return "sctlr_el12"
	case 59522:
		return "cpacr_el12"
	case 59537:
		return "trfcr_el12"
	case 59648:
		return "ttbr0_el12"
	case 59649:
		return "ttbr1_el12"
	case 59650:
		return "tcr_el12"
	case 59904:
		return "spsr_el12"
	case 59905:
		return "elr_el12"
	case 60040:
		return "afsr0_el12"
	case 60041:
		return "afsr1_el12"
	case 60048:
		return "esr_el12"
	case 60080:
		return "tfsr_el12"
	case 60160:
		return "far_el12"
	case 60616:
		return "pmscr_el12"
	case 60688:
		return "mair_el12"
	case 60696:
		return "amair_el12"
	case 60712:
		return "mpam1_el12"
	case 60928:
		return "vbar_el12"
	case 61057:
		return "contextidr_el12"
	case 61063:
		return "scxtnum_el12"
	case 61192:
		return "cntkctl_el12"
	case 61200:
		return "cntp_tval_el02"
	case 61201:
		return "cntp_ctl_el02"
	case 61202:
		return "cntp_cval_el02"
	case 61208:
		return "cntv_tval_el02"
	case 61209:
		return "cntv_ctl_el02"
	case 61210:
		return "cntv_cval_el02"
	case 61568:
		return "sctlr_el3"
	case 61569:
		return "actlr_el3"
	case 61576:
		return "scr_el3"
	case 61577:
		return "sder32_el3"
	case 61578:
		return "cptr_el3"
	case 61593:
		return "mdcr_el3"
	case 61696:
		return "ttbr0_el3"
	case 61698:
		return "tcr_el3"
	case 61952:
		return "spsr_el3"
	case 61953:
		return "elr_el3"
	case 61960:
		return "sp_el2"
	case 62088:
		return "afsr0_el3"
	case 62089:
		return "afsr1_el3"
	case 62096:
		return "esr_el3"
	case 62128:
		return "tfsr_el3"
	case 62208:
		return "far_el3"
	case 62736:
		return "mair_el3"
	case 62744:
		return "amair_el3"
	case 62760:
		return "mpam3_el3"
	case 62976:
		return "vbar_el3"
	case 62978:
		return "rmr_el3"
	case 63076:
		return "icc_ctlr_el3"
	case 63077:
		return "icc_sre_el3"
	case 63079:
		return "icc_igrpen1_el3"
	case 63106:
		return "tpidr_el3"
	case 63111:
		return "scxtnum_el3"
	case 65296:
		return "cntps_tval_el1"
	case 65297:
		return "cntps_ctl_el1"
	case 65298:
		return "cntps_cval_el1"
	case 32896: /* From https://github.com/AsahiLinux/m1n1 */
		return "mdrar_el1"
	case 32908:
		return "oslsr_el1"
	case 33782:
		return "dbgauthstatus_el1"
	case 38920:
		return "mdccsr_el0"
	case 49152:
		return "midr_el1"
	case 49157:
		return "mpidr_el1"
	case 49158:
		return "revidr_el1"
	case 49160:
		return "id_pfr0_el1"
	case 49161:
		return "id_pfr1_el1"
	case 49162:
		return "id_dfr0_el1"
	case 49163:
		return "id_afr0_el1"
	case 49164:
		return "id_mmfr0_el1"
	case 49165:
		return "id_mmfr1_el1"
	case 49166:
		return "id_mmfr2_el1"
	case 49167:
		return "id_mmfr3_el1"
	case 49168:
		return "id_isar0_el1"
	case 49169:
		return "id_isar1_el1"
	case 49170:
		return "id_isar2_el1"
	case 49171:
		return "id_isar3_el1"
	case 49172:
		return "id_isar4_el1"
	case 49173:
		return "id_isar5_el1"
	case 49174:
		return "id_mmfr4_el1"
	case 49175:
		return "id_isar6_el1"
	case 49176:
		return "mvfr0_el1"
	case 49177:
		return "mvfr1_el1"
	case 49178:
		return "mvfr2_el1"
	case 49180:
		return "id_pfr2_el1"
	case 49181:
		return "id_dfr1_el1"
	case 49182:
		return "id_mmfr5_el1"
	case 49184:
		return "id_aa64pfr0_el1"
	case 49185:
		return "id_aa64pfr1_el1"
	case 49188:
		return "id_aa64zfr0_el1"
	case 49192:
		return "id_aa64dfr0_el1"
	case 49193:
		return "id_aa64dfr1_el1"
	case 49196:
		return "id_aa64afr0_el1"
	case 49197:
		return "id_aa64afr1_el1"
	case 49200:
		return "id_aa64isar0_el1"
	case 49201:
		return "id_aa64isar1_el1"
	case 49202:
		return "id_aa64isar2_el1"
	case 49208:
		return "id_aa64mmfr0_el1"
	case 49209:
		return "id_aa64mmfr1_el1"
	case 49210:
		return "id_aa64mmfr2_el1"
	case 49296:
		return "zcr_el1"
	case 49816:
		return "erridr_el1"
	case 49824:
		return "erxfr_el1"
	case 49828:
		return "erxpfgf_el1"
	case 50377:
		return "pmsnevfr_el1"
	case 50468:
		return "mpamidr_el1"
	case 50471:
		return "lorid_el1"
	case 50689:
		return "rvbar_el1"
	case 50696:
		return "isr_el1"
	case 50752:
		return "icc_iar0_el1"
	case 50754:
		return "icc_hppir0_el1"
	case 50779:
		return "icc_rpr_el1"
	case 50784:
		return "icc_iar1_el1"
	case 50786:
		return "icc_hppir1_el1"
	case 50821:
		return "accdata_el1"
	case 51072:
		return "hid0_el1"
	case 51073:
		return "ehid0_el1"
	case 51080:
		return "hid1_el1"
	case 51081:
		return "ehid1_el1"
	case 51082:
		return "ehid20_el1"
	case 51083:
		return "hid21_el1"
	case 51088:
		return "hid2_el1"
	case 51089:
		return "ehid2_el1"
	case 51096:
		return "hid3_el1"
	case 51097:
		return "ehid3_el1"
	case 51104:
		return "hid4_el1"
	case 51105:
		return "ehid4_el1"
	case 51112:
		return "hid5_el1"
	case 51113:
		return "ehid5_el1"
	case 51120:
		return "hid6_el1"
	case 51128:
		return "hid7_el1"
	case 51129:
		return "ehid7_el1"
	case 51136:
		return "hid8_el1"
	case 51144:
		return "hid9_el1"
	case 51145:
		return "ehid9_el1"
	case 51152:
		return "hid10_el1"
	case 51153:
		return "ehid10_el1"
	case 51160:
		return "hid11_el1"
	case 51161:
		return "ehid11_el1"
	case 51162:
		return "hid18_el1"
	case 51184:
		return "hid13_el1"
	case 51192:
		return "hid14_el1"
	case 51194:
		return "hid16_el1"
	case 51197:
		return "hid17_el1"
	case 51200:
		return "ccsidr_el1"
	case 51201:
		return "clidr_el1"
	case 51202:
		return "ccsidr2_el1"
	case 51204:
		return "gmid_el1"
	case 51207:
		return "aidr_el1"
	case 53120:
		return "pmcr0_el1"
	case 53128:
		return "pmcr1_el1"
	case 53136:
		return "pmcr2_el1"
	case 53144:
		return "pmcr3_el1"
	case 53152:
		return "pmcr4_el1"
	case 53160:
		return "pmesr0_el1"
	case 53168:
		return "pmesr1_el1"
	case 53224:
		return "pmsr_el1"
	case 55168:
		return "pmc0_el1"
	case 55176:
		return "pmc1_el1"
	case 55184:
		return "pmc2_el1"
	case 55192:
		return "pmc3_el1"
	case 55200:
		return "pmc4_el1"
	case 55208:
		return "pmc5_el1"
	case 55216:
		return "pmc6_el1"
	case 55224:
		return "pmc7_el1"
	case 55240:
		return "pmc8_el1"
	case 55248:
		return "pmc9_el1"
	case 55297:
		return "ctr_el0"
	case 55303:
		return "dczid_el0"
	case 55584:
		return "rndr"
	case 55585:
		return "rndrrs"
	case 56550:
		return "pmceid0_el0"
	case 56551:
		return "pmceid1_el0"
	case 56977:
		return "amcfgr_el0"
	case 56978:
		return "amcgcr_el0"
	case 56982:
		return "amcg1idr_el0"
	case 57008:
		return "amevtyper00_el0"
	case 57089:
		return "cntpct_el0"
	case 57090:
		return "cntvct_el0"
	case 57093:
		return "cntpctss_el0"
	case 57094:
		return "cntvctss_el0"
	case 57216:
		return "lsu_err_sts_el1"
	case 57224:
		return "lsu_err_ctl_el1"
	case 57232:
		return "e_lsu_err_sts_el1"
	case 57280:
		return "l2c_err_sts_el1"
	case 57288:
		return "l2c_err_adr_el1"
	case 57296:
		return "l2c_err_inf_el1"
	case 57484:
		return "hfgrtr_el2"
	case 57485:
		return "hfgwtr_el2"
	case 57486:
		return "hfgitr_el2"
	case 57488:
		return "zcr_el2"
	case 57490:
		return "hcrx_el2"
	case 57740:
		return "hdfgrtr_el2"
	case 57741:
		return "hdfgwtr_el2"
	case 57742:
		return "hafgrtr_el2"
	case 58881:
		return "rvbar_el2"
	case 58969:
		return "ich_vtr_el2"
	case 58971:
		return "ich_eisr_el2"
	case 58973:
		return "ich_elrsr_el2"
	case 59072:
		return "amevcntvoff00_el2"
	case 59088:
		return "amevcntvoff10_el2"
	case 59142:
		return "cntpoff_el2"
	case 59264:
		return "fed_err_sts_el1"
	case 59266:
		return "e_fed_err_sts_el1"
	case 59268:
		return "apctl_el1"
	case 59272:
		return "kernelkeylo_el1"
	case 59273:
		return "kernelkeyhi_el1"
	case 59274:
		return "vmsa_lock_el1"
	case 59276:
		return "amx_ctl_el1"
	case 59280:
		return "aprr_el0"
	case 59281:
		return "aprr_el1"
	case 59282:
		return "ctrr_lock_el1"
	case 59283:
		return "ctrr_a_lwr_el1"
	case 59284:
		return "ctrr_a_upr_el1"
	case 59285:
		return "ctrr_ctl_el1"
	case 59286:
		return "aprr_jit_enable_el2"
	case 59287:
		return "aprr_jit_mask_el2"
	case 59302:
		return "amx_ctl_el12"
	case 59303:
		return "amx_ctl_el2"
	case 59305:
		return "sprr_perm_el20_silly_thing"
	case 59306:
		return "sprr_perm_el02"
	case 59320:
		return "sprr_umask0_el2"
	case 59321:
		return "sprr_umask1_el2"
	case 59322:
		return "sprr_umask2_el2"
	case 59323:
		return "sprr_umask3_el2"
	case 59328:
		return "sprr_umask0_el12"
	case 59329:
		return "sprr_umask1_el12"
	case 59330:
		return "sprr_umask2_el12"
	case 59331:
		return "sprr_umask3_el12"
	case 59349:
		return "cntpct_alias_el0"
	case 59350:
		return "cntvct_alias_el0"
	case 59352:
		return "ctrr_a_lwr_el2"
	case 59353:
		return "ctrr_a_upr_el2"
	case 59356:
		return "ctrr_ctl_el2"
	case 59357:
		return "ctrr_lock_el2"
	case 59536:
		return "zcr_el12"
	case 61312:
		return "ipi_rr_local_el1"
	case 61313:
		return "ipi_rr_global_el1"
	case 61317:
		return "dpc_err_sts_el1"
	case 61321:
		return "ipi_sr_el1"
	case 61322:
		return "vm_tmr_lr_el2"
	case 61323:
		return "vm_tmr_fiq_ena_el2"
	case 61337:
		return "ipi_cr_el1"
	case 61344:
		return "acc_cfg_el1"
	case 61352:
		return "cyc_ovrd_el1"
	case 61360:
		return "acc_ovrd_el1"
	case 61361:
		return "acc_eblk_ovrd_el1"
	case 61584:
		return "zcr_el3"
	case 62977:
		return "rvbar_el3"
	case 63360:
		return "mmu_err_sts_el1"
	case 63361:
		return "afsr1_gl1"
	case 63362:
		return "afsr1_gl2"
	case 63363:
		return "afsr1_gl12"
	case 63368:
		return "sprr_config_el1"
	case 63370:
		return "gxf_config_el1"
	case 63371:
		return "sprr_unk1_el1"
	case 63372:
		return "gxf_config_el2"
	case 63373:
		return "sprr_perm_el0"
	case 63374:
		return "sprr_perm_el1"
	case 63375:
		return "sprr_perm_el2"
	case 63376:
		return "e_mmu_err_sts_el1"
	case 63377:
		return "apgakeylo_el12"
	case 63378:
		return "apgakeyhi_el12"
	case 63379:
		return "kernelkeylo_el12"
	case 63380:
		return "kernelkeyhi_el12"
	case 63381:
		return "afpcr_el0"
	case 63383:
		return "aidr2_el1"
	case 63384:
		return "sprr_umask0_el1"
	case 63385:
		return "sprr_kmask0_el1"
	case 63386:
		return "sprr_kmask0_el2"
	case 63387:
		return "sprr_umask1_el1"
	case 63388:
		return "sprr_umask2_el1"
	case 63389:
		return "sprr_umask3_el1"
	case 63394:
		return "sprr_kmask1_el1"
	case 63395:
		return "sprr_kmask2_el1"
	case 63396:
		return "sprr_kmask3_el1"
	case 63401:
		return "sprr_kmask1_el2"
	case 63402:
		return "sprr_kmask2_el2"
	case 63403:
		return "sprr_kmask3_el2"
	case 63408:
		return "sprr_kmask0_el12"
	case 63409:
		return "sprr_kmask1_el12"
	case 63410:
		return "sprr_kmask2_el12"
	case 63411:
		return "sprr_kmask3_el12"
	case 63416:
		return "apiakeylo_el12"
	case 63417:
		return "apiakeyhi_el12"
	case 63418:
		return "apibkeylo_el12"
	case 63419:
		return "apibkeyhi_el12"
	case 63420:
		return "apdakeylo_el12"
	case 63421:
		return "apdakeyhi_el12"
	case 63422:
		return "apdbkeylo_el12"
	case 63423:
		return "apdbkeyhi_el12"
	case 63424:
		return "gxf_status_el1"
	case 63425:
		return "gxf_enter_el1"
	case 63426:
		return "gxf_abort_el1"
	case 63434:
		return "vbar_gl12"
	case 63435:
		return "spsr_gl12"
	case 63436:
		return "aspsr_gl12"
	case 63437:
		return "esr_gl12"
	case 63438:
		return "elr_gl12"
	case 63440:
		return "sp_gl12"
	case 63441:
		return "tpidr_gl1"
	case 63442:
		return "vbar_gl1"
	case 63443:
		return "spsr_gl1"
	case 63444:
		return "aspsr_gl1"
	case 63445:
		return "esr_gl1"
	case 63446:
		return "elr_gl1"
	case 63447:
		return "far_gl1"
	case 63449:
		return "tpidr_gl2"
	case 63450:
		return "vbar_gl2"
	case 63451:
		return "spsr_gl2"
	case 63452:
		return "aspsr_gl2"
	case 63453:
		return "esr_gl2"
	case 63454:
		return "elr_gl2"
	case 63455:
		return "far_gl2"
	case 63456:
		return "gxf_enter_el2"
	case 63457:
		return "gxf_abort_el2"
	case 63458:
		return "apctl_el2"
	case 63459:
		return "apsts_el2_maybe"
	case 63460:
		return "apsts_el1"
	case 63474:
		return "sprr_config_el2"
	case 63475:
		return "sprr_unk1_el2"
	case 63476:
		return "apvmkeylo_el2"
	case 63477:
		return "apvmkeyhi_el2"
	case 63478:
		return "actlr_el12"
	case 63479:
		return "apsts_el12"
	case 63480:
		return "apctl_el12"
	case 63481:
		return "gxf_config_el12"
	case 63482:
		return "gxf_enter_el12"
	case 63483:
		return "gxf_abort_el12"
	case 63484:
		return "sprr_config_el12"
	case 63485:
		return "sprr_unk1_el12"
	case 63487:
		return "sprr_perm_el12"
	case 65412:
		return "upmcr0_el1"
	case 65413:
		return "upmc8_el1"
	case 65420:
		return "upmesr0_el1"
	case 65421:
		return "upmc9_el1"
	case 65429:
		return "upmc10_el1"
	case 65436:
		return "upmecm0_el1"
	case 65437:
		return "upmc11_el1"
	case 65444:
		return "upmecm1_el1"
	case 65445:
		return "upmc12_el1"
	case 65452:
		return "upmpcm_el1"
	case 65453:
		return "upmc13_el1"
	case 65460:
		return "upmsr_el1"
	case 65461:
		return "upmc14_el1"
	case 65468:
		return "upmc0_el1"
	case 65469:
		return "upmc15_el1"
	case 65476:
		return "upmc1_el1"
	case 65477:
		return "upmecm2_el1"
	case 65484:
		return "upmc2_el1"
	case 65485:
		return "upmecm3_el1"
	case 65492:
		return "upmc3_el1"
	case 65500:
		return "upmc4_el1"
	case 65501:
		return "upmesr1_el1"
	case 65508:
		return "upmc5_el1"
	case 65516:
		return "upmc6_el1"
	case 65524:
		return "upmc7_el1"
	case 65299:
		return "spsel"
	case 65300:
		return "unknown_catchall"
	default:
		return ""
	}
}
