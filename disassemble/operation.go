package disassemble

type operation uint32

const (
	ARM64_ERROR     operation = 0
	ARM64_ABS       operation = 1
	ARM64_ADC       operation = 2
	ARM64_ADCLB     operation = 3
	ARM64_ADCLT     operation = 4
	ARM64_ADCS      operation = 5
	ARM64_ADD       operation = 6
	ARM64_ADDG      operation = 7
	ARM64_ADDHA     operation = 8
	ARM64_ADDHN     operation = 9
	ARM64_ADDHN2    operation = 10
	ARM64_ADDHNB    operation = 11
	ARM64_ADDHNT    operation = 12
	ARM64_ADDP      operation = 13
	ARM64_ADDPL     operation = 14
	ARM64_ADDS      operation = 15
	ARM64_ADDV      operation = 16
	ARM64_ADDVA     operation = 17
	ARM64_ADDVL     operation = 18
	ARM64_ADR       operation = 19
	ARM64_ADRP      operation = 20
	ARM64_AESD      operation = 21
	ARM64_AESE      operation = 22
	ARM64_AESIMC    operation = 23
	ARM64_AESMC     operation = 24
	ARM64_AND       operation = 25
	ARM64_ANDS      operation = 26
	ARM64_ANDV      operation = 27
	ARM64_ASR       operation = 28
	ARM64_ASRD      operation = 29
	ARM64_ASRR      operation = 30
	ARM64_ASRV      operation = 31
	ARM64_AT        operation = 32
	ARM64_AUTDA     operation = 33
	ARM64_AUTDB     operation = 34
	ARM64_AUTDZA    operation = 35
	ARM64_AUTDZB    operation = 36
	ARM64_AUTIA     operation = 37
	ARM64_AUTIA1716 operation = 38
	ARM64_AUTIASP   operation = 39
	ARM64_AUTIAZ    operation = 40
	ARM64_AUTIB     operation = 41
	ARM64_AUTIB1716 operation = 42
	ARM64_AUTIBSP   operation = 43
	ARM64_AUTIBZ    operation = 44
	ARM64_AUTIZA    operation = 45
	ARM64_AUTIZB    operation = 46
	ARM64_AXFLAG    operation = 47
	ARM64_B         operation = 48
	ARM64_BCAX      operation = 49
	ARM64_BDEP      operation = 50
	ARM64_BEXT      operation = 51
	ARM64_BFC       operation = 52
	ARM64_BFCVT     operation = 53
	ARM64_BFCVTN    operation = 54
	ARM64_BFCVTN2   operation = 55
	ARM64_BFCVTNT   operation = 56
	ARM64_BFDOT     operation = 57
	ARM64_BFI       operation = 58
	ARM64_BFM       operation = 59
	ARM64_BFMLAL    operation = 60
	ARM64_BFMLALB   operation = 61
	ARM64_BFMLALT   operation = 62
	ARM64_BFMMLA    operation = 63
	ARM64_BFMOPA    operation = 64
	ARM64_BFMOPS    operation = 65
	ARM64_BFXIL     operation = 66
	ARM64_BGRP      operation = 67
	ARM64_BIC       operation = 68
	ARM64_BICS      operation = 69
	ARM64_BIF       operation = 70
	ARM64_BIT       operation = 71
	ARM64_BL        operation = 72
	ARM64_BLR       operation = 73
	ARM64_BLRAA     operation = 74
	ARM64_BLRAAZ    operation = 75
	ARM64_BLRAB     operation = 76
	ARM64_BLRABZ    operation = 77
	ARM64_BR        operation = 78
	ARM64_BRAA      operation = 79
	ARM64_BRAAZ     operation = 80
	ARM64_BRAB      operation = 81
	ARM64_BRABZ     operation = 82
	ARM64_BRK       operation = 83
	ARM64_BRKA      operation = 84
	ARM64_BRKAS     operation = 85
	ARM64_BRKB      operation = 86
	ARM64_BRKBS     operation = 87
	ARM64_BRKN      operation = 88
	ARM64_BRKNS     operation = 89
	ARM64_BRKPA     operation = 90
	ARM64_BRKPAS    operation = 91
	ARM64_BRKPB     operation = 92
	ARM64_BRKPBS    operation = 93
	ARM64_BSL       operation = 94
	ARM64_BSL1N     operation = 95
	ARM64_BSL2N     operation = 96
	ARM64_BTI       operation = 97
	ARM64_B_AL      operation = 98
	ARM64_B_CC      operation = 99
	ARM64_B_CS      operation = 100
	ARM64_B_EQ      operation = 101
	ARM64_B_GE      operation = 102
	ARM64_B_GT      operation = 103
	ARM64_B_HI      operation = 104
	ARM64_B_LE      operation = 105
	ARM64_B_LS      operation = 106
	ARM64_B_LT      operation = 107
	ARM64_B_MI      operation = 108
	ARM64_B_NE      operation = 109
	ARM64_B_NV      operation = 110
	ARM64_B_PL      operation = 111
	ARM64_B_VC      operation = 112
	ARM64_B_VS      operation = 113
	ARM64_CADD      operation = 114
	ARM64_CAS       operation = 115
	ARM64_CASA      operation = 116
	ARM64_CASAB     operation = 117
	ARM64_CASAH     operation = 118
	ARM64_CASAL     operation = 119
	ARM64_CASALB    operation = 120
	ARM64_CASALH    operation = 121
	ARM64_CASB      operation = 122
	ARM64_CASH      operation = 123
	ARM64_CASL      operation = 124
	ARM64_CASLB     operation = 125
	ARM64_CASLH     operation = 126
	ARM64_CASP      operation = 127
	ARM64_CASPA     operation = 128
	ARM64_CASPAL    operation = 129
	ARM64_CASPL     operation = 130
	ARM64_CBNZ      operation = 131
	ARM64_CBZ       operation = 132
	ARM64_CCMN      operation = 133
	ARM64_CCMP      operation = 134
	ARM64_CDOT      operation = 135
	ARM64_CFINV     operation = 136
	ARM64_CFP       operation = 137
	ARM64_CINC      operation = 138
	ARM64_CINV      operation = 139
	ARM64_CLASTA    operation = 140
	ARM64_CLASTB    operation = 141
	ARM64_CLREX     operation = 142
	ARM64_CLS       operation = 143
	ARM64_CLZ       operation = 144
	ARM64_CMEQ      operation = 145
	ARM64_CMGE      operation = 146
	ARM64_CMGT      operation = 147
	ARM64_CMHI      operation = 148
	ARM64_CMHS      operation = 149
	ARM64_CMLA      operation = 150
	ARM64_CMLE      operation = 151
	ARM64_CMLT      operation = 152
	ARM64_CMN       operation = 153
	ARM64_CMP       operation = 154
	ARM64_CMPEQ     operation = 155
	ARM64_CMPGE     operation = 156
	ARM64_CMPGT     operation = 157
	ARM64_CMPHI     operation = 158
	ARM64_CMPHS     operation = 159
	ARM64_CMPLE     operation = 160
	ARM64_CMPLO     operation = 161
	ARM64_CMPLS     operation = 162
	ARM64_CMPLT     operation = 163
	ARM64_CMPNE     operation = 164
	ARM64_CMPP      operation = 165
	ARM64_CMTST     operation = 166
	ARM64_CNEG      operation = 167
	ARM64_CNOT      operation = 168
	ARM64_CNT       operation = 169
	ARM64_CNTB      operation = 170
	ARM64_CNTD      operation = 171
	ARM64_CNTH      operation = 172
	ARM64_CNTP      operation = 173
	ARM64_CNTW      operation = 174
	ARM64_COMPACT   operation = 175
	ARM64_CPP       operation = 176
	ARM64_CPY       operation = 177
	ARM64_CRC32B    operation = 178
	ARM64_CRC32CB   operation = 179
	ARM64_CRC32CH   operation = 180
	ARM64_CRC32CW   operation = 181
	ARM64_CRC32CX   operation = 182
	ARM64_CRC32H    operation = 183
	ARM64_CRC32W    operation = 184
	ARM64_CRC32X    operation = 185
	ARM64_CSDB      operation = 186
	ARM64_CSEL      operation = 187
	ARM64_CSET      operation = 188
	ARM64_CSETM     operation = 189
	ARM64_CSINC     operation = 190
	ARM64_CSINV     operation = 191
	ARM64_CSNEG     operation = 192
	ARM64_CTERMEQ   operation = 193
	ARM64_CTERMNE   operation = 194
	ARM64_DC        operation = 195
	ARM64_DCPS1     operation = 196
	ARM64_DCPS2     operation = 197
	ARM64_DCPS3     operation = 198
	ARM64_DECB      operation = 199
	ARM64_DECD      operation = 200
	ARM64_DECH      operation = 201
	ARM64_DECP      operation = 202
	ARM64_DECW      operation = 203
	ARM64_DGH       operation = 204
	ARM64_DMB       operation = 205
	ARM64_DRPS      operation = 206
	ARM64_DSB       operation = 207
	ARM64_DUP       operation = 208
	ARM64_DUPM      operation = 209
	ARM64_DVP       operation = 210
	ARM64_EON       operation = 211
	ARM64_EOR       operation = 212
	ARM64_EOR3      operation = 213
	ARM64_EORBT     operation = 214
	ARM64_EORS      operation = 215
	ARM64_EORTB     operation = 216
	ARM64_EORV      operation = 217
	ARM64_ERET      operation = 218
	ARM64_ERETAA    operation = 219
	ARM64_ERETAB    operation = 220
	ARM64_ESB       operation = 221
	ARM64_EXT       operation = 222
	ARM64_EXTR      operation = 223
	ARM64_FABD      operation = 224
	ARM64_FABS      operation = 225
	ARM64_FACGE     operation = 226
	ARM64_FACGT     operation = 227
	ARM64_FACLE     operation = 228
	ARM64_FACLT     operation = 229
	ARM64_FADD      operation = 230
	ARM64_FADDA     operation = 231
	ARM64_FADDP     operation = 232
	ARM64_FADDV     operation = 233
	ARM64_FCADD     operation = 234
	ARM64_FCCMP     operation = 235
	ARM64_FCCMPE    operation = 236
	ARM64_FCMEQ     operation = 237
	ARM64_FCMGE     operation = 238
	ARM64_FCMGT     operation = 239
	ARM64_FCMLA     operation = 240
	ARM64_FCMLE     operation = 241
	ARM64_FCMLT     operation = 242
	ARM64_FCMNE     operation = 243
	ARM64_FCMP      operation = 244
	ARM64_FCMPE     operation = 245
	ARM64_FCMUO     operation = 246
	ARM64_FCPY      operation = 247
	ARM64_FCSEL     operation = 248
	ARM64_FCVT      operation = 249
	ARM64_FCVTAS    operation = 250
	ARM64_FCVTAU    operation = 251
	ARM64_FCVTL     operation = 252
	ARM64_FCVTL2    operation = 253
	ARM64_FCVTLT    operation = 254
	ARM64_FCVTMS    operation = 255
	ARM64_FCVTMU    operation = 256
	ARM64_FCVTN     operation = 257
	ARM64_FCVTN2    operation = 258
	ARM64_FCVTNS    operation = 259
	ARM64_FCVTNT    operation = 260
	ARM64_FCVTNU    operation = 261
	ARM64_FCVTPS    operation = 262
	ARM64_FCVTPU    operation = 263
	ARM64_FCVTX     operation = 264
	ARM64_FCVTXN    operation = 265
	ARM64_FCVTXN2   operation = 266
	ARM64_FCVTXNT   operation = 267
	ARM64_FCVTZS    operation = 268
	ARM64_FCVTZU    operation = 269
	ARM64_FDIV      operation = 270
	ARM64_FDIVR     operation = 271
	ARM64_FDUP      operation = 272
	ARM64_FEXPA     operation = 273
	ARM64_FJCVTZS   operation = 274
	ARM64_FLOGB     operation = 275
	ARM64_FMAD      operation = 276
	ARM64_FMADD     operation = 277
	ARM64_FMAX      operation = 278
	ARM64_FMAXNM    operation = 279
	ARM64_FMAXNMP   operation = 280
	ARM64_FMAXNMV   operation = 281
	ARM64_FMAXP     operation = 282
	ARM64_FMAXV     operation = 283
	ARM64_FMIN      operation = 284
	ARM64_FMINNM    operation = 285
	ARM64_FMINNMP   operation = 286
	ARM64_FMINNMV   operation = 287
	ARM64_FMINP     operation = 288
	ARM64_FMINV     operation = 289
	ARM64_FMLA      operation = 290
	ARM64_FMLAL     operation = 291
	ARM64_FMLAL2    operation = 292
	ARM64_FMLALB    operation = 293
	ARM64_FMLALT    operation = 294
	ARM64_FMLS      operation = 295
	ARM64_FMLSL     operation = 296
	ARM64_FMLSL2    operation = 297
	ARM64_FMLSLB    operation = 298
	ARM64_FMLSLT    operation = 299
	ARM64_FMMLA     operation = 300
	ARM64_FMOPA     operation = 301
	ARM64_FMOPS     operation = 302
	ARM64_FMOV      operation = 303
	ARM64_FMSB      operation = 304
	ARM64_FMSUB     operation = 305
	ARM64_FMUL      operation = 306
	ARM64_FMULX     operation = 307
	ARM64_FNEG      operation = 308
	ARM64_FNMAD     operation = 309
	ARM64_FNMADD    operation = 310
	ARM64_FNMLA     operation = 311
	ARM64_FNMLS     operation = 312
	ARM64_FNMSB     operation = 313
	ARM64_FNMSUB    operation = 314
	ARM64_FNMUL     operation = 315
	ARM64_FRECPE    operation = 316
	ARM64_FRECPS    operation = 317
	ARM64_FRECPX    operation = 318
	ARM64_FRINT32X  operation = 319
	ARM64_FRINT32Z  operation = 320
	ARM64_FRINT64X  operation = 321
	ARM64_FRINT64Z  operation = 322
	ARM64_FRINTA    operation = 323
	ARM64_FRINTI    operation = 324
	ARM64_FRINTM    operation = 325
	ARM64_FRINTN    operation = 326
	ARM64_FRINTP    operation = 327
	ARM64_FRINTX    operation = 328
	ARM64_FRINTZ    operation = 329
	ARM64_FRSQRTE   operation = 330
	ARM64_FRSQRTS   operation = 331
	ARM64_FSCALE    operation = 332
	ARM64_FSQRT     operation = 333
	ARM64_FSUB      operation = 334
	ARM64_FSUBR     operation = 335
	ARM64_FTMAD     operation = 336
	ARM64_FTSMUL    operation = 337
	ARM64_FTSSEL    operation = 338
	ARM64_GMI       operation = 339
	ARM64_HINT      operation = 340
	ARM64_HISTCNT   operation = 341
	ARM64_HISTSEG   operation = 342
	ARM64_HLT       operation = 343
	ARM64_HVC       operation = 344
	ARM64_IC        operation = 345
	ARM64_INCB      operation = 346
	ARM64_INCD      operation = 347
	ARM64_INCH      operation = 348
	ARM64_INCP      operation = 349
	ARM64_INCW      operation = 350
	ARM64_INDEX     operation = 351
	ARM64_INS       operation = 352
	ARM64_INSR      operation = 353
	ARM64_IRG       operation = 354
	ARM64_ISB       operation = 355
	ARM64_LASTA     operation = 356
	ARM64_LASTB     operation = 357
	ARM64_LD1       operation = 358
	ARM64_LD1B      operation = 359
	ARM64_LD1D      operation = 360
	ARM64_LD1H      operation = 361
	ARM64_LD1Q      operation = 362
	ARM64_LD1R      operation = 363
	ARM64_LD1RB     operation = 364
	ARM64_LD1RD     operation = 365
	ARM64_LD1RH     operation = 366
	ARM64_LD1ROB    operation = 367
	ARM64_LD1ROD    operation = 368
	ARM64_LD1ROH    operation = 369
	ARM64_LD1ROW    operation = 370
	ARM64_LD1RQB    operation = 371
	ARM64_LD1RQD    operation = 372
	ARM64_LD1RQH    operation = 373
	ARM64_LD1RQW    operation = 374
	ARM64_LD1RSB    operation = 375
	ARM64_LD1RSH    operation = 376
	ARM64_LD1RSW    operation = 377
	ARM64_LD1RW     operation = 378
	ARM64_LD1SB     operation = 379
	ARM64_LD1SH     operation = 380
	ARM64_LD1SW     operation = 381
	ARM64_LD1W      operation = 382
	ARM64_LD2       operation = 383
	ARM64_LD2B      operation = 384
	ARM64_LD2D      operation = 385
	ARM64_LD2H      operation = 386
	ARM64_LD2R      operation = 387
	ARM64_LD2W      operation = 388
	ARM64_LD3       operation = 389
	ARM64_LD3B      operation = 390
	ARM64_LD3D      operation = 391
	ARM64_LD3H      operation = 392
	ARM64_LD3R      operation = 393
	ARM64_LD3W      operation = 394
	ARM64_LD4       operation = 395
	ARM64_LD4B      operation = 396
	ARM64_LD4D      operation = 397
	ARM64_LD4H      operation = 398
	ARM64_LD4R      operation = 399
	ARM64_LD4W      operation = 400
	ARM64_LD64B     operation = 401
	ARM64_LDADD     operation = 402
	ARM64_LDADDA    operation = 403
	ARM64_LDADDAB   operation = 404
	ARM64_LDADDAH   operation = 405
	ARM64_LDADDAL   operation = 406
	ARM64_LDADDALB  operation = 407
	ARM64_LDADDALH  operation = 408
	ARM64_LDADDB    operation = 409
	ARM64_LDADDH    operation = 410
	ARM64_LDADDL    operation = 411
	ARM64_LDADDLB   operation = 412
	ARM64_LDADDLH   operation = 413
	ARM64_LDAPR     operation = 414
	ARM64_LDAPRB    operation = 415
	ARM64_LDAPRH    operation = 416
	ARM64_LDAPUR    operation = 417
	ARM64_LDAPURB   operation = 418
	ARM64_LDAPURH   operation = 419
	ARM64_LDAPURSB  operation = 420
	ARM64_LDAPURSH  operation = 421
	ARM64_LDAPURSW  operation = 422
	ARM64_LDAR      operation = 423
	ARM64_LDARB     operation = 424
	ARM64_LDARH     operation = 425
	ARM64_LDAXP     operation = 426
	ARM64_LDAXR     operation = 427
	ARM64_LDAXRB    operation = 428
	ARM64_LDAXRH    operation = 429
	ARM64_LDCLR     operation = 430
	ARM64_LDCLRA    operation = 431
	ARM64_LDCLRAB   operation = 432
	ARM64_LDCLRAH   operation = 433
	ARM64_LDCLRAL   operation = 434
	ARM64_LDCLRALB  operation = 435
	ARM64_LDCLRALH  operation = 436
	ARM64_LDCLRB    operation = 437
	ARM64_LDCLRH    operation = 438
	ARM64_LDCLRL    operation = 439
	ARM64_LDCLRLB   operation = 440
	ARM64_LDCLRLH   operation = 441
	ARM64_LDEOR     operation = 442
	ARM64_LDEORA    operation = 443
	ARM64_LDEORAB   operation = 444
	ARM64_LDEORAH   operation = 445
	ARM64_LDEORAL   operation = 446
	ARM64_LDEORALB  operation = 447
	ARM64_LDEORALH  operation = 448
	ARM64_LDEORB    operation = 449
	ARM64_LDEORH    operation = 450
	ARM64_LDEORL    operation = 451
	ARM64_LDEORLB   operation = 452
	ARM64_LDEORLH   operation = 453
	ARM64_LDFF1B    operation = 454
	ARM64_LDFF1D    operation = 455
	ARM64_LDFF1H    operation = 456
	ARM64_LDFF1SB   operation = 457
	ARM64_LDFF1SH   operation = 458
	ARM64_LDFF1SW   operation = 459
	ARM64_LDFF1W    operation = 460
	ARM64_LDG       operation = 461
	ARM64_LDGM      operation = 462
	ARM64_LDLAR     operation = 463
	ARM64_LDLARB    operation = 464
	ARM64_LDLARH    operation = 465
	ARM64_LDNF1B    operation = 466
	ARM64_LDNF1D    operation = 467
	ARM64_LDNF1H    operation = 468
	ARM64_LDNF1SB   operation = 469
	ARM64_LDNF1SH   operation = 470
	ARM64_LDNF1SW   operation = 471
	ARM64_LDNF1W    operation = 472
	ARM64_LDNP      operation = 473
	ARM64_LDNT1B    operation = 474
	ARM64_LDNT1D    operation = 475
	ARM64_LDNT1H    operation = 476
	ARM64_LDNT1SB   operation = 477
	ARM64_LDNT1SH   operation = 478
	ARM64_LDNT1SW   operation = 479
	ARM64_LDNT1W    operation = 480
	ARM64_LDP       operation = 481
	ARM64_LDPSW     operation = 482
	ARM64_LDR       operation = 483
	ARM64_LDRAA     operation = 484
	ARM64_LDRAB     operation = 485
	ARM64_LDRB      operation = 486
	ARM64_LDRH      operation = 487
	ARM64_LDRSB     operation = 488
	ARM64_LDRSH     operation = 489
	ARM64_LDRSW     operation = 490
	ARM64_LDSET     operation = 491
	ARM64_LDSETA    operation = 492
	ARM64_LDSETAB   operation = 493
	ARM64_LDSETAH   operation = 494
	ARM64_LDSETAL   operation = 495
	ARM64_LDSETALB  operation = 496
	ARM64_LDSETALH  operation = 497
	ARM64_LDSETB    operation = 498
	ARM64_LDSETH    operation = 499
	ARM64_LDSETL    operation = 500
	ARM64_LDSETLB   operation = 501
	ARM64_LDSETLH   operation = 502
	ARM64_LDSMAX    operation = 503
	ARM64_LDSMAXA   operation = 504
	ARM64_LDSMAXAB  operation = 505
	ARM64_LDSMAXAH  operation = 506
	ARM64_LDSMAXAL  operation = 507
	ARM64_LDSMAXALB operation = 508
	ARM64_LDSMAXALH operation = 509
	ARM64_LDSMAXB   operation = 510
	ARM64_LDSMAXH   operation = 511
	ARM64_LDSMAXL   operation = 512
	ARM64_LDSMAXLB  operation = 513
	ARM64_LDSMAXLH  operation = 514
	ARM64_LDSMIN    operation = 515
	ARM64_LDSMINA   operation = 516
	ARM64_LDSMINAB  operation = 517
	ARM64_LDSMINAH  operation = 518
	ARM64_LDSMINAL  operation = 519
	ARM64_LDSMINALB operation = 520
	ARM64_LDSMINALH operation = 521
	ARM64_LDSMINB   operation = 522
	ARM64_LDSMINH   operation = 523
	ARM64_LDSMINL   operation = 524
	ARM64_LDSMINLB  operation = 525
	ARM64_LDSMINLH  operation = 526
	ARM64_LDTR      operation = 527
	ARM64_LDTRB     operation = 528
	ARM64_LDTRH     operation = 529
	ARM64_LDTRSB    operation = 530
	ARM64_LDTRSH    operation = 531
	ARM64_LDTRSW    operation = 532
	ARM64_LDUMAX    operation = 533
	ARM64_LDUMAXA   operation = 534
	ARM64_LDUMAXAB  operation = 535
	ARM64_LDUMAXAH  operation = 536
	ARM64_LDUMAXAL  operation = 537
	ARM64_LDUMAXALB operation = 538
	ARM64_LDUMAXALH operation = 539
	ARM64_LDUMAXB   operation = 540
	ARM64_LDUMAXH   operation = 541
	ARM64_LDUMAXL   operation = 542
	ARM64_LDUMAXLB  operation = 543
	ARM64_LDUMAXLH  operation = 544
	ARM64_LDUMIN    operation = 545
	ARM64_LDUMINA   operation = 546
	ARM64_LDUMINAB  operation = 547
	ARM64_LDUMINAH  operation = 548
	ARM64_LDUMINAL  operation = 549
	ARM64_LDUMINALB operation = 550
	ARM64_LDUMINALH operation = 551
	ARM64_LDUMINB   operation = 552
	ARM64_LDUMINH   operation = 553
	ARM64_LDUMINL   operation = 554
	ARM64_LDUMINLB  operation = 555
	ARM64_LDUMINLH  operation = 556
	ARM64_LDUR      operation = 557
	ARM64_LDURB     operation = 558
	ARM64_LDURH     operation = 559
	ARM64_LDURSB    operation = 560
	ARM64_LDURSH    operation = 561
	ARM64_LDURSW    operation = 562
	ARM64_LDXP      operation = 563
	ARM64_LDXR      operation = 564
	ARM64_LDXRB     operation = 565
	ARM64_LDXRH     operation = 566
	ARM64_LSL       operation = 567
	ARM64_LSLR      operation = 568
	ARM64_LSLV      operation = 569
	ARM64_LSR       operation = 570
	ARM64_LSRR      operation = 571
	ARM64_LSRV      operation = 572
	ARM64_MAD       operation = 573
	ARM64_MADD      operation = 574
	ARM64_MATCH     operation = 575
	ARM64_MLA       operation = 576
	ARM64_MLS       operation = 577
	ARM64_MNEG      operation = 578
	ARM64_MOV       operation = 579
	ARM64_MOVA      operation = 580
	ARM64_MOVI      operation = 581
	ARM64_MOVK      operation = 582
	ARM64_MOVN      operation = 583
	ARM64_MOVPRFX   operation = 584
	ARM64_MOVS      operation = 585
	ARM64_MOVZ      operation = 586
	ARM64_MRS       operation = 587
	ARM64_MSB       operation = 588
	ARM64_MSR       operation = 589
	ARM64_MSUB      operation = 590
	ARM64_MUL       operation = 591
	ARM64_MVN       operation = 592
	ARM64_MVNI      operation = 593
	ARM64_NAND      operation = 594
	ARM64_NANDS     operation = 595
	ARM64_NBSL      operation = 596
	ARM64_NEG       operation = 597
	ARM64_NEGS      operation = 598
	ARM64_NGC       operation = 599
	ARM64_NGCS      operation = 600
	ARM64_NMATCH    operation = 601
	ARM64_NOP       operation = 602
	ARM64_NOR       operation = 603
	ARM64_NORS      operation = 604
	ARM64_NOT       operation = 605
	ARM64_NOTS      operation = 606
	ARM64_ORN       operation = 607
	ARM64_ORNS      operation = 608
	ARM64_ORR       operation = 609
	ARM64_ORRS      operation = 610
	ARM64_ORV       operation = 611
	ARM64_PACDA     operation = 612
	ARM64_PACDB     operation = 613
	ARM64_PACDZA    operation = 614
	ARM64_PACDZB    operation = 615
	ARM64_PACGA     operation = 616
	ARM64_PACIA     operation = 617
	ARM64_PACIA1716 operation = 618
	ARM64_PACIASP   operation = 619
	ARM64_PACIAZ    operation = 620
	ARM64_PACIB     operation = 621
	ARM64_PACIB1716 operation = 622
	ARM64_PACIBSP   operation = 623
	ARM64_PACIBZ    operation = 624
	ARM64_PACIZA    operation = 625
	ARM64_PACIZB    operation = 626
	ARM64_PFALSE    operation = 627
	ARM64_PFIRST    operation = 628
	ARM64_PMUL      operation = 629
	ARM64_PMULL     operation = 630
	ARM64_PMULL2    operation = 631
	ARM64_PMULLB    operation = 632
	ARM64_PMULLT    operation = 633
	ARM64_PNEXT     operation = 634
	ARM64_PRFB      operation = 635
	ARM64_PRFD      operation = 636
	ARM64_PRFH      operation = 637
	ARM64_PRFM      operation = 638
	ARM64_PRFUM     operation = 639
	ARM64_PRFW      operation = 640
	ARM64_PSB       operation = 641
	ARM64_PSSBB     operation = 642
	ARM64_PTEST     operation = 643
	ARM64_PTRUE     operation = 644
	ARM64_PTRUES    operation = 645
	ARM64_PUNPKHI   operation = 646
	ARM64_PUNPKLO   operation = 647
	ARM64_RADDHN    operation = 648
	ARM64_RADDHN2   operation = 649
	ARM64_RADDHNB   operation = 650
	ARM64_RADDHNT   operation = 651
	ARM64_RAX1      operation = 652
	ARM64_RBIT      operation = 653
	ARM64_RDFFR     operation = 654
	ARM64_RDFFRS    operation = 655
	ARM64_RDVL      operation = 656
	ARM64_RET       operation = 657
	ARM64_RETAA     operation = 658
	ARM64_RETAB     operation = 659
	ARM64_REV       operation = 660
	ARM64_REV16     operation = 661
	ARM64_REV32     operation = 662
	ARM64_REV64     operation = 663
	ARM64_REVB      operation = 664
	ARM64_REVD      operation = 665
	ARM64_REVH      operation = 666
	ARM64_REVW      operation = 667
	ARM64_RMIF      operation = 668
	ARM64_ROR       operation = 669
	ARM64_RORV      operation = 670
	ARM64_RSHRN     operation = 671
	ARM64_RSHRN2    operation = 672
	ARM64_RSHRNB    operation = 673
	ARM64_RSHRNT    operation = 674
	ARM64_RSUBHN    operation = 675
	ARM64_RSUBHN2   operation = 676
	ARM64_RSUBHNB   operation = 677
	ARM64_RSUBHNT   operation = 678
	ARM64_SABA      operation = 679
	ARM64_SABAL     operation = 680
	ARM64_SABAL2    operation = 681
	ARM64_SABALB    operation = 682
	ARM64_SABALT    operation = 683
	ARM64_SABD      operation = 684
	ARM64_SABDL     operation = 685
	ARM64_SABDL2    operation = 686
	ARM64_SABDLB    operation = 687
	ARM64_SABDLT    operation = 688
	ARM64_SADALP    operation = 689
	ARM64_SADDL     operation = 690
	ARM64_SADDL2    operation = 691
	ARM64_SADDLB    operation = 692
	ARM64_SADDLBT   operation = 693
	ARM64_SADDLP    operation = 694
	ARM64_SADDLT    operation = 695
	ARM64_SADDLV    operation = 696
	ARM64_SADDV     operation = 697
	ARM64_SADDW     operation = 698
	ARM64_SADDW2    operation = 699
	ARM64_SADDWB    operation = 700
	ARM64_SADDWT    operation = 701
	ARM64_SB        operation = 702
	ARM64_SBC       operation = 703
	ARM64_SBCLB     operation = 704
	ARM64_SBCLT     operation = 705
	ARM64_SBCS      operation = 706
	ARM64_SBFIZ     operation = 707
	ARM64_SBFM      operation = 708
	ARM64_SBFX      operation = 709
	ARM64_SCLAMP    operation = 710
	ARM64_SCVTF     operation = 711
	ARM64_SDIV      operation = 712
	ARM64_SDIVR     operation = 713
	ARM64_SDOT      operation = 714
	ARM64_SEL       operation = 715
	ARM64_SETF16    operation = 716
	ARM64_SETF8     operation = 717
	ARM64_SETFFR    operation = 718
	ARM64_SEV       operation = 719
	ARM64_SEVL      operation = 720
	ARM64_SHA1C     operation = 721
	ARM64_SHA1H     operation = 722
	ARM64_SHA1M     operation = 723
	ARM64_SHA1P     operation = 724
	ARM64_SHA1SU0   operation = 725
	ARM64_SHA1SU1   operation = 726
	ARM64_SHA256H   operation = 727
	ARM64_SHA256H2  operation = 728
	ARM64_SHA256SU0 operation = 729
	ARM64_SHA256SU1 operation = 730
	ARM64_SHA512H   operation = 731
	ARM64_SHA512H2  operation = 732
	ARM64_SHA512SU0 operation = 733
	ARM64_SHA512SU1 operation = 734
	ARM64_SHADD     operation = 735
	ARM64_SHL       operation = 736
	ARM64_SHLL      operation = 737
	ARM64_SHLL2     operation = 738
	ARM64_SHRN      operation = 739
	ARM64_SHRN2     operation = 740
	ARM64_SHRNB     operation = 741
	ARM64_SHRNT     operation = 742
	ARM64_SHSUB     operation = 743
	ARM64_SHSUBR    operation = 744
	ARM64_SLI       operation = 745
	ARM64_SM3PARTW1 operation = 746
	ARM64_SM3PARTW2 operation = 747
	ARM64_SM3SS1    operation = 748
	ARM64_SM3TT1A   operation = 749
	ARM64_SM3TT1B   operation = 750
	ARM64_SM3TT2A   operation = 751
	ARM64_SM3TT2B   operation = 752
	ARM64_SM4E      operation = 753
	ARM64_SM4EKEY   operation = 754
	ARM64_SMADDL    operation = 755
	ARM64_SMAX      operation = 756
	ARM64_SMAXP     operation = 757
	ARM64_SMAXV     operation = 758
	ARM64_SMC       operation = 759
	ARM64_SMIN      operation = 760
	ARM64_SMINP     operation = 761
	ARM64_SMINV     operation = 762
	ARM64_SMLAL     operation = 763
	ARM64_SMLAL2    operation = 764
	ARM64_SMLALB    operation = 765
	ARM64_SMLALT    operation = 766
	ARM64_SMLSL     operation = 767
	ARM64_SMLSL2    operation = 768
	ARM64_SMLSLB    operation = 769
	ARM64_SMLSLT    operation = 770
	ARM64_SMMLA     operation = 771
	ARM64_SMNEGL    operation = 772
	ARM64_SMOPA     operation = 773
	ARM64_SMOPS     operation = 774
	ARM64_SMOV      operation = 775
	ARM64_SMSTART   operation = 776
	ARM64_SMSTOP    operation = 777
	ARM64_SMSUBL    operation = 778
	ARM64_SMULH     operation = 779
	ARM64_SMULL     operation = 780
	ARM64_SMULL2    operation = 781
	ARM64_SMULLB    operation = 782
	ARM64_SMULLT    operation = 783
	ARM64_SPLICE    operation = 784
	ARM64_SQABS     operation = 785
	ARM64_SQADD     operation = 786
	ARM64_SQCADD    operation = 787
	ARM64_SQDECB    operation = 788
	ARM64_SQDECD    operation = 789
	ARM64_SQDECH    operation = 790
	ARM64_SQDECP    operation = 791
	ARM64_SQDECW    operation = 792
	ARM64_SQDMLAL   operation = 793
	ARM64_SQDMLAL2  operation = 794
	ARM64_SQDMLALB  operation = 795
	ARM64_SQDMLALBT operation = 796
	ARM64_SQDMLALT  operation = 797
	ARM64_SQDMLSL   operation = 798
	ARM64_SQDMLSL2  operation = 799
	ARM64_SQDMLSLB  operation = 800
	ARM64_SQDMLSLBT operation = 801
	ARM64_SQDMLSLT  operation = 802
	ARM64_SQDMULH   operation = 803
	ARM64_SQDMULL   operation = 804
	ARM64_SQDMULL2  operation = 805
	ARM64_SQDMULLB  operation = 806
	ARM64_SQDMULLT  operation = 807
	ARM64_SQINCB    operation = 808
	ARM64_SQINCD    operation = 809
	ARM64_SQINCH    operation = 810
	ARM64_SQINCP    operation = 811
	ARM64_SQINCW    operation = 812
	ARM64_SQNEG     operation = 813
	ARM64_SQRDCMLAH operation = 814
	ARM64_SQRDMLAH  operation = 815
	ARM64_SQRDMLSH  operation = 816
	ARM64_SQRDMULH  operation = 817
	ARM64_SQRSHL    operation = 818
	ARM64_SQRSHLR   operation = 819
	ARM64_SQRSHRN   operation = 820
	ARM64_SQRSHRN2  operation = 821
	ARM64_SQRSHRNB  operation = 822
	ARM64_SQRSHRNT  operation = 823
	ARM64_SQRSHRUN  operation = 824
	ARM64_SQRSHRUN2 operation = 825
	ARM64_SQRSHRUNB operation = 826
	ARM64_SQRSHRUNT operation = 827
	ARM64_SQSHL     operation = 828
	ARM64_SQSHLR    operation = 829
	ARM64_SQSHLU    operation = 830
	ARM64_SQSHRN    operation = 831
	ARM64_SQSHRN2   operation = 832
	ARM64_SQSHRNB   operation = 833
	ARM64_SQSHRNT   operation = 834
	ARM64_SQSHRUN   operation = 835
	ARM64_SQSHRUN2  operation = 836
	ARM64_SQSHRUNB  operation = 837
	ARM64_SQSHRUNT  operation = 838
	ARM64_SQSUB     operation = 839
	ARM64_SQSUBR    operation = 840
	ARM64_SQXTN     operation = 841
	ARM64_SQXTN2    operation = 842
	ARM64_SQXTNB    operation = 843
	ARM64_SQXTNT    operation = 844
	ARM64_SQXTUN    operation = 845
	ARM64_SQXTUN2   operation = 846
	ARM64_SQXTUNB   operation = 847
	ARM64_SQXTUNT   operation = 848
	ARM64_SRHADD    operation = 849
	ARM64_SRI       operation = 850
	ARM64_SRSHL     operation = 851
	ARM64_SRSHLR    operation = 852
	ARM64_SRSHR     operation = 853
	ARM64_SRSRA     operation = 854
	ARM64_SSBB      operation = 855
	ARM64_SSHL      operation = 856
	ARM64_SSHLL     operation = 857
	ARM64_SSHLL2    operation = 858
	ARM64_SSHLLB    operation = 859
	ARM64_SSHLLT    operation = 860
	ARM64_SSHR      operation = 861
	ARM64_SSRA      operation = 862
	ARM64_SSUBL     operation = 863
	ARM64_SSUBL2    operation = 864
	ARM64_SSUBLB    operation = 865
	ARM64_SSUBLBT   operation = 866
	ARM64_SSUBLT    operation = 867
	ARM64_SSUBLTB   operation = 868
	ARM64_SSUBW     operation = 869
	ARM64_SSUBW2    operation = 870
	ARM64_SSUBWB    operation = 871
	ARM64_SSUBWT    operation = 872
	ARM64_ST1       operation = 873
	ARM64_ST1B      operation = 874
	ARM64_ST1D      operation = 875
	ARM64_ST1H      operation = 876
	ARM64_ST1Q      operation = 877
	ARM64_ST1W      operation = 878
	ARM64_ST2       operation = 879
	ARM64_ST2B      operation = 880
	ARM64_ST2D      operation = 881
	ARM64_ST2G      operation = 882
	ARM64_ST2H      operation = 883
	ARM64_ST2W      operation = 884
	ARM64_ST3       operation = 885
	ARM64_ST3B      operation = 886
	ARM64_ST3D      operation = 887
	ARM64_ST3H      operation = 888
	ARM64_ST3W      operation = 889
	ARM64_ST4       operation = 890
	ARM64_ST4B      operation = 891
	ARM64_ST4D      operation = 892
	ARM64_ST4H      operation = 893
	ARM64_ST4W      operation = 894
	ARM64_ST64B     operation = 895
	ARM64_ST64BV    operation = 896
	ARM64_ST64BV0   operation = 897
	ARM64_STADD     operation = 898
	ARM64_STADDB    operation = 899
	ARM64_STADDH    operation = 900
	ARM64_STADDL    operation = 901
	ARM64_STADDLB   operation = 902
	ARM64_STADDLH   operation = 903
	ARM64_STCLR     operation = 904
	ARM64_STCLRB    operation = 905
	ARM64_STCLRH    operation = 906
	ARM64_STCLRL    operation = 907
	ARM64_STCLRLB   operation = 908
	ARM64_STCLRLH   operation = 909
	ARM64_STEOR     operation = 910
	ARM64_STEORB    operation = 911
	ARM64_STEORH    operation = 912
	ARM64_STEORL    operation = 913
	ARM64_STEORLB   operation = 914
	ARM64_STEORLH   operation = 915
	ARM64_STG       operation = 916
	ARM64_STGM      operation = 917
	ARM64_STGP      operation = 918
	ARM64_STLLR     operation = 919
	ARM64_STLLRB    operation = 920
	ARM64_STLLRH    operation = 921
	ARM64_STLR      operation = 922
	ARM64_STLRB     operation = 923
	ARM64_STLRH     operation = 924
	ARM64_STLUR     operation = 925
	ARM64_STLURB    operation = 926
	ARM64_STLURH    operation = 927
	ARM64_STLXP     operation = 928
	ARM64_STLXR     operation = 929
	ARM64_STLXRB    operation = 930
	ARM64_STLXRH    operation = 931
	ARM64_STNP      operation = 932
	ARM64_STNT1B    operation = 933
	ARM64_STNT1D    operation = 934
	ARM64_STNT1H    operation = 935
	ARM64_STNT1W    operation = 936
	ARM64_STP       operation = 937
	ARM64_STR       operation = 938
	ARM64_STRB      operation = 939
	ARM64_STRH      operation = 940
	ARM64_STSET     operation = 941
	ARM64_STSETB    operation = 942
	ARM64_STSETH    operation = 943
	ARM64_STSETL    operation = 944
	ARM64_STSETLB   operation = 945
	ARM64_STSETLH   operation = 946
	ARM64_STSMAX    operation = 947
	ARM64_STSMAXB   operation = 948
	ARM64_STSMAXH   operation = 949
	ARM64_STSMAXL   operation = 950
	ARM64_STSMAXLB  operation = 951
	ARM64_STSMAXLH  operation = 952
	ARM64_STSMIN    operation = 953
	ARM64_STSMINB   operation = 954
	ARM64_STSMINH   operation = 955
	ARM64_STSMINL   operation = 956
	ARM64_STSMINLB  operation = 957
	ARM64_STSMINLH  operation = 958
	ARM64_STTR      operation = 959
	ARM64_STTRB     operation = 960
	ARM64_STTRH     operation = 961
	ARM64_STUMAX    operation = 962
	ARM64_STUMAXB   operation = 963
	ARM64_STUMAXH   operation = 964
	ARM64_STUMAXL   operation = 965
	ARM64_STUMAXLB  operation = 966
	ARM64_STUMAXLH  operation = 967
	ARM64_STUMIN    operation = 968
	ARM64_STUMINB   operation = 969
	ARM64_STUMINH   operation = 970
	ARM64_STUMINL   operation = 971
	ARM64_STUMINLB  operation = 972
	ARM64_STUMINLH  operation = 973
	ARM64_STUR      operation = 974
	ARM64_STURB     operation = 975
	ARM64_STURH     operation = 976
	ARM64_STXP      operation = 977
	ARM64_STXR      operation = 978
	ARM64_STXRB     operation = 979
	ARM64_STXRH     operation = 980
	ARM64_STZ2G     operation = 981
	ARM64_STZG      operation = 982
	ARM64_STZGM     operation = 983
	ARM64_SUB       operation = 984
	ARM64_SUBG      operation = 985
	ARM64_SUBHN     operation = 986
	ARM64_SUBHN2    operation = 987
	ARM64_SUBHNB    operation = 988
	ARM64_SUBHNT    operation = 989
	ARM64_SUBP      operation = 990
	ARM64_SUBPS     operation = 991
	ARM64_SUBR      operation = 992
	ARM64_SUBS      operation = 993
	ARM64_SUDOT     operation = 994
	ARM64_SUMOPA    operation = 995
	ARM64_SUMOPS    operation = 996
	ARM64_SUNPKHI   operation = 997
	ARM64_SUNPKLO   operation = 998
	ARM64_SUQADD    operation = 999
	ARM64_SVC       operation = 1000
	ARM64_SWP       operation = 1001
	ARM64_SWPA      operation = 1002
	ARM64_SWPAB     operation = 1003
	ARM64_SWPAH     operation = 1004
	ARM64_SWPAL     operation = 1005
	ARM64_SWPALB    operation = 1006
	ARM64_SWPALH    operation = 1007
	ARM64_SWPB      operation = 1008
	ARM64_SWPH      operation = 1009
	ARM64_SWPL      operation = 1010
	ARM64_SWPLB     operation = 1011
	ARM64_SWPLH     operation = 1012
	ARM64_SXTB      operation = 1013
	ARM64_SXTH      operation = 1014
	ARM64_SXTL      operation = 1015
	ARM64_SXTL2     operation = 1016
	ARM64_SXTW      operation = 1017
	ARM64_SYS       operation = 1018
	ARM64_SYSL      operation = 1019
	ARM64_TBL       operation = 1020
	ARM64_TBNZ      operation = 1021
	ARM64_TBX       operation = 1022
	ARM64_TBZ       operation = 1023
	ARM64_TCANCEL   operation = 1024
	ARM64_TCOMMIT   operation = 1025
	ARM64_TLBI      operation = 1026
	ARM64_TRN1      operation = 1027
	ARM64_TRN2      operation = 1028
	ARM64_TSB       operation = 1029
	ARM64_TST       operation = 1030
	ARM64_TSTART    operation = 1031
	ARM64_TTEST     operation = 1032
	ARM64_UABA      operation = 1033
	ARM64_UABAL     operation = 1034
	ARM64_UABAL2    operation = 1035
	ARM64_UABALB    operation = 1036
	ARM64_UABALT    operation = 1037
	ARM64_UABD      operation = 1038
	ARM64_UABDL     operation = 1039
	ARM64_UABDL2    operation = 1040
	ARM64_UABDLB    operation = 1041
	ARM64_UABDLT    operation = 1042
	ARM64_UADALP    operation = 1043
	ARM64_UADDL     operation = 1044
	ARM64_UADDL2    operation = 1045
	ARM64_UADDLB    operation = 1046
	ARM64_UADDLP    operation = 1047
	ARM64_UADDLT    operation = 1048
	ARM64_UADDLV    operation = 1049
	ARM64_UADDV     operation = 1050
	ARM64_UADDW     operation = 1051
	ARM64_UADDW2    operation = 1052
	ARM64_UADDWB    operation = 1053
	ARM64_UADDWT    operation = 1054
	ARM64_UBFIZ     operation = 1055
	ARM64_UBFM      operation = 1056
	ARM64_UBFX      operation = 1057
	ARM64_UCLAMP    operation = 1058
	ARM64_UCVTF     operation = 1059
	ARM64_UDF       operation = 1060
	ARM64_UDIV      operation = 1061
	ARM64_UDIVR     operation = 1062
	ARM64_UDOT      operation = 1063
	ARM64_UHADD     operation = 1064
	ARM64_UHSUB     operation = 1065
	ARM64_UHSUBR    operation = 1066
	ARM64_UMADDL    operation = 1067
	ARM64_UMAX      operation = 1068
	ARM64_UMAXP     operation = 1069
	ARM64_UMAXV     operation = 1070
	ARM64_UMIN      operation = 1071
	ARM64_UMINP     operation = 1072
	ARM64_UMINV     operation = 1073
	ARM64_UMLAL     operation = 1074
	ARM64_UMLAL2    operation = 1075
	ARM64_UMLALB    operation = 1076
	ARM64_UMLALT    operation = 1077
	ARM64_UMLSL     operation = 1078
	ARM64_UMLSL2    operation = 1079
	ARM64_UMLSLB    operation = 1080
	ARM64_UMLSLT    operation = 1081
	ARM64_UMMLA     operation = 1082
	ARM64_UMNEGL    operation = 1083
	ARM64_UMOPA     operation = 1084
	ARM64_UMOPS     operation = 1085
	ARM64_UMOV      operation = 1086
	ARM64_UMSUBL    operation = 1087
	ARM64_UMULH     operation = 1088
	ARM64_UMULL     operation = 1089
	ARM64_UMULL2    operation = 1090
	ARM64_UMULLB    operation = 1091
	ARM64_UMULLT    operation = 1092
	ARM64_UQADD     operation = 1093
	ARM64_UQDECB    operation = 1094
	ARM64_UQDECD    operation = 1095
	ARM64_UQDECH    operation = 1096
	ARM64_UQDECP    operation = 1097
	ARM64_UQDECW    operation = 1098
	ARM64_UQINCB    operation = 1099
	ARM64_UQINCD    operation = 1100
	ARM64_UQINCH    operation = 1101
	ARM64_UQINCP    operation = 1102
	ARM64_UQINCW    operation = 1103
	ARM64_UQRSHL    operation = 1104
	ARM64_UQRSHLR   operation = 1105
	ARM64_UQRSHRN   operation = 1106
	ARM64_UQRSHRN2  operation = 1107
	ARM64_UQRSHRNB  operation = 1108
	ARM64_UQRSHRNT  operation = 1109
	ARM64_UQSHL     operation = 1110
	ARM64_UQSHLR    operation = 1111
	ARM64_UQSHRN    operation = 1112
	ARM64_UQSHRN2   operation = 1113
	ARM64_UQSHRNB   operation = 1114
	ARM64_UQSHRNT   operation = 1115
	ARM64_UQSUB     operation = 1116
	ARM64_UQSUBR    operation = 1117
	ARM64_UQXTN     operation = 1118
	ARM64_UQXTN2    operation = 1119
	ARM64_UQXTNB    operation = 1120
	ARM64_UQXTNT    operation = 1121
	ARM64_URECPE    operation = 1122
	ARM64_URHADD    operation = 1123
	ARM64_URSHL     operation = 1124
	ARM64_URSHLR    operation = 1125
	ARM64_URSHR     operation = 1126
	ARM64_URSQRTE   operation = 1127
	ARM64_URSRA     operation = 1128
	ARM64_USDOT     operation = 1129
	ARM64_USHL      operation = 1130
	ARM64_USHLL     operation = 1131
	ARM64_USHLL2    operation = 1132
	ARM64_USHLLB    operation = 1133
	ARM64_USHLLT    operation = 1134
	ARM64_USHR      operation = 1135
	ARM64_USMMLA    operation = 1136
	ARM64_USMOPA    operation = 1137
	ARM64_USMOPS    operation = 1138
	ARM64_USQADD    operation = 1139
	ARM64_USRA      operation = 1140
	ARM64_USUBL     operation = 1141
	ARM64_USUBL2    operation = 1142
	ARM64_USUBLB    operation = 1143
	ARM64_USUBLT    operation = 1144
	ARM64_USUBW     operation = 1145
	ARM64_USUBW2    operation = 1146
	ARM64_USUBWB    operation = 1147
	ARM64_USUBWT    operation = 1148
	ARM64_UUNPKHI   operation = 1149
	ARM64_UUNPKLO   operation = 1150
	ARM64_UXTB      operation = 1151
	ARM64_UXTH      operation = 1152
	ARM64_UXTL      operation = 1153
	ARM64_UXTL2     operation = 1154
	ARM64_UXTW      operation = 1155
	ARM64_UZP1      operation = 1156
	ARM64_UZP2      operation = 1157
	ARM64_WFE       operation = 1158
	ARM64_WFET      operation = 1159
	ARM64_WFI       operation = 1160
	ARM64_WFIT      operation = 1161
	ARM64_WHILEGE   operation = 1162
	ARM64_WHILEGT   operation = 1163
	ARM64_WHILEHI   operation = 1164
	ARM64_WHILEHS   operation = 1165
	ARM64_WHILELE   operation = 1166
	ARM64_WHILELO   operation = 1167
	ARM64_WHILELS   operation = 1168
	ARM64_WHILELT   operation = 1169
	ARM64_WHILERW   operation = 1170
	ARM64_WHILEWR   operation = 1171
	ARM64_WRFFR     operation = 1172
	ARM64_XAFLAG    operation = 1173
	ARM64_XAR       operation = 1174
	ARM64_XPACD     operation = 1175
	ARM64_XPACI     operation = 1176
	ARM64_XPACLRI   operation = 1177
	ARM64_XTN       operation = 1178
	ARM64_XTN2      operation = 1179
	ARM64_YIELD     operation = 1180
	ARM64_ZERO      operation = 1181
	ARM64_ZIP1      operation = 1182
	ARM64_ZIP2      operation = 1183
)

