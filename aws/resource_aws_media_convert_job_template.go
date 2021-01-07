package aws

import (
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsMediaConvertJobTemplate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsMediaConvertJobTemplateCreate,
		Read:   resourceAwsMediaConvertJobTemplateRead,
		Update: resourceAwsMediaConvertJobTemplateUpdate,
		Delete: resourceAwsMediaConvertJobTemplateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},
		Schema: map[string]*schema.Schema{
			"arn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"category": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"priority": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"queue": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"status_update_interval": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"acceleration_settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"mode": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice([]string{
								mediaconvert.AccelerationModeDisabled,
								mediaconvert.AccelerationModeEnabled,
								mediaconvert.AccelerationModePreferred,
							}, false),
						},
					},
				},
			},
			"hop_destinations": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"priority": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"queue": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"wait_minutes ": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntAtLeast(1),
						},
					},
				},
			},
			"settings": {
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ad_avail_offset": {
							Type:     schema.TypeInt,
							Optional: true,
						},
						"avail_blanking": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"avail_blanking_image": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"esam": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"manifest_confirm_condition_notification": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"mcc_xml": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"signal_processing_notification": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"scc_xml": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"response_signal_preroll": {
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"inputs": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// needed to change to a list of types to be able to parse it
									"audio_selector_groups": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_selector_group": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"audio_selector_names": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
														},
													},
												},
											},
										},
									},
									"audio_selectors": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_selector": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"custom_language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(3, 10),
															},
															"default_selection": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AudioDefaultSelection_Values(), false),
															},
															"external_audio_file_input": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
															},
															"offset": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"pids": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
															},
															"program_selection": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"remix_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"channel_mapping": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"output_channel": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"input_channels": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem:     &schema.Schema{Type: schema.TypeInt},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"channels_in": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																		"channels_out": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																	},
																},
															},
															"selector_type": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AudioSelectorType_Values(), false),
															},
															"tracks": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeInt},
															},
														},
													},
												},
											},
										},
									},
									"caption_selectors": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"caption_selector": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"name": {
																Type:     schema.TypeString,
																Required: true,
															},
															"custom_language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringLenBetween(3, 10),
															},
															"language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
															},
															"source_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"ancillary_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"convert_608_to_708": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AncillaryConvert608To708_Values(), false),
																					},
																					"source_ancillary_channel_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"terminate_captions": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AncillaryTerminateCaptions_Values(), false),
																					},
																				},
																			},
																		},
																		"dvb_sub_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"pid": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																				},
																			},
																		},
																		"embedded_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"convert_608_to_708": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.EmbeddedConvert608To708_Values(), false),
																					},
																					"source_608_channel_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"source_608_track_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"terminate_captions": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.EmbeddedTerminateCaptions_Values(), false),
																					},
																				},
																			},
																		},
																		"file_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"convert_608_to_708": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.FileSourceConvert608To708_Values(), false),
																					},
																					"framerate": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"framerate_denominator": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1),
																								},
																								"framerate_numerator": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1),
																								},
																							},
																						},
																					},
																					"source_file": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"time_delta": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"source_type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CaptionSourceType_Values(), false),
																		},
																		"teletext_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"page_number": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"track_source_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"track_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
											},
										},
									},
									"crop": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"height": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"width": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"x": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"y": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"deblock_filter": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputDeblockFilter_Values(), false),
									},
									"denoise_filter": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputDenoiseFilter_Values(), false),
									},
									"filter_enable": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputFilterEnable_Values(), false),
									},
									"filter_strength": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntBetween(-5, 5),
										Default:      0,
									},
									"image_inserter": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"insertable_images": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"duration": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"fade_in": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"fade_out": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"height": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"image_inserter_input": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"image_x": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"image_y": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"layer": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"opacity": {
																Type:     schema.TypeInt,
																Optional: true,
																Default:  50,
															},
															"start_time": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"width": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
									"input_clippings": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"end_timecode": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"start_timecode": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
									"input_scan_type": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputScanType_Values(), false),
									},
									"position": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"height": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"width": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(2),
												},
												"x": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"y": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"program_number": {
										Type:         schema.TypeInt,
										Optional:     true,
										ValidateFunc: validation.IntAtLeast(1),
									},
									"psi_control": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputPsiControl_Values(), false),
									},
									"timecode_source": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.InputTimecodeSource_Values(), false),
									},
									"timecode_start": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"video_selector": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"alpha_behavior": {
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice(mediaconvert.AlphaBehavior_Values(), false),
												},
												"color_space": {
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice(mediaconvert.ColorSpace_Values(), false),
												},
												"color_space_usage": {
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice(mediaconvert.ColorSpaceUsage_Values(), false),
												},
												"pid": {
													Type:         schema.TypeInt,
													Optional:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
												"program_number": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"rotate": {
													Type:         schema.TypeString,
													Optional:     true,
													ValidateFunc: validation.StringInSlice(mediaconvert.InputRotate_Values(), false),
												},
												"hdr10_metadata": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"blue_primary_x": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"blue_primary_y": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"green_primary_x": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"green_primary_y": {
																Type:     schema.TypeInt,
																Optional: true},
															"max_content_light_level": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"max_frame_average_light_level": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"max_luminance": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_luminance": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"red_primary_x": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"red_primary_y": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"white_point_x": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"white_point_y": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"motion_image_inserter": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"framerate": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"framerate_denominator": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
												"framerate_numerator": {
													Type:         schema.TypeInt,
													Required:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
											},
										},
									},
									"input": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"insertion_mode": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.MotionImageInsertionMode_Values(), false),
									},
									"offset": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"image_x": {
													Type:     schema.TypeInt,
													Optional: true,
												},
												"image_y": {
													Type:     schema.TypeInt,
													Optional: true,
												},
											},
										},
									},
									"playback": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.MotionImagePlayback_Values(), false),
									},
									"start_time": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"nielsen_configuration": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"breakout_code": {
										Type:     schema.TypeInt,
										Optional: true,
										Default:  0,
									},
									"distributor_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"nielsen_non_linear_watermark": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"active_watermark_process": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.NielsenActiveWatermarkProcessType_Values(), false),
									},
									"adi_filename": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"asset_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"asset_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"cbet_source_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"episode_id": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"metadata_destination": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"source_id": {
										Type:     schema.TypeInt,
										Optional: true,
									},
									"source_watermark_status": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.NielsenSourceWatermarkStatusType_Values(), false),
									},
									"tic_server_url": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"unique_tic_per_audio_track": {
										Type:         schema.TypeString,
										Required:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.NielsenUniqueTicPerAudioTrackType_Values(), false),
									},
								},
							},
						},
						"output_groups": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"automated_encoding_settings": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"abr_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"max_abr_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(100000),
																Default:      8000000,
															},
															"max_renditions": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(3),
																Default:      15,
															},
															"min_abr_bitrate": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(100000),
																Default:      600000,
															},
														},
													},
												},
											},
										},
									},
									"custom_name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"name": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"output_group_settings": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"cmaf_group_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_manifest": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"manifest_name_modifier": {
																			Type:     schema.TypeString,
																			Required: true,
																		},
																		"selected_outputs": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"base_url": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"client_cache": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafClientCache_Values(), false),
																Default:      mediaconvert.CmafClientCacheEnabled,
															},
															"code_specification": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafCodecSpecification_Values(), false),
															},
															"destination": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"s3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"canned_acl": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ObjectCannedAcl_Values(), false),
																								},
																							},
																						},
																					},
																					"encryption": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"encryption_type": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ServerSideEncryptionType_Values(), false),
																								},
																								"kms_key_arn": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															"encryption": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"constant_initialization_vector": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringLenBetween(32, 48),
																		},
																		"encryption_method": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmafEncryptionType_Values(), false),
																		},
																		"initialization_vector_in_manifest": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmafInitializationVectorInManifest_Values(), false),
																		},
																		"speke_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"certificate_arn": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"dash_signaled_system_ids": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																					},
																					"hls_signaled_system_ids": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																					},
																					"resource_id": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"static_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key_format": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"key_format_versions": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"static_key_value": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmafKeyProviderType_Values(), false),
																		},
																	},
																},
															},
															"fragment_length": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"manifest_compression": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafManifestCompression_Values(), false),
															},
															"manifest_duration_format": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafManifestDurationFormat_Values(), false),
															},
															"min_buffer_time": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_final_segment_length": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"mpd_profile": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafMpdProfile_Values(), false),
															},
															"segment_control": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafSegmentControl_Values(), false),
															},
															"segment_length": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"stream_inf_resolution": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafStreamInfResolution_Values(), false),
															},
															"write_dash_manifest": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafWriteDASHManifest_Values(), false),
															},
															"write_hls_manifest": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafWriteHLSManifest_Values(), false),
															},
															"write_segment_timeline_in_representation": {
																Type:         schema.TypeString,
																Required:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.CmafWriteSegmentTimelineInRepresentation_Values(), false),
															},
														},
													},
												},
												"dash_iso_group_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_manifest": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"manifest_name_modifier": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selected_outputs": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"base_url": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"s3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"canned_acl": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ObjectCannedAcl_Values(), false),
																								},
																							},
																						},
																					},
																					"encryption": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"encryption_type": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ServerSideEncryptionType_Values(), false),
																								},
																								"kms_key_arn": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															"encryption": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"playback_device_compatibility": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.DashIsoPlaybackDeviceCompatibility_Values(), false),
																		},
																		"speke_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"certificate_arn": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"resource_id": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"system_ids": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"fragment_length": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"hbbtv_compliance": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.DashIsoHbbtvCompliance_Values(), false),
															},
															"min_buffer_time": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"min_final_segment_length": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"mpd_profile": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.DashIsoMpdProfile_Values(), false),
															},
															"segment_control": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.DashIsoSegmentControl_Values(), false),
															},
															"segment_length": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"write_segment_timeline_in_representation": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.DashIsoWriteSegmentTimelineInRepresentation_Values(), false),
															},
														},
													},
												},
												"file_group_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"s3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"canned_acl": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ObjectCannedAcl_Values(), false),
																								},
																							},
																						},
																					},
																					"encryption": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"encryption_type": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ServerSideEncryptionType_Values(), false),
																								},
																								"kms_key_arn": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
														},
													},
												},
												"hls_group_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"ad_markers": {
																Type:     schema.TypeSet,
																Optional: true,
																Elem:     &schema.Schema{Type: schema.TypeString},
															},
															"additional_manifest": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"manifest_name_modifier": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selected_outputs": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"audio_only_header": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsAudioOnlyHeader_Values(), false),
															},
															"base_url": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"caption_language_mapping": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"caption_channel": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"custom_language_code": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringLenBetween(3, 3),
																		},
																		"language_code": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"language_description": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
															"caption_language_setting": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsCaptionLanguageSetting_Values(), false),
															},
															"client_cache": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsClientCache_Values(), false),
															},
															"codec_specification": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsCodecSpecification_Values(), false),
															},
															"destination": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"s3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"canned_acl": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ObjectCannedAcl_Values(), false),
																								},
																							},
																						},
																					},
																					"encryption": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"encryption_type": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ServerSideEncryptionType_Values(), false),
																								},
																								"kms_key_arn": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															"directory_structure": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsDirectoryStructure_Values(), false),
															},
															"encryption": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"constant_initialization_vector": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringLenBetween(32, 32),
																		},
																		"encryption_method": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsEncryptionType_Values(), false),
																		},
																		"initialization_vector_in_manifest": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsInitializationVectorInManifest_Values(), false),
																		},
																		"offline_encrypted": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsOfflineEncrypted_Values(), false),
																		},
																		"speke_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"certificate_arn": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"resource_id": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"system_ids": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"static_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"key_format": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"key_format_versions": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"static_key_value": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsKeyProviderType_Values(), false),
																		},
																	},
																},
															},
															"manifest_compression": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsManifestCompression_Values(), false),
															},
															"manifest_duration_format": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsManifestDurationFormat_Values(), false),
															},
															"min_final_segment_length": {
																Type:     schema.TypeFloat,
																Optional: true,
																Default:  0,
															},
															"min_segment_length": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"output_selection": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsOutputSelection_Values(), false),
															},
															"program_date_time": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsProgramDateTime_Values(), false),
															},
															"program_date_time_period": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"segment_control": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsSegmentControl_Values(), false),
															},
															"segment_length": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"segments_per_subdirectory": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"stream_inf_resolution": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsStreamInfResolution_Values(), false),
															},
															"timed_metadata_id3_frame": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.HlsTimedMetadataId3Frame_Values(), false),
															},
															"timed_metadata_id3_period": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"timestamp_delta_milliseconds": {
																Type:     schema.TypeInt,
																Optional: true,
															},
														},
													},
												},
												"ms_smooth_group_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"additional_manifest": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"manifest_name_modifier": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"selected_outputs": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeString},
																		},
																	},
																},
															},
															"audio_deduplication": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.MsSmoothAudioDeduplication_Values(), false),
															},
															"destination": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"s3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"access_control": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"canned_acl": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ObjectCannedAcl_Values(), false),
																								},
																							},
																						},
																					},
																					"encryption": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"encryption_type": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.S3ServerSideEncryptionType_Values(), false),
																								},
																								"kms_key_arn": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																	},
																},
															},
															"encryption": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"speke_key_provider": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"certificate_arn": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"resource_id": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																					"system_ids": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																					},
																					"url": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"fragment_length": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"manifest_encoding": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.MsSmoothManifestEncoding_Values(), false),
															},
														},
													},
												},
												"type": {
													Type:         schema.TypeString,
													Required:     true,
													ValidateFunc: validation.StringInSlice(mediaconvert.OutputGroupType_Values(), false),
												},
											},
										},
									},
									"outputs": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_description": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"audio_channel_tagging_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"channel_tag": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioChannelTag_Values(), false),
																		},
																	},
																},
															},
															"audio_normalization_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"algorithm": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioChannelTag_Values(), false),
																		},
																		"algorithm_control": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioNormalizationAlgorithmControl_Values(), false),
																		},
																		"correction_gate_level": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"loudness_logging": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioNormalizationLoudnessLogging_Values(), false),
																		},
																		"peak_calculation": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioNormalizationPeakCalculation_Values(), false),
																		},
																		"target_lkfs": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																	},
																},
															},
															"audio_source_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"audio_type": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntBetween(0, 255),
																Default:      0,
															},
															"audio_type_control": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AudioTypeControl_Values(), false),
															},
															"codec_settings": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"codec": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.AudioCodec_Values(), false),
																		},
																		"aac_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"audio_description_broadcaster_mix": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacAudioDescriptionBroadcasterMix_Values(), false),
																						Default:      nil,
																					},
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(6000),
																						Default:      6000,
																					},
																					"codec_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacCodecProfile_Values(), false),
																						Default:      nil,
																					},
																					"coding_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacCodingMode_Values(), false),
																						Default:      nil,
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacRateControlMode_Values(), false),
																						Default:      nil,
																					},
																					"raw_format": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacRawFormat_Values(), false),
																						Default:      nil,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(8000),
																						Default:      8000,
																					},
																					"specification": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacSpecification_Values(), false),
																						Default:      nil,
																					},
																					"vbr_quality": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AacVbrQuality_Values(), false),
																						Default:      nil,
																					},
																				},
																			},
																		},
																		"ac3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(64000),
																						Default:      64000,
																					},
																					"bitstream_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3BitstreamMode_Values(), false),
																						Default:      nil,
																					},
																					"coding_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3CodingMode_Values(), false),
																						Default:      nil,
																					},
																					"dialnorm": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"dynamic_range_compression_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3DynamicRangeCompressionProfile_Values(), false),
																						Default:      nil,
																					},
																					"lfe_filter": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3LfeFilter_Values(), false),
																						Default:      nil,
																					},
																					"metadata_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3MetadataControl_Values(), false),
																						Default:      nil,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(48000),
																						Default:      48000,
																					},
																				},
																			},
																		},
																		"aiff_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitdepth": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(16),
																						Default:      16,
																					},
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(8000),
																						Default:      8000,
																					},
																				},
																			},
																		},
																		"eac3_atmos_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(384000),
																						Default:      384000,
																					},
																					"bitstream_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosBitstreamMode_Values(), false),
																					},
																					"coding_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosCodingMode_Values(), false),
																					},
																					"dialogue_intelligence": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosDialogueIntelligence_Values(), false),
																					},
																					"dynamic_range_compression_line": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosDynamicRangeCompressionLine_Values(), false),
																					},
																					"dynamic_range_compression_rf": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosDynamicRangeCompressionRf_Values(), false),
																					},
																					"lo_ro_center_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lo_ro_surround_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lt_rt_center_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lt_rt_surround_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"metering_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosMeteringMode_Values(), false),
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(48000),
																						Default:      48000,
																					},
																					"speech_threshold": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"stereo_downmix": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosStereoDownmix_Values(), false),
																					},
																					"surround_ex_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AtmosSurroundExMode_Values(), false),
																					},
																				},
																			},
																		},
																		"eac3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"attenuation_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3AttenuationControl_Values(), false),
																					},
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(64000),
																						Default:      64000,
																					},
																					"bitstream_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3BitstreamMode_Values(), false),
																					},
																					"coding_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3CodingMode_Values(), false),
																					},
																					"dc_filter": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3DcFilter_Values(), false),
																					},
																					"dialnorm": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"dynamic_range_compression_line": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3DynamicRangeCompressionLine_Values(), false),
																					},
																					"dynamic_range_compression_rf": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3DynamicRangeCompressionRf_Values(), false),
																					},
																					"lfe_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3LfeControl_Values(), false),
																					},
																					"lfe_filter": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3LfeFilter_Values(), false),
																					},
																					"lo_ro_center_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lo_ro_surround_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lt_rt_center_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"lt_rt_surround_mix_level": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																						Default:  0,
																					},
																					"metadata_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Ac3MetadataControl_Values(), false),
																					},
																					"passthrough_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3PassthroughControl_Values(), false),
																					},
																					"phase_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3PhaseControl_Values(), false),
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(48000),
																						Default:      48000,
																					},
																					"stereo_downmix": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3StereoDownmix_Values(), false),
																					},
																					"surround_ex_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3SurroundExMode_Values(), false),
																					},
																					"surround_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Eac3SurroundMode_Values(), false),
																					},
																				},
																			},
																		},
																		"mp2_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(32000),
																						Default:      32000,
																					},
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(32000),
																						Default:      32000,
																					},
																				},
																			},
																		},
																		"mp3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(16000),
																						Default:      16000,
																					},
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mp3RateControlMode_Values(), false),
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(22050),
																						Default:      22050,
																					},
																					"vbr_quality": {
																						Type:     schema.TypeInt,
																						Optional: true,
																						Default:  0,
																					},
																				},
																			},
																		},
																		"opus_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(32000),
																						Default:      32000,
																					},
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(16000),
																						Default:      16000,
																					},
																				},
																			},
																		},
																		"vorbis_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"sample_rate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(22050),
																						Default:      22050,
																					},
																					"vbr_quality": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"wav_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitdepth": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(16),
																						Default:      16,
																					},
																					"channels": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																						Default:      1,
																					},
																					"format": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.WavFormat_Values(), false),
																					},
																					"sample_rate": {
																						Optional:     true,
																						Type:         schema.TypeInt,
																						ValidateFunc: validation.IntAtLeast(8000),
																						Default:      8000,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"custom_language_code": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
															},
															"language_code_control": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AudioLanguageCodeControl_Values(), false),
															},
															"remix_settings": {
																Type:     schema.TypeList,
																Optional: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"channel_mapping": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"output_channels": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"input_channels": {
																									Type:     schema.TypeSet,
																									Optional: true,
																									Elem:     &schema.Schema{Type: schema.TypeInt},
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"channels_in": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																		"channels_out": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(1),
																		},
																	},
																},
															},
															"stream_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"caption_description": {
													Type:     schema.TypeList,
													Optional: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"caption_selector_name": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"custom_language_code": {
																Type:     schema.TypeString,
																Optional: true,
															},
															"destination_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"burnin_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"alignment": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleAlignment_Values(), false),
																					},
																					"background_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleBackgroundColor_Values(), false),
																					},
																					"background_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"font_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleFontColor_Values(), false),
																					},
																					"font_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"font_resolution": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(96),
																					},
																					"font_script": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.FontScript_Values(), false),
																					},
																					"font_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"outline_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleOutlineColor_Values(), false),
																					},
																					"outline_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleShadowColor_Values(), false),
																					},
																					"shadow_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_x_offset": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_y_offset": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"teletext_spacing": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.BurninSubtitleTeletextSpacing_Values(), false),
																					},
																					"x_position": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"y_position": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"destination_type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CaptionDestinationType_Values(), false),
																		},
																		"dvb_sub_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"alignment": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleAlignment_Values(), false),
																					},
																					"background_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleBackgroundColor_Values(), false),
																					},
																					"background_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"font_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleFontColor_Values(), false),
																					},
																					"font_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"font_resolution": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(96),
																					},
																					"font_script": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.FontScript_Values(), false),
																					},
																					"font_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"outline_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleOutlineColor_Values(), false),
																					},
																					"outline_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_color": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleShadowColor_Values(), false),
																					},
																					"shadow_opacity": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_x_offset": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"shadow_y_offset": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"subtitling_type": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitlingType_Values(), false),
																					},
																					"teletext_spacing": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DvbSubtitleTeletextSpacing_Values(), false),
																					},
																					"x_position": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"y_position": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																				},
																			},
																		},
																		"embedded_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"destination_608_channel_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"destination_708_service_number": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																				},
																			},
																		},
																		"imsc_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"style_passthrough": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ImscStylePassthrough_Values(), false),
																					},
																				},
																			},
																		},
																		"scc_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"framerate": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.SccDestinationFramerate_Values(), false),
																					},
																				},
																			},
																		},
																		"teletext_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"page_number": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringLenBetween(3, 256),
																					},
																					"page_types": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeString},
																						Set:      schema.HashString,
																					},
																				},
																			},
																		},
																		"ttml_destination_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"style_passthrough": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.TtmlStylePassthrough_Values(), false),
																					},
																				},
																			},
																		},
																	},
																},
															},
															"language_code": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
															},
															"language_description": {
																Type:     schema.TypeString,
																Optional: true,
															},
														},
													},
												},
												"container_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"cmfc_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_duration": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmfcAudioDuration_Values(), false),
																		},
																		"scte35_esam": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmfcScte35Esam_Values(), false),
																		},
																		"scte35_source": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmfcScte35Source_Values(), false),
																		},
																	},
																},
															},
															"container": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.ContainerType_Values(), false),
															},
															"f4v_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"moov_placement": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.F4vMoovPlacement_Values(), false),
																		},
																	},
																},
															},
															"m2ts_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_buffer_model": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsAudioBufferModel_Values(), false),
																		},
																		"audio_duration": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsAudioDuration_Values(), false),
																		},
																		"audio_frames_per_pes": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"audio_pids": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeInt},
																		},
																		"bitrate": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"buffer_model": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsBufferModel_Values(), false),
																		},
																		"dvb_nit_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"network_id": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"network_name": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringLenBetween(1, 256),
																					},
																					"nit_interval": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(25),
																					},
																				},
																			},
																		},
																		"dvb_sdt_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"output_sdt": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.OutputSdt_Values(), false),
																					},
																					"sdt_interval": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(25),
																					},
																					"service_name": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringLenBetween(1, 256),
																					},
																					"service_provider_name": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringLenBetween(1, 256),
																					},
																				},
																			},
																		},
																		"dvb_sub_pids": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeInt},
																		},
																		"dvb_tdt_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"tdt_interval": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																				},
																			},
																		},
																		"dvb_teletext_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																			Default:      499,
																		},
																		"ebp_audio_interval": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsEbpAudioInterval_Values(), false),
																		},
																		"ebp_placement": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsEbpPlacement_Values(), false),
																		},
																		"es_rate_in_pes": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsEsRateInPes_Values(), false),
																		},
																		"force_ts_video_ebp_order": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsForceTsVideoEbpOrder_Values(), false),
																		},
																		"fragment_time": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																		"max_pcr_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"min_ebp_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"nielsen_id3": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsNielsenId3_Values(), false),
																		},
																		"null_packet_bitrate": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																		"pat_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"pcr_control": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsPcrControl_Values(), false),
																		},
																		"pcr_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"pmt_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"pmt_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																			Default:      48,
																		},
																		"private_metadata_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																			Default:      503,
																		},
																		"program_number": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Default:  1,
																		},
																		"rate_mode": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsRateMode_Values(), false),
																		},
																		"scte_35_esam": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"scte_35_esam_pid": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(32),
																					},
																				},
																			},
																		},
																		"scte_35_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"scte_35_source": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsScte35Source_Values(), false),
																		},
																		"segmentation_markers": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsSegmentationMarkers_Values(), false),
																		},
																		"segmentation_style": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M2tsSegmentationStyle_Values(), false),
																		},
																		"segmentation_time": {
																			Type:     schema.TypeFloat,
																			Optional: true,
																		},
																		"timed_metadata_pid": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Default:  32,
																		},
																		"transport_stream_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"video_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																	},
																},
															},
															"m3u8_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_duration": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M3u8AudioDuration_Values(), false)},
																		"audio_frames_per_pes": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"audio_pids": {
																			Type:     schema.TypeSet,
																			Optional: true,
																			Elem:     &schema.Schema{Type: schema.TypeInt},
																		},
																		"nielsen_id3": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M3u8NielsenId3_Values(), false),
																		},
																		"pat_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"pcr_control": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M3u8PcrControl_Values(), false),
																		},
																		"pcr_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"pmt_interval": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"pmt_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"private_metadata_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"program_number": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"scte_35_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"scte_35_source": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.M3u8Scte35Source_Values(), false),
																		},
																		"timed_metadata": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.TimedMetadata_Values(), false),
																		},
																		"timed_metadata_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																		"transport_stream_id": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"video_pid": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(32),
																		},
																	},
																},
															},
															"mov_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"clap_atom": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MovClapAtom_Values(), false),
																		},
																		"cslg_atom": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MovCslgAtom_Values(), false),
																		},
																		"mpeg2_fourcc_control": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MovMpeg2FourCCControl_Values(), false),
																		},
																		"padding_control": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MovPaddingControl_Values(), false),
																		},
																		"reference": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MovReference_Values(), false),
																		},
																	},
																},
															},
															"mp4_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_duration": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.CmfcAudioDuration_Values(), false),
																		},
																		"cslg_atom": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.Mp4CslgAtom_Values(), false),
																		},
																		"ctts_version": {
																			Type:     schema.TypeInt,
																			Optional: true,
																			Default:  0,
																		},
																		"free_space_box": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.Mp4FreeSpaceBox_Values(), false),
																		},
																		"moov_placement": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.Mp4MoovPlacement_Values(), false),
																		},
																		"mp4_major_brand": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
															"mpd_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"accessibility_caption_hints": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MpdAccessibilityCaptionHints_Values(), false),
																		},
																		"audio_duration": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MpdAudioDuration_Values(), false)},
																		"caption_container_type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MpdCaptionContainerType_Values(), false),
																		},
																		"scte_35_esam": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MpdScte35Esam_Values(), false),
																		},
																		"scte_35_source": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MpdScte35Source_Values(), false),
																		},
																	},
																},
															},
															"mxf_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"afd_signaling": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MxfAfdSignaling_Values(), false),
																		},
																		"profile": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.MxfProfile_Values(), false),
																		},
																	},
																},
															},
														},
													},
												},
												"extension": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"name_modifier": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"output_settings": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"hls_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"audio_group_id": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"audio_only_container": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsAudioOnlyContainer_Values(), false),
																		},
																		"audio_rendition_sets": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																		"audio_track_type": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsAudioTrackType_Values(), false),
																		},
																		"iframe_only_manifest": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.HlsIFrameOnlyManifest_Values(), false),
																		},
																		"segment_modifier": {
																			Type:     schema.TypeString,
																			Optional: true,
																		},
																	},
																},
															},
														},
													},
												},
												"preset": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"video_description": {
													Type:     schema.TypeList,
													Optional: true,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"afd_signaling": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AfdSignaling_Values(), false),
															},
															"anti_alias": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.AntiAlias_Values(), false),
															},
															"codec_settings": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"av1_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Av1AdaptiveQuantization_Values(), false),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Av1FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Av1FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"gop_size": {
																						Type:         schema.TypeFloat,
																						Optional:     true,
																						ValidateFunc: validation.FloatAtLeast(0),
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"number_b_frames_between_reference_frames": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntBetween(7, 15),
																					},
																					"qvbr_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"qvbr_quality_level": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1),
																								},
																								"qvbr_quality_level_fine_tune": {
																									Type:     schema.TypeFloat,
																									Optional: true,
																									Default:  0,
																								},
																							},
																						},
																					},
																					"rate_control_mode": {
																						Type:     schema.TypeString,
																						Optional: true,
																						ValidateFunc: validation.StringInSlice([]string{
																							mediaconvert.Av1RateControlModeQvbr,
																						}, false),
																					},
																					"slices": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"spatial_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Av1SpatialAdaptiveQuantization_Values(), false),
																						Default:      mediaconvert.Av1SpatialAdaptiveQuantizationEnabled,
																					},
																				},
																			},
																		},
																		"avc_intra_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"avc_intra_class": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraClass_Values(), false),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraFramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraFramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(24),
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraInterlaceMode_Values(), false),
																					},
																					"slow_pal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraSlowPal_Values(), false),
																						Default:      mediaconvert.AvcIntraSlowPalDisabled,
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.AvcIntraTelecine_Values(), false),
																						Default:      mediaconvert.AvcIntraTelecineNone,
																					},
																				},
																			},
																		},
																		"codec": {
																			Type:         schema.TypeString,
																			Optional:     true,
																			ValidateFunc: validation.StringInSlice(mediaconvert.VideoCodec_Values(), false),
																		},
																		"frame_capture_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"max_captures": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																				},
																			},
																		},
																		"h264_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264AdaptiveQuantization_Values(), false),
																					},
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"codec_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264CodecLevel_Values(), false),
																					},
																					"codec_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264CodecProfile_Values(), false),
																					},
																					"dynamic_sub_gop": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264DynamicSubGop_Values(), false),
																					},
																					"entropy_encoding": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264EntropyEncoding_Values(), false),
																					},
																					"field_encoding": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264FieldEncoding_Values(), false),
																					},
																					"flicker_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264FlickerAdaptiveQuantization_Values(), false),
																						Default:      mediaconvert.H264FlickerAdaptiveQuantizationEnabled,
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"gop_b_reference": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264GopBReference_Values(), false),
																					},
																					"gop_closed_cadence": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"gop_size": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																					},
																					"gop_size_units": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264GopSizeUnits_Values(), false),
																					},
																					"hrd_buffer_initial_fill_percentage": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"hrd_buffer_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264InterlaceMode_Values(), false),
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"min_i_interval": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"number_b_frames_between_reference_frames": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"number_reference_frames": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264ParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality_tuning_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264QualityTuningLevel_Values(), false),
																					},
																					"qvbr_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_average_bitrate": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1000),
																								},
																								"qvbr_quality_level": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntBetween(1, 10),
																								},
																								"qvbr_quality_level_fine_tune": {
																									Type:     schema.TypeFloat,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264RateControlMode_Values(), false),
																					},
																					"repeat_pps": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264RepeatPps_Values(), false),
																					},
																					"scene_change_detect": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264SceneChangeDetect_Values(), false),
																					},
																					"slices": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"slow_pal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264SlowPal_Values(), false),
																						Default:      mediaconvert.H264SlowPalDisabled,
																					},
																					"softness": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntBetween(0, 128),
																					},
																					"spatial_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264SpatialAdaptiveQuantization_Values(), false),
																					},
																					"syntax": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264Syntax_Values(), false),
																						Default:      mediaconvert.H264SyntaxDefault,
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264Telecine_Values(), false),
																						Default:      mediaconvert.Mpeg2TelecineNone,
																					},
																					"temporal_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264TemporalAdaptiveQuantization_Values(), false),
																						Default:      mediaconvert.H264TemporalAdaptiveQuantizationEnabled,
																					},
																					"unregistered_sei_timecode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H264UnregisteredSeiTimecode_Values(), false),
																					},
																				},
																			},
																		},
																		"h265_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265AdaptiveQuantization_Values(), false),
																					},
																					"alternate_transfer_function_sei": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265AlternateTransferFunctionSei_Values(), false),
																					},
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"codec_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265CodecLevel_Values(), false),
																					},
																					"codec_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265CodecProfile_Values(), false),
																					},
																					"dynamic_sub_gop": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265DynamicSubGop_Values(), false),
																					},
																					"flicker_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265FlickerAdaptiveQuantization_Values(), false),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"gop_b_reference": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265GopBReference_Values(), false),
																					},
																					"gop_closed_cadence": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"gop_size": {
																						Type:     schema.TypeFloat,
																						Optional: true,
																					},
																					"gop_size_units": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265GopSizeUnits_Values(), false),
																					},
																					"hrd_buffer_initial_fill_percentage": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"hrd_buffer_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265InterlaceMode_Values(), false),
																						Default:      mediaconvert.H265InterlaceModeProgressive,
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"min_i_interval": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"number_b_frames_between_reference_frames": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"number_reference_frames": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265ParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality_tuning_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265QualityTuningLevel_Values(), false),
																					},
																					"qvbr_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_average_bitrate": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1000),
																								},
																								"qvbr_quality_level": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntBetween(1, 10),
																								},
																								"qvbr_quality_level_fine_tune": {
																									Type:     schema.TypeFloat,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265RateControlMode_Values(), false),
																					},
																					"sample_adaptive_offset_filter_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265SampleAdaptiveOffsetFilterMode_Values(), false),
																					},
																					"scene_change_detect": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265SceneChangeDetect_Values(), false),
																					},
																					"slices": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"slow_pal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265SlowPal_Values(), false),
																						Default:      mediaconvert.H265SlowPalDisabled,
																					},
																					"spatial_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265SpatialAdaptiveQuantization_Values(), false),
																						Default:      mediaconvert.H265SpatialAdaptiveQuantizationEnabled,
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265Telecine_Values(), false),
																					},
																					"temporal_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265TemporalAdaptiveQuantization_Values(), false),
																					},
																					"temporal_ids": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265TemporalIds_Values(), false),
																					},
																					"tiles": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265Tiles_Values(), false),
																					},
																					"unregistered_sei_timecode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265UnregisteredSeiTimecode_Values(), false),
																					},
																					"write_mp4_packaging_type": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.H265WriteMp4PackagingType_Values(), false),
																					},
																				},
																			},
																		},
																		"mpeg2_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2AdaptiveQuantization_Values(), false),
																					},
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"codec_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2CodecLevel_Values(), false),
																					},
																					"codec_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2CodecProfile_Values(), false),
																					},
																					"dynamic_sub_gop": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2DynamicSubGop_Values(), false),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(24),
																					},
																					"gop_closed_cadence": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"gop_size": {
																						Type:         schema.TypeFloat,
																						Optional:     true,
																						ValidateFunc: validation.FloatAtLeast(0),
																					},
																					"gop_size_units": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2GopSizeUnits_Values(), false),
																					},
																					"hrd_buffer_initial_fill_percentage": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"hrd_buffer_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2InterlaceMode_Values(), false),
																						Default:      mediaconvert.Mpeg2InterlaceModeProgressive,
																					},
																					"intra_dc_precision": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2IntraDcPrecision_Values(), false),
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"min_i_interval": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"number_b_frames_between_reference_frames": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2ParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality_tuning_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2QualityTuningLevel_Values(), false),
																						Default:      mediaconvert.Mpeg2QualityTuningLevelSinglePass,
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2RateControlMode_Values(), false),
																					},
																					"scene_change_detect": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2SceneChangeDetect_Values(), false),
																					},
																					"slowpal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2SlowPal_Values(), false),
																						Default:      mediaconvert.Mpeg2SlowPalDisabled,
																					},
																					"softness": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntBetween(17, 128),
																					},
																					"spatial_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2SpatialAdaptiveQuantization_Values(), false),
																					},
																					"syntax": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2Syntax_Values(), false),
																						Default:      mediaconvert.Mpeg2SyntaxDefault,
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2Telecine_Values(), false),
																						Default:      mediaconvert.Mpeg2TelecineNone,
																					},
																					"temporal_adaptive_quantization": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Mpeg2TemporalAdaptiveQuantization_Values(), false),
																						Default:      mediaconvert.Mpeg2TemporalAdaptiveQuantizationEnabled,
																					},
																				},
																			},
																		},
																		"prores_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"codec_profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresCodecProfile_Values(), false),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresFramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresFramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresInterlaceMode_Values(), false),
																						Default:      mediaconvert.ProresInterlaceModeProgressive,
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"slow_pal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresSlowPal_Values(), false),
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ProresTelecine_Values(), false),
																						Default:      mediaconvert.ProresTelecineNone,
																					},
																				},
																			},
																		},
																		"vc3_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(24),
																					},
																					"interlace_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3InterlaceMode_Values(), false),
																					},
																					"slowpal": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3SlowPal_Values(), false),
																					},
																					"telecine": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3Telecine_Values(), false),
																					},
																					"vc3_class": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vc3Class_Values(), false),
																					},
																				},
																			},
																		},
																		"vp8_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp8FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp8FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"gop_size": {
																						Type:         schema.TypeFloat,
																						Optional:     true,
																						ValidateFunc: validation.FloatAtLeast(0),
																					},
																					"hrd_buffer_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp8ParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality_tuning_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp8QualityTuningLevel_Values(), false),
																						Default:      mediaconvert.Vp8QualityTuningLevelMultiPass,
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp8RateControlMode_Values(), false),
																					},
																				},
																			},
																		},
																		"vp9_settings": {
																			Type:     schema.TypeList,
																			Optional: true,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"framerate_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp9FramerateControl_Values(), false),
																					},
																					"framerate_conversion_algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp9FramerateConversionAlgorithm_Values(), false),
																					},
																					"framerate_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"framerate_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"gop_size": {
																						Type:         schema.TypeFloat,
																						Optional:     true,
																						ValidateFunc: validation.FloatAtLeast(0),
																					},
																					"hrd_buffer_size": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"max_bitrate": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1000),
																					},
																					"par_control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp9ParControl_Values(), false),
																					},
																					"par_denominator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"par_numerator": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"quality_tuning_level": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp9QualityTuningLevel_Values(), false),
																						Default:      mediaconvert.Vp9QualityTuningLevelMultiPass,
																					},
																					"rate_control_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.Vp9RateControlMode_Values(), false),
																					},
																				},
																			},
																		},
																	},
																},
															},
															"color_metadata": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.ColorMetadata_Values(), false),
															},
															"crop": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"height": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(2),
																		},
																		"width": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(2),
																		},
																		"x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"drop_frame_timecode": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.DropFrameTimecode_Values(), false),
															},
															"fixed_afd": {
																Type:     schema.TypeInt,
																Optional: true,
															},
															"height": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32),
															},
															"position": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"height": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(2),
																		},
																		"width": {
																			Type:         schema.TypeInt,
																			Optional:     true,
																			ValidateFunc: validation.IntAtLeast(2),
																		},
																		"x": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																		"y": {
																			Type:     schema.TypeInt,
																			Optional: true,
																		},
																	},
																},
															},
															"respond_to_afd": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.RespondToAfd_Values(), false),
															},
															"scaling_behavior": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.ScalingBehavior_Values(), false),
															},
															"sharpness": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(0),
																Default:      0,
															},
															"timecode_insertion": {
																Type:         schema.TypeString,
																Optional:     true,
																ValidateFunc: validation.StringInSlice(mediaconvert.VideoTimecodeInsertion_Values(), false),
															},
															"video_preprocessors": {
																Type:     schema.TypeList,
																Optional: true,
																MaxItems: 1,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"color_corrector": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"brightness": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"color_space_conversion": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.ColorSpaceConversion_Values(), false),
																					},
																					"contrast": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																					"hdr10_metadata": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"blue_primary_x": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"blue_primary_y": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"green_primary_x": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"green_primary_y": {
																									Type:     schema.TypeInt,
																									Optional: true},
																								"max_content_light_level": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"max_frame_average_light_level": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"max_luminance": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"min_luminance": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"red_primary_x": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"red_primary_y": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"white_point_x": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"white_point_y": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"hue": {
																						Type:     schema.TypeInt,
																						Optional: true,
																					},
																					"saturation": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(1),
																					},
																				},
																			},
																		},
																		"deinterlacer": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"algorithm": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DeinterlaceAlgorithm_Values(), false),
																					},
																					"control": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DeinterlacerControl_Values(), false),
																					},
																					"mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DeinterlacerMode_Values(), false),
																					},
																				},
																			},
																		},
																		"dolby_vision": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"l6_metadata": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"max_cll": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"max_fall": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"l6_mode": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DolbyVisionLevel6Mode_Values(), false),
																					},
																					"profile": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.DolbyVisionProfile_Values(), false),
																					},
																				},
																			},
																		},
																		"image_inserter": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"insertable_image": {
																						Type:     schema.TypeList,
																						Optional: true,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"duration": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"fade_in": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"fade_out": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"height": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"image_inserter_input": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringLenBetween(14, 4000),
																								},
																								"image_x": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"image_y": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"layer": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"opacity": {
																									Type:     schema.TypeInt,
																									Optional: true,
																									Default:  50,
																								},
																								"start_time": {
																									Type:     schema.TypeString,
																									Optional: true,
																								},
																								"width": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"noise_reducer": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"filter": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.NoiseReducerFilter_Values(), false),
																					},
																					"filter_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"strength": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"spatial_filter_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"post_filter_sharpen_strength": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"speed": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"strength": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																					"temporal_filter_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"aggressive_mode": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"post_temporal_sharpening": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.NoiseFilterPostTemporalSharpening_Values(), false),
																									Default:      mediaconvert.NoiseFilterPostTemporalSharpeningAuto,
																								},
																								"speed": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																								"strength": {
																									Type:     schema.TypeInt,
																									Optional: true,
																								},
																							},
																						},
																					},
																				},
																			},
																		},
																		"partner_watermarking": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"nexguard_file_marker_settings": {
																						Type:     schema.TypeList,
																						Optional: true,
																						MaxItems: 1,
																						Elem: &schema.Resource{
																							Schema: map[string]*schema.Schema{
																								"license": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringIsNotEmpty,
																								},
																								"payload": {
																									Type:         schema.TypeInt,
																									Optional:     true,
																									ValidateFunc: validation.IntAtLeast(1),
																								},
																								"preset": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringIsNotEmpty,
																								},
																								"strength": {
																									Type:         schema.TypeString,
																									Optional:     true,
																									ValidateFunc: validation.StringInSlice(mediaconvert.WatermarkingStrength_Values(), false),
																									Default:      mediaconvert.WatermarkingStrengthDefault,
																								},
																							},
																						},
																					},
																				},
																			}},
																		"timecode_burnin": {
																			Type:     schema.TypeList,
																			Optional: true,
																			MaxItems: 1,
																			Elem: &schema.Resource{
																				Schema: map[string]*schema.Schema{
																					"font_size": {
																						Type:         schema.TypeInt,
																						Optional:     true,
																						ValidateFunc: validation.IntAtLeast(10),
																					},
																					"position": {
																						Type:         schema.TypeString,
																						Optional:     true,
																						ValidateFunc: validation.StringInSlice(mediaconvert.TimecodeBurninPosition_Values(), false),
																					},
																					"prefix": {
																						Type:     schema.TypeString,
																						Optional: true,
																					},
																				},
																			},
																		},
																	},
																},
															},
															"Width ": {
																Type:         schema.TypeInt,
																Optional:     true,
																ValidateFunc: validation.IntAtLeast(32),
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"timecode_config": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"anchor": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"source": {
										Type:         schema.TypeString,
										Optional:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.TimecodeSource_Values(), false),
									},
									"start": {
										Type:     schema.TypeString,
										Optional: true,
									},
									"timestamp_offset": {
										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},
						"timed_metadata_insertion": {
							Type:     schema.TypeList,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id3_insertion": {
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id3": {
													Type:     schema.TypeString,
													Optional: true,
												},
												"timecode": {
													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAwsMediaConvertJobTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func resourceAwsMediaConvertJobTemplateRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsMediaConvertJobTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func resourceAwsMediaConvertJobTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}