func (o operation) String() string {
	switch o {
	case ARM64_ABS:
		return "abs"
	case ARM64_ADC:
		return "adc"
	case ARM64_ADCLB:
		return "adclb"
	case ARM64_ADCLT:
		return "adclt"
	case ARM64_ADCS:
		return "adcs"
	case ARM64_ADD:
		return "add"
	case ARM64_ADDG:
		return "addg"
	case ARM64_ADDHA:
		return "addha"
	case ARM64_ADDHN:
		return "addhn"
	case ARM64_ADDHN2:
		return "addhn2"
	case ARM64_ADDHNB:
		return "addhnb"
	case ARM64_ADDHNT:
		return "addhnt"
	case ARM64_ADDP:
		return "addp"
	case ARM64_ADDPL:
		return "addpl"
	case ARM64_ADDS:
		return "adds"
	case ARM64_ADDV:
		return "addv"
	case ARM64_ADDVA:
		return "addva"
	case ARM64_ADDVL:
		return "addvl"
	case ARM64_ADR:
		return "adr"
	case ARM64_ADRP:
		return "adrp"
	case ARM64_AESD:
		return "aesd"
	case ARM64_AESE:
		return "aese"
	case ARM64_AESIMC:
		return "aesimc"
	case ARM64_AESMC:
		return "aesmc"
	case ARM64_AND:
		return "and"
	case ARM64_ANDS:
		return "ands"
	case ARM64_ANDV:
		return "andv"
	case ARM64_ASR:
		return "asr"
	case ARM64_ASRD:
		return "asrd"
	case ARM64_ASRR:
		return "asrr"
	case ARM64_ASRV:
		return "asrv"
	case ARM64_AT:
		return "at"
	case ARM64_AUTDA:
		return "autda"
	case ARM64_AUTDB:
		return "autdb"
	case ARM64_AUTDZA:
		return "autdza"
	case ARM64_AUTDZB:
		return "autdzb"
	case ARM64_AUTIA:
		return "autia"
	case ARM64_AUTIA1716:
		return "autia1716"
	case ARM64_AUTIASP:
		return "autiasp"
	case ARM64_AUTIAZ:
		return "autiaz"
	case ARM64_AUTIB:
		return "autib"
	case ARM64_AUTIB1716:
		return "autib1716"
	case ARM64_AUTIBSP:
		return "autibsp"
	case ARM64_AUTIBZ:
		return "autibz"
	case ARM64_AUTIZA:
		return "autiza"
	case ARM64_AUTIZB:
		return "autizb"
	case ARM64_AXFLAG:
		return "axflag"
	case ARM64_B:
		return "b"
	case ARM64_BCAX:
		return "bcax"
	case ARM64_BDEP:
		return "bdep"
	case ARM64_BEXT:
		return "bext"
	case ARM64_BFC:
		return "bfc"
	case ARM64_BFCVT:
		return "bfcvt"
	case ARM64_BFCVTN:
		return "bfcvtn"
	case ARM64_BFCVTN2:
		return "bfcvtn2"
	case ARM64_BFCVTNT:
		return "bfcvtnt"
	case ARM64_BFDOT:
		return "bfdot"
	case ARM64_BFI:
		return "bfi"
	case ARM64_BFM:
		return "bfm"
	case ARM64_BFMLAL:
		return "bfmlal"
	case ARM64_BFMLALB:
		return "bfmlalb"
	case ARM64_BFMLALT:
		return "bfmlalt"
	case ARM64_BFMMLA:
		return "bfmmla"
	case ARM64_BFMOPA:
		return "bfmopa"
	case ARM64_BFMOPS:
		return "bfmops"
	case ARM64_BFXIL:
		return "bfxil"
	case ARM64_BGRP:
		return "bgrp"
	case ARM64_BIC:
		return "bic"
	case ARM64_BICS:
		return "bics"
	case ARM64_BIF:
		return "bif"
	case ARM64_BIT:
		return "bit"
	case ARM64_BL:
		return "bl"
	case ARM64_BLR:
		return "blr"
	case ARM64_BLRAA:
		return "blraa"
	case ARM64_BLRAAZ:
		return "blraaz"
	case ARM64_BLRAB:
		return "blrab"
	case ARM64_BLRABZ:
		return "blrabz"
	case ARM64_BR:
		return "br"
	case ARM64_BRAA:
		return "braa"
	case ARM64_BRAAZ:
		return "braaz"
	case ARM64_BRAB:
		return "brab"
	case ARM64_BRABZ:
		return "brabz"
	case ARM64_BRK:
		return "brk"
	case ARM64_BRKA:
		return "brka"
	case ARM64_BRKAS:
		return "brkas"
	case ARM64_BRKB:
		return "brkb"
	case ARM64_BRKBS:
		return "brkbs"
	case ARM64_BRKN:
		return "brkn"
	case ARM64_BRKNS:
		return "brkns"
	case ARM64_BRKPA:
		return "brkpa"
	case ARM64_BRKPAS:
		return "brkpas"
	case ARM64_BRKPB:
		return "brkpb"
	case ARM64_BRKPBS:
		return "brkpbs"
	case ARM64_BSL:
		return "bsl"
	case ARM64_BSL1N:
		return "bsl1n"
	case ARM64_BSL2N:
		return "bsl2n"
	case ARM64_BTI:
		return "bti"
	case ARM64_B_AL:
		return "b.al"
	case ARM64_B_CC:
		return "b.lo"
	case ARM64_B_CS:
		return "b.hs"
	case ARM64_B_EQ:
		return "b.eq"
	case ARM64_B_GE:
		return "b.ge"
	case ARM64_B_GT:
		return "b.gt"
	case ARM64_B_HI:
		return "b.hi"
	case ARM64_B_LE:
		return "b.le"
	case ARM64_B_LS:
		return "b.ls"
	case ARM64_B_LT:
		return "b.lt"
	case ARM64_B_MI:
		return "b.mi"
	case ARM64_B_NE:
		return "b.ne"
	case ARM64_B_NV:
		return "b.nv"
	case ARM64_B_PL:
		return "b.pl"
	case ARM64_B_VC:
		return "b.vc"
	case ARM64_B_VS:
		return "b.vs"
	case ARM64_CADD:
		return "cadd"
	case ARM64_CAS:
		return "cas"
	case ARM64_CASA:
		return "casa"
	case ARM64_CASAB:
		return "casab"
	case ARM64_CASAH:
		return "casah"
	case ARM64_CASAL:
		return "casal"
	case ARM64_CASALB:
		return "casalb"
	case ARM64_CASALH:
		return "casalh"
	case ARM64_CASB:
		return "casb"
	case ARM64_CASH:
		return "cash"
	case ARM64_CASL:
		return "casl"
	case ARM64_CASLB:
		return "caslb"
	case ARM64_CASLH:
		return "caslh"
	case ARM64_CASP:
		return "casp"
	case ARM64_CASPA:
		return "caspa"
	case ARM64_CASPAL:
		return "caspal"
	case ARM64_CASPL:
		return "caspl"
	case ARM64_CBNZ:
		return "cbnz"
	case ARM64_CBZ:
		return "cbz"
	case ARM64_CCMN:
		return "ccmn"
	case ARM64_CCMP:
		return "ccmp"
	case ARM64_CDOT:
		return "cdot"
	case ARM64_CFINV:
		return "cfinv"
	case ARM64_CFP:
		return "cfp"
	case ARM64_CINC:
		return "cinc"
	case ARM64_CINV:
		return "cinv"
	case ARM64_CLASTA:
		return "clasta"
	case ARM64_CLASTB:
		return "clastb"
	case ARM64_CLREX:
		return "clrex"
	case ARM64_CLS:
		return "cls"
	case ARM64_CLZ:
		return "clz"
	case ARM64_CMEQ:
		return "cmeq"
	case ARM64_CMGE:
		return "cmge"
	case ARM64_CMGT:
		return "cmgt"
	case ARM64_CMHI:
		return "cmhi"
	case ARM64_CMHS:
		return "cmhs"
	case ARM64_CMLA:
		return "cmla"
	case ARM64_CMLE:
		return "cmle"
	case ARM64_CMLT:
		return "cmlt"
	case ARM64_CMN:
		return "cmn"
	case ARM64_CMP:
		return "cmp"
	case ARM64_CMPEQ:
		return "cmpeq"
	case ARM64_CMPGE:
		return "cmpge"
	case ARM64_CMPGT:
		return "cmpgt"
	case ARM64_CMPHI:
		return "cmphi"
	case ARM64_CMPHS:
		return "cmphs"
	case ARM64_CMPLE:
		return "cmple"
	case ARM64_CMPLO:
		return "cmplo"
	case ARM64_CMPLS:
		return "cmpls"
	case ARM64_CMPLT:
		return "cmplt"
	case ARM64_CMPNE:
		return "cmpne"
	case ARM64_CMPP:
		return "cmpp"
	case ARM64_CMTST:
		return "cmtst"
	case ARM64_CNEG:
		return "cneg"
	case ARM64_CNOT:
		return "cnot"
	case ARM64_CNT:
		return "cnt"
	case ARM64_CNTB:
		return "cntb"
	case ARM64_CNTD:
		return "cntd"
	case ARM64_CNTH:
		return "cnth"
	case ARM64_CNTP:
		return "cntp"
	case ARM64_CNTW:
		return "cntw"
	case ARM64_COMPACT:
		return "compact"
	case ARM64_CPP:
		return "cpp"
	case ARM64_CPY:
		return "cpy"
	case ARM64_CRC32B:
		return "crc32b"
	case ARM64_CRC32CB:
		return "crc32cb"
	case ARM64_CRC32CH:
		return "crc32ch"
	case ARM64_CRC32CW:
		return "crc32cw"
	case ARM64_CRC32CX:
		return "crc32cx"
	case ARM64_CRC32H:
		return "crc32h"
	case ARM64_CRC32W:
		return "crc32w"
	case ARM64_CRC32X:
		return "crc32x"
	case ARM64_CSDB:
		return "csdb"
	case ARM64_CSEL:
		return "csel"
	case ARM64_CSET:
		return "cset"
	case ARM64_CSETM:
		return "csetm"
	case ARM64_CSINC:
		return "csinc"
	case ARM64_CSINV:
		return "csinv"
	case ARM64_CSNEG:
		return "csneg"
	case ARM64_CTERMEQ:
		return "ctermeq"
	case ARM64_CTERMNE:
		return "ctermne"
	case ARM64_DC:
		return "dc"
	case ARM64_DCPS1:
		return "dcps1"
	case ARM64_DCPS2:
		return "dcps2"
	case ARM64_DCPS3:
		return "dcps3"
	case ARM64_DECB:
		return "decb"
	case ARM64_DECD:
		return "decd"
	case ARM64_DECH:
		return "dech"
	case ARM64_DECP:
		return "decp"
	case ARM64_DECW:
		return "decw"
	case ARM64_DGH:
		return "dgh"
	case ARM64_DMB:
		return "dmb"
	case ARM64_DRPS:
		return "drps"
	case ARM64_DSB:
		return "dsb"
	case ARM64_DUP:
		return "dup"
	case ARM64_DUPM:
		return "dupm"
	case ARM64_DVP:
		return "dvp"
	case ARM64_EON:
		return "eon"
	case ARM64_EOR:
		return "eor"
	case ARM64_EOR3:
		return "eor3"
	case ARM64_EORBT:
		return "eorbt"
	case ARM64_EORS:
		return "eors"
	case ARM64_EORTB:
		return "eortb"
	case ARM64_EORV:
		return "eorv"
	case ARM64_ERET:
		return "eret"
	case ARM64_ERETAA:
		return "eretaa"
	case ARM64_ERETAB:
		return "eretab"
	case ARM64_ESB:
		return "esb"
	case ARM64_EXT:
		return "ext"
	case ARM64_EXTR:
		return "extr"
	case ARM64_FABD:
		return "fabd"
	case ARM64_FABS:
		return "fabs"
	case ARM64_FACGE:
		return "facge"
	case ARM64_FACGT:
		return "facgt"
	case ARM64_FACLE:
		return "facle"
	case ARM64_FACLT:
		return "faclt"
	case ARM64_FADD:
		return "fadd"
	case ARM64_FADDA:
		return "fadda"
	case ARM64_FADDP:
		return "faddp"
	case ARM64_FADDV:
		return "faddv"
	case ARM64_FCADD:
		return "fcadd"
	case ARM64_FCCMP:
		return "fccmp"
	case ARM64_FCCMPE:
		return "fccmpe"
	case ARM64_FCMEQ:
		return "fcmeq"
	case ARM64_FCMGE:
		return "fcmge"
	case ARM64_FCMGT:
		return "fcmgt"
	case ARM64_FCMLA:
		return "fcmla"
	case ARM64_FCMLE:
		return "fcmle"
	case ARM64_FCMLT:
		return "fcmlt"
	case ARM64_FCMNE:
		return "fcmne"
	case ARM64_FCMP:
		return "fcmp"
	case ARM64_FCMPE:
		return "fcmpe"
	case ARM64_FCMUO:
		return "fcmuo"
	case ARM64_FCPY:
		return "fcpy"
	case ARM64_FCSEL:
		return "fcsel"
	case ARM64_FCVT:
		return "fcvt"
	case ARM64_FCVTAS:
		return "fcvtas"
	case ARM64_FCVTAU:
		return "fcvtau"
	case ARM64_FCVTL:
		return "fcvtl"
	case ARM64_FCVTL2:
		return "fcvtl2"
	case ARM64_FCVTLT:
		return "fcvtlt"
	case ARM64_FCVTMS:
		return "fcvtms"
	case ARM64_FCVTMU:
		return "fcvtmu"
	case ARM64_FCVTN:
		return "fcvtn"
	case ARM64_FCVTN2:
		return "fcvtn2"
	case ARM64_FCVTNS:
		return "fcvtns"
	case ARM64_FCVTNT:
		return "fcvtnt"
	case ARM64_FCVTNU:
		return "fcvtnu"
	case ARM64_FCVTPS:
		return "fcvtps"
	case ARM64_FCVTPU:
		return "fcvtpu"
	case ARM64_FCVTX:
		return "fcvtx"
	case ARM64_FCVTXN:
		return "fcvtxn"
	case ARM64_FCVTXN2:
		return "fcvtxn2"
	case ARM64_FCVTXNT:
		return "fcvtxnt"
	case ARM64_FCVTZS:
		return "fcvtzs"
	case ARM64_FCVTZU:
		return "fcvtzu"
	case ARM64_FDIV:
		return "fdiv"
	case ARM64_FDIVR:
		return "fdivr"
	case ARM64_FDUP:
		return "fdup"
	case ARM64_FEXPA:
		return "fexpa"
	case ARM64_FJCVTZS:
		return "fjcvtzs"
	case ARM64_FLOGB:
		return "flogb"
	case ARM64_FMAD:
		return "fmad"
	case ARM64_FMADD:
		return "fmadd"
	case ARM64_FMAX:
		return "fmax"
	case ARM64_FMAXNM:
		return "fmaxnm"
	case ARM64_FMAXNMP:
		return "fmaxnmp"
	case ARM64_FMAXNMV:
		return "fmaxnmv"
	case ARM64_FMAXP:
		return "fmaxp"
	case ARM64_FMAXV:
		return "fmaxv"
	case ARM64_FMIN:
		return "fmin"
	case ARM64_FMINNM:
		return "fminnm"
	case ARM64_FMINNMP:
		return "fminnmp"
	case ARM64_FMINNMV:
		return "fminnmv"
	case ARM64_FMINP:
		return "fminp"
	case ARM64_FMINV:
		return "fminv"
	case ARM64_FMLA:
		return "fmla"
	case ARM64_FMLAL:
		return "fmlal"
	case ARM64_FMLAL2:
		return "fmlal2"
	case ARM64_FMLALB:
		return "fmlalb"
	case ARM64_FMLALT:
		return "fmlalt"
	case ARM64_FMLS:
		return "fmls"
	case ARM64_FMLSL:
		return "fmlsl"
	case ARM64_FMLSL2:
		return "fmlsl2"
	case ARM64_FMLSLB:
		return "fmlslb"
	case ARM64_FMLSLT:
		return "fmlslt"
	case ARM64_FMMLA:
		return "fmmla"
	case ARM64_FMOPA:
		return "fmopa"
	case ARM64_FMOPS:
		return "fmops"
	case ARM64_FMOV:
		return "fmov"
	case ARM64_FMSB:
		return "fmsb"
	case ARM64_FMSUB:
		return "fmsub"
	case ARM64_FMUL:
		return "fmul"
	case ARM64_FMULX:
		return "fmulx"
	case ARM64_FNEG:
		return "fneg"
	case ARM64_FNMAD:
		return "fnmad"
	case ARM64_FNMADD:
		return "fnmadd"
	case ARM64_FNMLA:
		return "fnmla"
	case ARM64_FNMLS:
		return "fnmls"
	case ARM64_FNMSB:
		return "fnmsb"
	case ARM64_FNMSUB:
		return "fnmsub"
	case ARM64_FNMUL:
		return "fnmul"
	case ARM64_FRECPE:
		return "frecpe"
	case ARM64_FRECPS:
		return "frecps"
	case ARM64_FRECPX:
		return "frecpx"
	case ARM64_FRINT32X:
		return "frint32x"
	case ARM64_FRINT32Z:
		return "frint32z"
	case ARM64_FRINT64X:
		return "frint64x"
	case ARM64_FRINT64Z:
		return "frint64z"
	case ARM64_FRINTA:
		return "frinta"
	case ARM64_FRINTI:
		return "frinti"
	case ARM64_FRINTM:
		return "frintm"
	case ARM64_FRINTN:
		return "frintn"
	case ARM64_FRINTP:
		return "frintp"
	case ARM64_FRINTX:
		return "frintx"
	case ARM64_FRINTZ:
		return "frintz"
	case ARM64_FRSQRTE:
		return "frsqrte"
	case ARM64_FRSQRTS:
		return "frsqrts"
	case ARM64_FSCALE:
		return "fscale"
	case ARM64_FSQRT:
		return "fsqrt"
	case ARM64_FSUB:
		return "fsub"
	case ARM64_FSUBR:
		return "fsubr"
	case ARM64_FTMAD:
		return "ftmad"
	case ARM64_FTSMUL:
		return "ftsmul"
	case ARM64_FTSSEL:
		return "ftssel"
	case ARM64_GMI:
		return "gmi"
	case ARM64_HINT:
		return "hint"
	case ARM64_HISTCNT:
		return "histcnt"
	case ARM64_HISTSEG:
		return "histseg"
	case ARM64_HLT:
		return "hlt"
	case ARM64_HVC:
		return "hvc"
	case ARM64_IC:
		return "ic"
	case ARM64_INCB:
		return "incb"
	case ARM64_INCD:
		return "incd"
	case ARM64_INCH:
		return "inch"
	case ARM64_INCP:
		return "incp"
	case ARM64_INCW:
		return "incw"
	case ARM64_INDEX:
		return "index"
	case ARM64_INS:
		return "ins"
	case ARM64_INSR:
		return "insr"
	case ARM64_IRG:
		return "irg"
	case ARM64_ISB:
		return "isb"
	case ARM64_LASTA:
		return "lasta"
	case ARM64_LASTB:
		return "lastb"
	case ARM64_LD1:
		return "ld1"
	case ARM64_LD1B:
		return "ld1b"
	case ARM64_LD1D:
		return "ld1d"
	case ARM64_LD1H:
		return "ld1h"
	case ARM64_LD1Q:
		return "ld1q"
	case ARM64_LD1R:
		return "ld1r"
	case ARM64_LD1RB:
		return "ld1rb"
	case ARM64_LD1RD:
		return "ld1rd"
	case ARM64_LD1RH:
		return "ld1rh"
	case ARM64_LD1ROB:
		return "ld1rob"
	case ARM64_LD1ROD:
		return "ld1rod"
	case ARM64_LD1ROH:
		return "ld1roh"
	case ARM64_LD1ROW:
		return "ld1row"
	case ARM64_LD1RQB:
		return "ld1rqb"
	case ARM64_LD1RQD:
		return "ld1rqd"
	case ARM64_LD1RQH:
		return "ld1rqh"
	case ARM64_LD1RQW:
		return "ld1rqw"
	case ARM64_LD1RSB:
		return "ld1rsb"
	case ARM64_LD1RSH:
		return "ld1rsh"
	case ARM64_LD1RSW:
		return "ld1rsw"
	case ARM64_LD1RW:
		return "ld1rw"
	case ARM64_LD1SB:
		return "ld1sb"
	case ARM64_LD1SH:
		return "ld1sh"
	case ARM64_LD1SW:
		return "ld1sw"
	case ARM64_LD1W:
		return "ld1w"
	case ARM64_LD2:
		return "ld2"
	case ARM64_LD2B:
		return "ld2b"
	case ARM64_LD2D:
		return "ld2d"
	case ARM64_LD2H:
		return "ld2h"
	case ARM64_LD2R:
		return "ld2r"
	case ARM64_LD2W:
		return "ld2w"
	case ARM64_LD3:
		return "ld3"
	case ARM64_LD3B:
		return "ld3b"
	case ARM64_LD3D:
		return "ld3d"
	case ARM64_LD3H:
		return "ld3h"
	case ARM64_LD3R:
		return "ld3r"
	case ARM64_LD3W:
		return "ld3w"
	case ARM64_LD4:
		return "ld4"
	case ARM64_LD4B:
		return "ld4b"
	case ARM64_LD4D:
		return "ld4d"
	case ARM64_LD4H:
		return "ld4h"
	case ARM64_LD4R:
		return "ld4r"
	case ARM64_LD4W:
		return "ld4w"
	case ARM64_LD64B:
		return "ld64b"
	case ARM64_LDADD:
		return "ldadd"
	case ARM64_LDADDA:
		return "ldadda"
	case ARM64_LDADDAB:
		return "ldaddab"
	case ARM64_LDADDAH:
		return "ldaddah"
	case ARM64_LDADDAL:
		return "ldaddal"
	case ARM64_LDADDALB:
		return "ldaddalb"
	case ARM64_LDADDALH:
		return "ldaddalh"
	case ARM64_LDADDB:
		return "ldaddb"
	case ARM64_LDADDH:
		return "ldaddh"
	case ARM64_LDADDL:
		return "ldaddl"
	case ARM64_LDADDLB:
		return "ldaddlb"
	case ARM64_LDADDLH:
		return "ldaddlh"
	case ARM64_LDAPR:
		return "ldapr"
	case ARM64_LDAPRB:
		return "ldaprb"
	case ARM64_LDAPRH:
		return "ldaprh"
	case ARM64_LDAPUR:
		return "ldapur"
	case ARM64_LDAPURB:
		return "ldapurb"
	case ARM64_LDAPURH:
		return "ldapurh"
	case ARM64_LDAPURSB:
		return "ldapursb"
	case ARM64_LDAPURSH:
		return "ldapursh"
	case ARM64_LDAPURSW:
		return "ldapursw"
	case ARM64_LDAR:
		return "ldar"
	case ARM64_LDARB:
		return "ldarb"
	case ARM64_LDARH:
		return "ldarh"
	case ARM64_LDAXP:
		return "ldaxp"
	case ARM64_LDAXR:
		return "ldaxr"
	case ARM64_LDAXRB:
		return "ldaxrb"
	case ARM64_LDAXRH:
		return "ldaxrh"
	case ARM64_LDCLR:
		return "ldclr"
	case ARM64_LDCLRA:
		return "ldclra"
	case ARM64_LDCLRAB:
		return "ldclrab"
	case ARM64_LDCLRAH:
		return "ldclrah"
	case ARM64_LDCLRAL:
		return "ldclral"
	case ARM64_LDCLRALB:
		return "ldclralb"
	case ARM64_LDCLRALH:
		return "ldclralh"
	case ARM64_LDCLRB:
		return "ldclrb"
	case ARM64_LDCLRH:
		return "ldclrh"
	case ARM64_LDCLRL:
		return "ldclrl"
	case ARM64_LDCLRLB:
		return "ldclrlb"
	case ARM64_LDCLRLH:
		return "ldclrlh"
	case ARM64_LDEOR:
		return "ldeor"
	case ARM64_LDEORA:
		return "ldeora"
	case ARM64_LDEORAB:
		return "ldeorab"
	case ARM64_LDEORAH:
		return "ldeorah"
	case ARM64_LDEORAL:
		return "ldeoral"
	case ARM64_LDEORALB:
		return "ldeoralb"
	case ARM64_LDEORALH:
		return "ldeoralh"
	case ARM64_LDEORB:
		return "ldeorb"
	case ARM64_LDEORH:
		return "ldeorh"
	case ARM64_LDEORL:
		return "ldeorl"
	case ARM64_LDEORLB:
		return "ldeorlb"
	case ARM64_LDEORLH:
		return "ldeorlh"
	case ARM64_LDFF1B:
		return "ldff1b"
	case ARM64_LDFF1D:
		return "ldff1d"
	case ARM64_LDFF1H:
		return "ldff1h"
	case ARM64_LDFF1SB:
		return "ldff1sb"
	case ARM64_LDFF1SH:
		return "ldff1sh"
	case ARM64_LDFF1SW:
		return "ldff1sw"
	case ARM64_LDFF1W:
		return "ldff1w"
	case ARM64_LDG:
		return "ldg"
	case ARM64_LDGM:
		return "ldgm"
	case ARM64_LDLAR:
		return "ldlar"
	case ARM64_LDLARB:
		return "ldlarb"
	case ARM64_LDLARH:
		return "ldlarh"
	case ARM64_LDNF1B:
		return "ldnf1b"
	case ARM64_LDNF1D:
		return "ldnf1d"
	case ARM64_LDNF1H:
		return "ldnf1h"
	case ARM64_LDNF1SB:
		return "ldnf1sb"
	case ARM64_LDNF1SH:
		return "ldnf1sh"
	case ARM64_LDNF1SW:
		return "ldnf1sw"
	case ARM64_LDNF1W:
		return "ldnf1w"
	case ARM64_LDNP:
		return "ldnp"
	case ARM64_LDNT1B:
		return "ldnt1b"
	case ARM64_LDNT1D:
		return "ldnt1d"
	case ARM64_LDNT1H:
		return "ldnt1h"
	case ARM64_LDNT1SB:
		return "ldnt1sb"
	case ARM64_LDNT1SH:
		return "ldnt1sh"
	case ARM64_LDNT1SW:
		return "ldnt1sw"
	case ARM64_LDNT1W:
		return "ldnt1w"
	case ARM64_LDP:
		return "ldp"
	case ARM64_LDPSW:
		return "ldpsw"
	case ARM64_LDR:
		return "ldr"
	case ARM64_LDRAA:
		return "ldraa"
	case ARM64_LDRAB:
		return "ldrab"
	case ARM64_LDRB:
		return "ldrb"
	case ARM64_LDRH:
		return "ldrh"
	case ARM64_LDRSB:
		return "ldrsb"
	case ARM64_LDRSH:
		return "ldrsh"
	case ARM64_LDRSW:
		return "ldrsw"
	case ARM64_LDSET:
		return "ldset"
	case ARM64_LDSETA:
		return "ldseta"
	case ARM64_LDSETAB:
		return "ldsetab"
	case ARM64_LDSETAH:
		return "ldsetah"
	case ARM64_LDSETAL:
		return "ldsetal"
	case ARM64_LDSETALB:
		return "ldsetalb"
	case ARM64_LDSETALH:
		return "ldsetalh"
	case ARM64_LDSETB:
		return "ldsetb"
	case ARM64_LDSETH:
		return "ldseth"
	case ARM64_LDSETL:
		return "ldsetl"
	case ARM64_LDSETLB:
		return "ldsetlb"
	case ARM64_LDSETLH:
		return "ldsetlh"
	case ARM64_LDSMAX:
		return "ldsmax"
	case ARM64_LDSMAXA:
		return "ldsmaxa"
	case ARM64_LDSMAXAB:
		return "ldsmaxab"
	case ARM64_LDSMAXAH:
		return "ldsmaxah"
	case ARM64_LDSMAXAL:
		return "ldsmaxal"
	case ARM64_LDSMAXALB:
		return "ldsmaxalb"
	case ARM64_LDSMAXALH:
		return "ldsmaxalh"
	case ARM64_LDSMAXB:
		return "ldsmaxb"
	case ARM64_LDSMAXH:
		return "ldsmaxh"
	case ARM64_LDSMAXL:
		return "ldsmaxl"
	case ARM64_LDSMAXLB:
		return "ldsmaxlb"
	case ARM64_LDSMAXLH:
		return "ldsmaxlh"
	case ARM64_LDSMIN:
		return "ldsmin"
	case ARM64_LDSMINA:
		return "ldsmina"
	case ARM64_LDSMINAB:
		return "ldsminab"
	case ARM64_LDSMINAH:
		return "ldsminah"
	case ARM64_LDSMINAL:
		return "ldsminal"
	case ARM64_LDSMINALB:
		return "ldsminalb"
	case ARM64_LDSMINALH:
		return "ldsminalh"
	case ARM64_LDSMINB:
		return "ldsminb"
	case ARM64_LDSMINH:
		return "ldsminh"
	case ARM64_LDSMINL:
		return "ldsminl"
	case ARM64_LDSMINLB:
		return "ldsminlb"
	case ARM64_LDSMINLH:
		return "ldsminlh"
	case ARM64_LDTR:
		return "ldtr"
	case ARM64_LDTRB:
		return "ldtrb"
	case ARM64_LDTRH:
		return "ldtrh"
	case ARM64_LDTRSB:
		return "ldtrsb"
	case ARM64_LDTRSH:
		return "ldtrsh"
	case ARM64_LDTRSW:
		return "ldtrsw"
	case ARM64_LDUMAX:
		return "ldumax"
	case ARM64_LDUMAXA:
		return "ldumaxa"
	case ARM64_LDUMAXAB:
		return "ldumaxab"
	case ARM64_LDUMAXAH:
		return "ldumaxah"
	case ARM64_LDUMAXAL:
		return "ldumaxal"
	case ARM64_LDUMAXALB:
		return "ldumaxalb"
	case ARM64_LDUMAXALH:
		return "ldumaxalh"
	case ARM64_LDUMAXB:
		return "ldumaxb"
	case ARM64_LDUMAXH:
		return "ldumaxh"
	case ARM64_LDUMAXL:
		return "ldumaxl"
	case ARM64_LDUMAXLB:
		return "ldumaxlb"
	case ARM64_LDUMAXLH:
		return "ldumaxlh"
	case ARM64_LDUMIN:
		return "ldumin"
	case ARM64_LDUMINA:
		return "ldumina"
	case ARM64_LDUMINAB:
		return "lduminab"
	case ARM64_LDUMINAH:
		return "lduminah"
	case ARM64_LDUMINAL:
		return "lduminal"
	case ARM64_LDUMINALB:
		return "lduminalb"
	case ARM64_LDUMINALH:
		return "lduminalh"
	case ARM64_LDUMINB:
		return "lduminb"
	case ARM64_LDUMINH:
		return "lduminh"
	case ARM64_LDUMINL:
		return "lduminl"
	case ARM64_LDUMINLB:
		return "lduminlb"
	case ARM64_LDUMINLH:
		return "lduminlh"
	case ARM64_LDUR:
		return "ldur"
	case ARM64_LDURB:
		return "ldurb"
	case ARM64_LDURH:
		return "ldurh"
	case ARM64_LDURSB:
		return "ldursb"
	case ARM64_LDURSH:
		return "ldursh"
	case ARM64_LDURSW:
		return "ldursw"
	case ARM64_LDXP:
		return "ldxp"
	case ARM64_LDXR:
		return "ldxr"
	case ARM64_LDXRB:
		return "ldxrb"
	case ARM64_LDXRH:
		return "ldxrh"
	case ARM64_LSL:
		return "lsl"
	case ARM64_LSLR:
		return "lslr"
	case ARM64_LSLV:
		return "lslv"
	case ARM64_LSR:
		return "lsr"
	case ARM64_LSRR:
		return "lsrr"
	case ARM64_LSRV:
		return "lsrv"
	case ARM64_MAD:
		return "mad"
	case ARM64_MADD:
		return "madd"
	case ARM64_MATCH:
		return "match"
	case ARM64_MLA:
		return "mla"
	case ARM64_MLS:
		return "mls"
	case ARM64_MNEG:
		return "mneg"
	case ARM64_MOV:
		return "mov"
	case ARM64_MOVA:
		return "mova"
	case ARM64_MOVI:
		return "movi"
	case ARM64_MOVK:
		return "movk"
	case ARM64_MOVN:
		return "movn"
	case ARM64_MOVPRFX:
		return "movprfx"
	case ARM64_MOVS:
		return "movs"
	case ARM64_MOVZ:
		return "movz"
	case ARM64_MRS:
		return "mrs"
	case ARM64_MSB:
		return "msb"
	case ARM64_MSR:
		return "msr"
	case ARM64_MSUB:
		return "msub"
	case ARM64_MUL:
		return "mul"
	case ARM64_MVN:
		return "mvn"
	case ARM64_MVNI:
		return "mvni"
	case ARM64_NAND:
		return "nand"
	case ARM64_NANDS:
		return "nands"
	case ARM64_NBSL:
		return "nbsl"
	case ARM64_NEG:
		return "neg"
	case ARM64_NEGS:
		return "negs"
	case ARM64_NGC:
		return "ngc"
	case ARM64_NGCS:
		return "ngcs"
	case ARM64_NMATCH:
		return "nmatch"
	case ARM64_NOP:
		return "nop"
	case ARM64_NOR:
		return "nor"
	case ARM64_NORS:
		return "nors"
	case ARM64_NOT:
		return "not"
	case ARM64_NOTS:
		return "nots"
	case ARM64_ORN:
		return "orn"
	case ARM64_ORNS:
		return "orns"
	case ARM64_ORR:
		return "orr"
	case ARM64_ORRS:
		return "orrs"
	case ARM64_ORV:
		return "orv"
	case ARM64_PACDA:
		return "pacda"
	case ARM64_PACDB:
		return "pacdb"
	case ARM64_PACDZA:
		return "pacdza"
	case ARM64_PACDZB:
		return "pacdzb"
	case ARM64_PACGA:
		return "pacga"
	case ARM64_PACIA:
		return "pacia"
	case ARM64_PACIA1716:
		return "pacia1716"
	case ARM64_PACIASP:
		return "paciasp"
	case ARM64_PACIAZ:
		return "paciaz"
	case ARM64_PACIB:
		return "pacib"
	case ARM64_PACIB1716:
		return "pacib1716"
	case ARM64_PACIBSP:
		return "pacibsp"
	case ARM64_PACIBZ:
		return "pacibz"
	case ARM64_PACIZA:
		return "paciza"
	case ARM64_PACIZB:
		return "pacizb"
	case ARM64_PFALSE:
		return "pfalse"
	case ARM64_PFIRST:
		return "pfirst"
	case ARM64_PMUL:
		return "pmul"
	case ARM64_PMULL:
		return "pmull"
	case ARM64_PMULL2:
		return "pmull2"
	case ARM64_PMULLB:
		return "pmullb"
	case ARM64_PMULLT:
		return "pmullt"
	case ARM64_PNEXT:
		return "pnext"
	case ARM64_PRFB:
		return "prfb"
	case ARM64_PRFD:
		return "prfd"
	case ARM64_PRFH:
		return "prfh"
	case ARM64_PRFM:
		return "prfm"
	case ARM64_PRFUM:
		return "prfum"
	case ARM64_PRFW:
		return "prfw"
	case ARM64_PSB:
		return "psb"
	case ARM64_PSSBB:
		return "pssbb"
	case ARM64_PTEST:
		return "ptest"
	case ARM64_PTRUE:
		return "ptrue"
	case ARM64_PTRUES:
		return "ptrues"
	case ARM64_PUNPKHI:
		return "punpkhi"
	case ARM64_PUNPKLO:
		return "punpklo"
	case ARM64_RADDHN:
		return "raddhn"
	case ARM64_RADDHN2:
		return "raddhn2"
	case ARM64_RADDHNB:
		return "raddhnb"
	case ARM64_RADDHNT:
		return "raddhnt"
	case ARM64_RAX1:
		return "rax1"
	case ARM64_RBIT:
		return "rbit"
	case ARM64_RDFFR:
		return "rdffr"
	case ARM64_RDFFRS:
		return "rdffrs"
	case ARM64_RDVL:
		return "rdvl"
	case ARM64_RET:
		return "ret"
	case ARM64_RETAA:
		return "retaa"
	case ARM64_RETAB:
		return "retab"
	case ARM64_REV:
		return "rev"
	case ARM64_REV16:
		return "rev16"
	case ARM64_REV32:
		return "rev32"
	case ARM64_REV64:
		return "rev64"
	case ARM64_REVB:
		return "revb"
	case ARM64_REVD:
		return "revd"
	case ARM64_REVH:
		return "revh"
	case ARM64_REVW:
		return "revw"
	case ARM64_RMIF:
		return "rmif"
	case ARM64_ROR:
		return "ror"
	case ARM64_RORV:
		return "rorv"
	case ARM64_RSHRN:
		return "rshrn"
	case ARM64_RSHRN2:
		return "rshrn2"
	case ARM64_RSHRNB:
		return "rshrnb"
	case ARM64_RSHRNT:
		return "rshrnt"
	case ARM64_RSUBHN:
		return "rsubhn"
	case ARM64_RSUBHN2:
		return "rsubhn2"
	case ARM64_RSUBHNB:
		return "rsubhnb"
	case ARM64_RSUBHNT:
		return "rsubhnt"
	case ARM64_SABA:
		return "saba"
	case ARM64_SABAL:
		return "sabal"
	case ARM64_SABAL2:
		return "sabal2"
	case ARM64_SABALB:
		return "sabalb"
	case ARM64_SABALT:
		return "sabalt"
	case ARM64_SABD:
		return "sabd"
	case ARM64_SABDL:
		return "sabdl"
	case ARM64_SABDL2:
		return "sabdl2"
	case ARM64_SABDLB:
		return "sabdlb"
	case ARM64_SABDLT:
		return "sabdlt"
	case ARM64_SADALP:
		return "sadalp"
	case ARM64_SADDL:
		return "saddl"
	case ARM64_SADDL2:
		return "saddl2"
	case ARM64_SADDLB:
		return "saddlb"
	case ARM64_SADDLBT:
		return "saddlbt"
	case ARM64_SADDLP:
		return "saddlp"
	case ARM64_SADDLT:
		return "saddlt"
	case ARM64_SADDLV:
		return "saddlv"
	case ARM64_SADDV:
		return "saddv"
	case ARM64_SADDW:
		return "saddw"
	case ARM64_SADDW2:
		return "saddw2"
	case ARM64_SADDWB:
		return "saddwb"
	case ARM64_SADDWT:
		return "saddwt"
	case ARM64_SB:
		return "sb"
	case ARM64_SBC:
		return "sbc"
	case ARM64_SBCLB:
		return "sbclb"
	case ARM64_SBCLT:
		return "sbclt"
	case ARM64_SBCS:
		return "sbcs"
	case ARM64_SBFIZ:
		return "sbfiz"
	case ARM64_SBFM:
		return "sbfm"
	case ARM64_SBFX:
		return "sbfx"
	case ARM64_SCLAMP:
		return "sclamp"
	case ARM64_SCVTF:
		return "scvtf"
	case ARM64_SDIV:
		return "sdiv"
	case ARM64_SDIVR:
		return "sdivr"
	case ARM64_SDOT:
		return "sdot"
	case ARM64_SEL:
		return "sel"
	case ARM64_SETF16:
		return "setf16"
	case ARM64_SETF8:
		return "setf8"
	case ARM64_SETFFR:
		return "setffr"
	case ARM64_SEV:
		return "sev"
	case ARM64_SEVL:
		return "sevl"
	case ARM64_SHA1C:
		return "sha1c"
	case ARM64_SHA1H:
		return "sha1h"
	case ARM64_SHA1M:
		return "sha1m"
	case ARM64_SHA1P:
		return "sha1p"
	case ARM64_SHA1SU0:
		return "sha1su0"
	case ARM64_SHA1SU1:
		return "sha1su1"
	case ARM64_SHA256H:
		return "sha256h"
	case ARM64_SHA256H2:
		return "sha256h2"
	case ARM64_SHA256SU0:
		return "sha256su0"
	case ARM64_SHA256SU1:
		return "sha256su1"
	case ARM64_SHA512H:
		return "sha512h"
	case ARM64_SHA512H2:
		return "sha512h2"
	case ARM64_SHA512SU0:
		return "sha512su0"
	case ARM64_SHA512SU1:
		return "sha512su1"
	case ARM64_SHADD:
		return "shadd"
	case ARM64_SHL:
		return "shl"
	case ARM64_SHLL:
		return "shll"
	case ARM64_SHLL2:
		return "shll2"
	case ARM64_SHRN:
		return "shrn"
	case ARM64_SHRN2:
		return "shrn2"
	case ARM64_SHRNB:
		return "shrnb"
	case ARM64_SHRNT:
		return "shrnt"
	case ARM64_SHSUB:
		return "shsub"
	case ARM64_SHSUBR:
		return "shsubr"
	case ARM64_SLI:
		return "sli"
	case ARM64_SM3PARTW1:
		return "sm3partw1"
	case ARM64_SM3PARTW2:
		return "sm3partw2"
	case ARM64_SM3SS1:
		return "sm3ss1"
	case ARM64_SM3TT1A:
		return "sm3tt1a"
	case ARM64_SM3TT1B:
		return "sm3tt1b"
	case ARM64_SM3TT2A:
		return "sm3tt2a"
	case ARM64_SM3TT2B:
		return "sm3tt2b"
	case ARM64_SM4E:
		return "sm4e"
	case ARM64_SM4EKEY:
		return "sm4ekey"
	case ARM64_SMADDL:
		return "smaddl"
	case ARM64_SMAX:
		return "smax"
	case ARM64_SMAXP:
		return "smaxp"
	case ARM64_SMAXV:
		return "smaxv"
	case ARM64_SMC:
		return "smc"
	case ARM64_SMIN:
		return "smin"
	case ARM64_SMINP:
		return "sminp"
	case ARM64_SMINV:
		return "sminv"
	case ARM64_SMLAL:
		return "smlal"
	case ARM64_SMLAL2:
		return "smlal2"
	case ARM64_SMLALB:
		return "smlalb"
	case ARM64_SMLALT:
		return "smlalt"
	case ARM64_SMLSL:
		return "smlsl"
	case ARM64_SMLSL2:
		return "smlsl2"
	case ARM64_SMLSLB:
		return "smlslb"
	case ARM64_SMLSLT:
		return "smlslt"
	case ARM64_SMMLA:
		return "smmla"
	case ARM64_SMNEGL:
		return "smnegl"
	case ARM64_SMOPA:
		return "smopa"
	case ARM64_SMOPS:
		return "smops"
	case ARM64_SMOV:
		return "smov"
	case ARM64_SMSTART:
		return "smstart"
	case ARM64_SMSTOP:
		return "smstop"
	case ARM64_SMSUBL:
		return "smsubl"
	case ARM64_SMULH:
		return "smulh"
	case ARM64_SMULL:
		return "smull"
	case ARM64_SMULL2:
		return "smull2"
	case ARM64_SMULLB:
		return "smullb"
	case ARM64_SMULLT:
		return "smullt"
	case ARM64_SPLICE:
		return "splice"
	case ARM64_SQABS:
		return "sqabs"
	case ARM64_SQADD:
		return "sqadd"
	case ARM64_SQCADD:
		return "sqcadd"
	case ARM64_SQDECB:
		return "sqdecb"
	case ARM64_SQDECD:
		return "sqdecd"
	case ARM64_SQDECH:
		return "sqdech"
	case ARM64_SQDECP:
		return "sqdecp"
	case ARM64_SQDECW:
		return "sqdecw"
	case ARM64_SQDMLAL:
		return "sqdmlal"
	case ARM64_SQDMLAL2:
		return "sqdmlal2"
	case ARM64_SQDMLALB:
		return "sqdmlalb"
	case ARM64_SQDMLALBT:
		return "sqdmlalbt"
	case ARM64_SQDMLALT:
		return "sqdmlalt"
	case ARM64_SQDMLSL:
		return "sqdmlsl"
	case ARM64_SQDMLSL2:
		return "sqdmlsl2"
	case ARM64_SQDMLSLB:
		return "sqdmlslb"
	case ARM64_SQDMLSLBT:
		return "sqdmlslbt"
	case ARM64_SQDMLSLT:
		return "sqdmlslt"
	case ARM64_SQDMULH:
		return "sqdmulh"
	case ARM64_SQDMULL:
		return "sqdmull"
	case ARM64_SQDMULL2:
		return "sqdmull2"
	case ARM64_SQDMULLB:
		return "sqdmullb"
	case ARM64_SQDMULLT:
		return "sqdmullt"
	case ARM64_SQINCB:
		return "sqincb"
	case ARM64_SQINCD:
		return "sqincd"
	case ARM64_SQINCH:
		return "sqinch"
	case ARM64_SQINCP:
		return "sqincp"
	case ARM64_SQINCW:
		return "sqincw"
	case ARM64_SQNEG:
		return "sqneg"
	case ARM64_SQRDCMLAH:
		return "sqrdcmlah"
	case ARM64_SQRDMLAH:
		return "sqrdmlah"
	case ARM64_SQRDMLSH:
		return "sqrdmlsh"
	case ARM64_SQRDMULH:
		return "sqrdmulh"
	case ARM64_SQRSHL:
		return "sqrshl"
	case ARM64_SQRSHLR:
		return "sqrshlr"
	case ARM64_SQRSHRN:
		return "sqrshrn"
	case ARM64_SQRSHRN2:
		return "sqrshrn2"
	case ARM64_SQRSHRNB:
		return "sqrshrnb"
	case ARM64_SQRSHRNT:
		return "sqrshrnt"
	case ARM64_SQRSHRUN:
		return "sqrshrun"
	case ARM64_SQRSHRUN2:
		return "sqrshrun2"
	case ARM64_SQRSHRUNB:
		return "sqrshrunb"
	case ARM64_SQRSHRUNT:
		return "sqrshrunt"
	case ARM64_SQSHL:
		return "sqshl"
	case ARM64_SQSHLR:
		return "sqshlr"
	case ARM64_SQSHLU:
		return "sqshlu"
	case ARM64_SQSHRN:
		return "sqshrn"
	case ARM64_SQSHRN2:
		return "sqshrn2"
	case ARM64_SQSHRNB:
		return "sqshrnb"
	case ARM64_SQSHRNT:
		return "sqshrnt"
	case ARM64_SQSHRUN:
		return "sqshrun"
	case ARM64_SQSHRUN2:
		return "sqshrun2"
	case ARM64_SQSHRUNB:
		return "sqshrunb"
	case ARM64_SQSHRUNT:
		return "sqshrunt"
	case ARM64_SQSUB:
		return "sqsub"
	case ARM64_SQSUBR:
		return "sqsubr"
	case ARM64_SQXTN:
		return "sqxtn"
	case ARM64_SQXTN2:
		return "sqxtn2"
	case ARM64_SQXTNB:
		return "sqxtnb"
	case ARM64_SQXTNT:
		return "sqxtnt"
	case ARM64_SQXTUN:
		return "sqxtun"
	case ARM64_SQXTUN2:
		return "sqxtun2"
	case ARM64_SQXTUNB:
		return "sqxtunb"
	case ARM64_SQXTUNT:
		return "sqxtunt"
	case ARM64_SRHADD:
		return "srhadd"
	case ARM64_SRI:
		return "sri"
	case ARM64_SRSHL:
		return "srshl"
	case ARM64_SRSHLR:
		return "srshlr"
	case ARM64_SRSHR:
		return "srshr"
	case ARM64_SRSRA:
		return "srsra"
	case ARM64_SSBB:
		return "ssbb"
	case ARM64_SSHL:
		return "sshl"
	case ARM64_SSHLL:
		return "sshll"
	case ARM64_SSHLL2:
		return "sshll2"
	case ARM64_SSHLLB:
		return "sshllb"
	case ARM64_SSHLLT:
		return "sshllt"
	case ARM64_SSHR:
		return "sshr"
	case ARM64_SSRA:
		return "ssra"
	case ARM64_SSUBL:
		return "ssubl"
	case ARM64_SSUBL2:
		return "ssubl2"
	case ARM64_SSUBLB:
		return "ssublb"
	case ARM64_SSUBLBT:
		return "ssublbt"
	case ARM64_SSUBLT:
		return "ssublt"
	case ARM64_SSUBLTB:
		return "ssubltb"
	case ARM64_SSUBW:
		return "ssubw"
	case ARM64_SSUBW2:
		return "ssubw2"
	case ARM64_SSUBWB:
		return "ssubwb"
	case ARM64_SSUBWT:
		return "ssubwt"
	case ARM64_ST1:
		return "st1"
	case ARM64_ST1B:
		return "st1b"
	case ARM64_ST1D:
		return "st1d"
	case ARM64_ST1H:
		return "st1h"
	case ARM64_ST1Q:
		return "st1q"
	case ARM64_ST1W:
		return "st1w"
	case ARM64_ST2:
		return "st2"
	case ARM64_ST2B:
		return "st2b"
	case ARM64_ST2D:
		return "st2d"
	case ARM64_ST2G:
		return "st2g"
	case ARM64_ST2H:
		return "st2h"
	case ARM64_ST2W:
		return "st2w"
	case ARM64_ST3:
		return "st3"
	case ARM64_ST3B:
		return "st3b"
	case ARM64_ST3D:
		return "st3d"
	case ARM64_ST3H:
		return "st3h"
	case ARM64_ST3W:
		return "st3w"
	case ARM64_ST4:
		return "st4"
	case ARM64_ST4B:
		return "st4b"
	case ARM64_ST4D:
		return "st4d"
	case ARM64_ST4H:
		return "st4h"
	case ARM64_ST4W:
		return "st4w"
	case ARM64_ST64B:
		return "st64b"
	case ARM64_ST64BV:
		return "st64bv"
	case ARM64_ST64BV0:
		return "st64bv0"
	case ARM64_STADD:
		return "stadd"
	case ARM64_STADDB:
		return "staddb"
	case ARM64_STADDH:
		return "staddh"
	case ARM64_STADDL:
		return "staddl"
	case ARM64_STADDLB:
		return "staddlb"
	case ARM64_STADDLH:
		return "staddlh"
	case ARM64_STCLR:
		return "stclr"
	case ARM64_STCLRB:
		return "stclrb"
	case ARM64_STCLRH:
		return "stclrh"
	case ARM64_STCLRL:
		return "stclrl"
	case ARM64_STCLRLB:
		return "stclrlb"
	case ARM64_STCLRLH:
		return "stclrlh"
	case ARM64_STEOR:
		return "steor"
	case ARM64_STEORB:
		return "steorb"
	case ARM64_STEORH:
		return "steorh"
	case ARM64_STEORL:
		return "steorl"
	case ARM64_STEORLB:
		return "steorlb"
	case ARM64_STEORLH:
		return "steorlh"
	case ARM64_STG:
		return "stg"
	case ARM64_STGM:
		return "stgm"
	case ARM64_STGP:
		return "stgp"
	case ARM64_STLLR:
		return "stllr"
	case ARM64_STLLRB:
		return "stllrb"
	case ARM64_STLLRH:
		return "stllrh"
	case ARM64_STLR:
		return "stlr"
	case ARM64_STLRB:
		return "stlrb"
	case ARM64_STLRH:
		return "stlrh"
	case ARM64_STLUR:
		return "stlur"
	case ARM64_STLURB:
		return "stlurb"
	case ARM64_STLURH:
		return "stlurh"
	case ARM64_STLXP:
		return "stlxp"
	case ARM64_STLXR:
		return "stlxr"
	case ARM64_STLXRB:
		return "stlxrb"
	case ARM64_STLXRH:
		return "stlxrh"
	case ARM64_STNP:
		return "stnp"
	case ARM64_STNT1B:
		return "stnt1b"
	case ARM64_STNT1D:
		return "stnt1d"
	case ARM64_STNT1H:
		return "stnt1h"
	case ARM64_STNT1W:
		return "stnt1w"
	case ARM64_STP:
		return "stp"
	case ARM64_STR:
		return "str"
	case ARM64_STRB:
		return "strb"
	case ARM64_STRH:
		return "strh"
	case ARM64_STSET:
		return "stset"
	case ARM64_STSETB:
		return "stsetb"
	case ARM64_STSETH:
		return "stseth"
	case ARM64_STSETL:
		return "stsetl"
	case ARM64_STSETLB:
		return "stsetlb"
	case ARM64_STSETLH:
		return "stsetlh"
	case ARM64_STSMAX:
		return "stsmax"
	case ARM64_STSMAXB:
		return "stsmaxb"
	case ARM64_STSMAXH:
		return "stsmaxh"
	case ARM64_STSMAXL:
		return "stsmaxl"
	case ARM64_STSMAXLB:
		return "stsmaxlb"
	case ARM64_STSMAXLH:
		return "stsmaxlh"
	case ARM64_STSMIN:
		return "stsmin"
	case ARM64_STSMINB:
		return "stsminb"
	case ARM64_STSMINH:
		return "stsminh"
	case ARM64_STSMINL:
		return "stsminl"
	case ARM64_STSMINLB:
		return "stsminlb"
	case ARM64_STSMINLH:
		return "stsminlh"
	case ARM64_STTR:
		return "sttr"
	case ARM64_STTRB:
		return "sttrb"
	case ARM64_STTRH:
		return "sttrh"
	case ARM64_STUMAX:
		return "stumax"
	case ARM64_STUMAXB:
		return "stumaxb"
	case ARM64_STUMAXH:
		return "stumaxh"
	case ARM64_STUMAXL:
		return "stumaxl"
	case ARM64_STUMAXLB:
		return "stumaxlb"
	case ARM64_STUMAXLH:
		return "stumaxlh"
	case ARM64_STUMIN:
		return "stumin"
	case ARM64_STUMINB:
		return "stuminb"
	case ARM64_STUMINH:
		return "stuminh"
	case ARM64_STUMINL:
		return "stuminl"
	case ARM64_STUMINLB:
		return "stuminlb"
	case ARM64_STUMINLH:
		return "stuminlh"
	case ARM64_STUR:
		return "stur"
	case ARM64_STURB:
		return "sturb"
	case ARM64_STURH:
		return "sturh"
	case ARM64_STXP:
		return "stxp"
	case ARM64_STXR:
		return "stxr"
	case ARM64_STXRB:
		return "stxrb"
	case ARM64_STXRH:
		return "stxrh"
	case ARM64_STZ2G:
		return "stz2g"
	case ARM64_STZG:
		return "stzg"
	case ARM64_STZGM:
		return "stzgm"
	case ARM64_SUB:
		return "sub"
	case ARM64_SUBG:
		return "subg"
	case ARM64_SUBHN:
		return "subhn"
	case ARM64_SUBHN2:
		return "subhn2"
	case ARM64_SUBHNB:
		return "subhnb"
	case ARM64_SUBHNT:
		return "subhnt"
	case ARM64_SUBP:
		return "subp"
	case ARM64_SUBPS:
		return "subps"
	case ARM64_SUBR:
		return "subr"
	case ARM64_SUBS:
		return "subs"
	case ARM64_SUDOT:
		return "sudot"
	case ARM64_SUMOPA:
		return "sumopa"
	case ARM64_SUMOPS:
		return "sumops"
	case ARM64_SUNPKHI:
		return "sunpkhi"
	case ARM64_SUNPKLO:
		return "sunpklo"
	case ARM64_SUQADD:
		return "suqadd"
	case ARM64_SVC:
		return "svc"
	case ARM64_SWP:
		return "swp"
	case ARM64_SWPA:
		return "swpa"
	case ARM64_SWPAB:
		return "swpab"
	case ARM64_SWPAH:
		return "swpah"
	case ARM64_SWPAL:
		return "swpal"
	case ARM64_SWPALB:
		return "swpalb"
	case ARM64_SWPALH:
		return "swpalh"
	case ARM64_SWPB:
		return "swpb"
	case ARM64_SWPH:
		return "swph"
	case ARM64_SWPL:
		return "swpl"
	case ARM64_SWPLB:
		return "swplb"
	case ARM64_SWPLH:
		return "swplh"
	case ARM64_SXTB:
		return "sxtb"
	case ARM64_SXTH:
		return "sxth"
	case ARM64_SXTL:
		return "sxtl"
	case ARM64_SXTL2:
		return "sxtl2"
	case ARM64_SXTW:
		return "sxtw"
	case ARM64_SYS:
		return "sys"
	case ARM64_SYSL:
		return "sysl"
	case ARM64_TBL:
		return "tbl"
	case ARM64_TBNZ:
		return "tbnz"
	case ARM64_TBX:
		return "tbx"
	case ARM64_TBZ:
		return "tbz"
	case ARM64_TCANCEL:
		return "tcancel"
	case ARM64_TCOMMIT:
		return "tcommit"
	case ARM64_TLBI:
		return "tlbi"
	case ARM64_TRN1:
		return "trn1"
	case ARM64_TRN2:
		return "trn2"
	case ARM64_TSB:
		return "tsb"
	case ARM64_TST:
		return "tst"
	case ARM64_TSTART:
		return "tstart"
	case ARM64_TTEST:
		return "ttest"
	case ARM64_UABA:
		return "uaba"
	case ARM64_UABAL:
		return "uabal"
	case ARM64_UABAL2:
		return "uabal2"
	case ARM64_UABALB:
		return "uabalb"
	case ARM64_UABALT:
		return "uabalt"
	case ARM64_UABD:
		return "uabd"
	case ARM64_UABDL:
		return "uabdl"
	case ARM64_UABDL2:
		return "uabdl2"
	case ARM64_UABDLB:
		return "uabdlb"
	case ARM64_UABDLT:
		return "uabdlt"
	case ARM64_UADALP:
		return "uadalp"
	case ARM64_UADDL:
		return "uaddl"
	case ARM64_UADDL2:
		return "uaddl2"
	case ARM64_UADDLB:
		return "uaddlb"
	case ARM64_UADDLP:
		return "uaddlp"
	case ARM64_UADDLT:
		return "uaddlt"
	case ARM64_UADDLV:
		return "uaddlv"
	case ARM64_UADDV:
		return "uaddv"
	case ARM64_UADDW:
		return "uaddw"
	case ARM64_UADDW2:
		return "uaddw2"
	case ARM64_UADDWB:
		return "uaddwb"
	case ARM64_UADDWT:
		return "uaddwt"
	case ARM64_UBFIZ:
		return "ubfiz"
	case ARM64_UBFM:
		return "ubfm"
	case ARM64_UBFX:
		return "ubfx"
	case ARM64_UCLAMP:
		return "uclamp"
	case ARM64_UCVTF:
		return "ucvtf"
	case ARM64_UDF:
		return "udf"
	case ARM64_UDIV:
		return "udiv"
	case ARM64_UDIVR:
		return "udivr"
	case ARM64_UDOT:
		return "udot"
	case ARM64_UHADD:
		return "uhadd"
	case ARM64_UHSUB:
		return "uhsub"
	case ARM64_UHSUBR:
		return "uhsubr"
	case ARM64_UMADDL:
		return "umaddl"
	case ARM64_UMAX:
		return "umax"
	case ARM64_UMAXP:
		return "umaxp"
	case ARM64_UMAXV:
		return "umaxv"
	case ARM64_UMIN:
		return "umin"
	case ARM64_UMINP:
		return "uminp"
	case ARM64_UMINV:
		return "uminv"
	case ARM64_UMLAL:
		return "umlal"
	case ARM64_UMLAL2:
		return "umlal2"
	case ARM64_UMLALB:
		return "umlalb"
	case ARM64_UMLALT:
		return "umlalt"
	case ARM64_UMLSL:
		return "umlsl"
	case ARM64_UMLSL2:
		return "umlsl2"
	case ARM64_UMLSLB:
		return "umlslb"
	case ARM64_UMLSLT:
		return "umlslt"
	case ARM64_UMMLA:
		return "ummla"
	case ARM64_UMNEGL:
		return "umnegl"
	case ARM64_UMOPA:
		return "umopa"
	case ARM64_UMOPS:
		return "umops"
	case ARM64_UMOV:
		return "umov"
	case ARM64_UMSUBL:
		return "umsubl"
	case ARM64_UMULH:
		return "umulh"
	case ARM64_UMULL:
		return "umull"
	case ARM64_UMULL2:
		return "umull2"
	case ARM64_UMULLB:
		return "umullb"
	case ARM64_UMULLT:
		return "umullt"
	case ARM64_UQADD:
		return "uqadd"
	case ARM64_UQDECB:
		return "uqdecb"
	case ARM64_UQDECD:
		return "uqdecd"
	case ARM64_UQDECH:
		return "uqdech"
	case ARM64_UQDECP:
		return "uqdecp"
	case ARM64_UQDECW:
		return "uqdecw"
	case ARM64_UQINCB:
		return "uqincb"
	case ARM64_UQINCD:
		return "uqincd"
	case ARM64_UQINCH:
		return "uqinch"
	case ARM64_UQINCP:
		return "uqincp"
	case ARM64_UQINCW:
		return "uqincw"
	case ARM64_UQRSHL:
		return "uqrshl"
	case ARM64_UQRSHLR:
		return "uqrshlr"
	case ARM64_UQRSHRN:
		return "uqrshrn"
	case ARM64_UQRSHRN2:
		return "uqrshrn2"
	case ARM64_UQRSHRNB:
		return "uqrshrnb"
	case ARM64_UQRSHRNT:
		return "uqrshrnt"
	case ARM64_UQSHL:
		return "uqshl"
	case ARM64_UQSHLR:
		return "uqshlr"
	case ARM64_UQSHRN:
		return "uqshrn"
	case ARM64_UQSHRN2:
		return "uqshrn2"
	case ARM64_UQSHRNB:
		return "uqshrnb"
	case ARM64_UQSHRNT:
		return "uqshrnt"
	case ARM64_UQSUB:
		return "uqsub"
	case ARM64_UQSUBR:
		return "uqsubr"
	case ARM64_UQXTN:
		return "uqxtn"
	case ARM64_UQXTN2:
		return "uqxtn2"
	case ARM64_UQXTNB:
		return "uqxtnb"
	case ARM64_UQXTNT:
		return "uqxtnt"
	case ARM64_URECPE:
		return "urecpe"
	case ARM64_URHADD:
		return "urhadd"
	case ARM64_URSHL:
		return "urshl"
	case ARM64_URSHLR:
		return "urshlr"
	case ARM64_URSHR:
		return "urshr"
	case ARM64_URSQRTE:
		return "ursqrte"
	case ARM64_URSRA:
		return "ursra"
	case ARM64_USDOT:
		return "usdot"
	case ARM64_USHL:
		return "ushl"
	case ARM64_USHLL:
		return "ushll"
	case ARM64_USHLL2:
		return "ushll2"
	case ARM64_USHLLB:
		return "ushllb"
	case ARM64_USHLLT:
		return "ushllt"
	case ARM64_USHR:
		return "ushr"
	case ARM64_USMMLA:
		return "usmmla"
	case ARM64_USMOPA:
		return "usmopa"
	case ARM64_USMOPS:
		return "usmops"
	case ARM64_USQADD:
		return "usqadd"
	case ARM64_USRA:
		return "usra"
	case ARM64_USUBL:
		return "usubl"
	case ARM64_USUBL2:
		return "usubl2"
	case ARM64_USUBLB:
		return "usublb"
	case ARM64_USUBLT:
		return "usublt"
	case ARM64_USUBW:
		return "usubw"
	case ARM64_USUBW2:
		return "usubw2"
	case ARM64_USUBWB:
		return "usubwb"
	case ARM64_USUBWT:
		return "usubwt"
	case ARM64_UUNPKHI:
		return "uunpkhi"
	case ARM64_UUNPKLO:
		return "uunpklo"
	case ARM64_UXTB:
		return "uxtb"
	case ARM64_UXTH:
		return "uxth"
	case ARM64_UXTL:
		return "uxtl"
	case ARM64_UXTL2:
		return "uxtl2"
	case ARM64_UXTW:
		return "uxtw"
	case ARM64_UZP1:
		return "uzp1"
	case ARM64_UZP2:
		return "uzp2"
	case ARM64_WFE:
		return "wfe"
	case ARM64_WFET:
		return "wfet"
	case ARM64_WFI:
		return "wfi"
	case ARM64_WFIT:
		return "wfit"
	case ARM64_WHILEGE:
		return "whilege"
	case ARM64_WHILEGT:
		return "whilegt"
	case ARM64_WHILEHI:
		return "whilehi"
	case ARM64_WHILEHS:
		return "whilehs"
	case ARM64_WHILELE:
		return "whilele"
	case ARM64_WHILELO:
		return "whilelo"
	case ARM64_WHILELS:
		return "whilels"
	case ARM64_WHILELT:
		return "whilelt"
	case ARM64_WHILERW:
		return "whilerw"
	case ARM64_WHILEWR:
		return "whilewr"
	case ARM64_WRFFR:
		return "wrffr"
	case ARM64_XAFLAG:
		return "xaflag"
	case ARM64_XAR:
		return "xar"
	case ARM64_XPACD:
		return "xpacd"
	case ARM64_XPACI:
		return "xpaci"
	case ARM64_XPACLRI:
		return "xpaclri"
	case ARM64_XTN:
		return "xtn"
	case ARM64_XTN2:
		return "xtn2"
	case ARM64_YIELD:
		return "yield"
	case ARM64_ZERO:
		return "zero"
	case ARM64_ZIP1:
		return "zip1"
	case ARM64_ZIP2:
		return "zip2"
	case ARM64_ERROR:
		fallthrough
	default:
		return "error"
	}
}
