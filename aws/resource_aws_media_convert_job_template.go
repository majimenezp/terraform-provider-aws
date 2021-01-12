package aws

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/terraform-providers/terraform-provider-aws/aws/internal/keyvaluetags"
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
						"wait_minutes": {
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
						"input": {
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// needed to change to a list of types to be able to parse it
									"audio_selector_group": {
										Type:     schema.TypeList,
										Optional: true,
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
									"audio_selector": {
										Type:     schema.TypeList,
										Optional: true,
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
																					"input_channels_fine_tune": {
																						Type:     schema.TypeSet,
																						Optional: true,
																						Elem:     &schema.Schema{Type: schema.TypeFloat},
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
									"caption_selector": {
										Type:     schema.TypeList,
										Optional: true,
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
										Default:      1,
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
													Default:      1,
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
						"output_group": {
							Type:     schema.TypeList,
							Required: true,
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
									"output": {
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
																						Required:     true,
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
															"width": {
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

func resourceAwsMediaConvertJobTemplateRead(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig
	getOpts := &mediaconvert.GetJobTemplateInput{
		Name: aws.String(d.Id()),
	}
	resp, err := conn.GetJobTemplate(getOpts)
	if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
		log.Printf("[WARN] Media Convert Job Template (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Job Template: %s", err)
	}
	d.Set("arn", resp.JobTemplate.Arn)
	d.Set("category", resp.JobTemplate.Category)
	d.Set("name", resp.JobTemplate.Name)
	d.Set("description", resp.JobTemplate.Description)
	d.Set("priority", resp.JobTemplate.Priority)
	d.Set("queue", resp.JobTemplate.Queue)
	d.Set("status_update_interval", resp.JobTemplate.StatusUpdateInterval)

	if err := d.Set("acceleration_settings", flattenMediaConvertAccelerationSettings(resp.JobTemplate.AccelerationSettings)); err != nil {
		return fmt.Errorf("Error setting Media Convert Job template AccelerationSettings: %s", err)
	}
	if err := d.Set("hop_destinations", flattenMediaConvertHopDestinations(resp.JobTemplate.HopDestinations)); err != nil {
		return fmt.Errorf("Error setting Media Convert Job template AccelerationSettings: %s", err)
	}

	if err := d.Set("settings", flattenMediaConvertJobTemplateSettings(resp.JobTemplate.Settings)); err != nil {
		return fmt.Errorf("Error setting Media Convert Job template Settings: %s", err)
	}

	tags, err := keyvaluetags.MediaconvertListTags(conn, aws.StringValue(resp.JobTemplate.Arn))
	if err != nil {
		return fmt.Errorf("error listing tags for Media Convert Preset (%s): %s", d.Id(), err)
	}
	if err := d.Set("tags", tags.IgnoreAws().IgnoreConfig(ignoreTagsConfig).Map()); err != nil {
		return fmt.Errorf("error setting tags: %s", err)
	}
	return nil
}

func resourceAwsMediaConvertJobTemplateUpdate(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	log.Printf("[INFO] Updating MediaConvert Job Template: %s", d.Get("name").(string))
	if d.HasChanges("description", "priority", "queue", "status_update_interval", "acceleration_settings", "hop_destinations", "settings") {
		updateOpts := &mediaconvert.UpdateJobTemplateInput{
			Name: aws.String(d.Id()),
		}
		if v, ok := d.GetOk("description"); ok {
			updateOpts.Description = aws.String(v.(string))
		}
		if v, ok := d.GetOk("priority"); ok {
			updateOpts.Description = aws.String(v.(string))
		}
		if v, ok := d.GetOk("queue"); ok {
			updateOpts.Queue = aws.String(v.(string))
		}
		if v, ok := d.GetOk("status_update_interval"); ok {
			updateOpts.StatusUpdateInterval = aws.String(v.(string))
		}
		if v, ok := d.GetOk("acceleration_settings"); ok {
			accelerationSettings := v.([]interface{})
			updateOpts.AccelerationSettings = expandMediaConvertJobTemplateAccelerationSettings(accelerationSettings)
		}
		if v, ok := d.GetOk("hop_destinations"); ok {
			hopDestinations := v.([]interface{})
			updateOpts.HopDestinations = expandMediaConvertJobTemplateHopDestinations(hopDestinations)
		}
		if v, ok := d.GetOk("settings"); ok {
			settings := v.([]interface{})
			updateOpts.Settings = expandMediaConvertJobTemplateSettings(settings)
		}
		_, err = conn.UpdateJobTemplate(updateOpts)
		if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
			log.Printf("[WARN] Media Convert Job Template (%s) not found, removing from state", d.Id())
			d.SetId("")
			return nil
		}
		if err != nil {
			return fmt.Errorf("Error updating Media Convert Job template: %s", err)
		}
	}
	if d.HasChange("tags") {
		o, n := d.GetChange("tags")
		if err := keyvaluetags.MediaconvertUpdateTags(conn, d.Get("arn").(string), o, n); err != nil {
			return fmt.Errorf("error updating tags: %s", err)
		}
	}
	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func resourceAwsMediaConvertJobTemplateDelete(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceAwsMediaConvertJobTemplateCreate(d *schema.ResourceData, meta interface{}) error {
	conn, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	input := &mediaconvert.CreateJobTemplateInput{
		Name:                 aws.String(d.Get("name").(string)),
		Description:          aws.String(d.Get("description").(string)),
		Priority:             aws.Int64(int64(d.Get("priority").(int))),
		StatusUpdateInterval: aws.String(d.Get("status_update_interval").(string)),
		Tags:                 keyvaluetags.New(d.Get("tags").(map[string]interface{})).IgnoreAws().MediaconvertTags(),
	}
	if attr, ok := d.GetOk("settings"); ok {
		input.Settings = expandMediaConvertJobTemplateSettings(attr.([]interface{}))
	}
	if attr, ok := d.GetOk("acceleration_settings"); ok {
		input.AccelerationSettings = expandMediaConvertJobTemplateAccelerationSettings(attr.([]interface{}))
	}
	if attr, ok := d.GetOk("hop_destinations"); ok {
		input.HopDestinations = expandMediaConvertJobTemplateHopDestinations(attr.([]interface{}))
	}
	if v, ok := d.GetOk("category"); ok && v.(string) != "" {
		input.Category = aws.String(v.(string))
	}
	if v, ok := d.GetOk("queue"); ok && v.(string) != "" {
		input.Queue = aws.String(v.(string))
	}

	resp, err := conn.CreateJobTemplate(input)
	if err != nil {
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		strA, _ := json.Marshal(resp)
		fmt.Println(string(strA))
		strB, _ := json.Marshal(input)
		fmt.Println(string(strB))
		fmt.Println("+++++++++++++++++++++++++++++++++++++++++")
		return fmt.Errorf("Error creating Media Convert Job Template: %s", err)
	}
	d.SetId(aws.StringValue(resp.JobTemplate.Name))

	return resourceAwsMediaConvertJobTemplateRead(d, meta)
}

func expandMediaConvertJobTemplateSettings(list []interface{}) *mediaconvert.JobTemplateSettings {

	result := &mediaconvert.JobTemplateSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["ad_avail_offset"].(int); ok {
		result.AdAvailOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["avail_blanking"]; ok {
		result.AvailBlanking = expandMediaConvertAvailBlanking(v.([]interface{}))
	}
	if v, ok := tfMap["esam"]; ok {
		result.Esam = expandMediaConvertEsamSettings(v.([]interface{}))
	}
	if v, ok := tfMap["input"]; ok {
		result.Inputs = expandMediaConvertInputTemplate(v.([]interface{}))
	}
	if v, ok := tfMap["motion_image_inserter"]; ok {
		result.MotionImageInserter = expandMediaConvertMotionImageInserter(v.([]interface{}))
	}
	if v, ok := tfMap["nielsen_configuration"]; ok {
		result.NielsenConfiguration = expandMediaConvertMotionNielsenConfiguration(v.([]interface{}))
	}
	if v, ok := tfMap["nielsen_non_linear_watermark"]; ok {
		result.NielsenNonLinearWatermark = expandMediaConvertNielsenNonLinearWatermarkSettings(v.([]interface{}))
	}
	if v, ok := tfMap["output_group"]; ok {
		result.OutputGroups = expandMediaConvertOutputGroup(v.([]interface{}))
	}
	if v, ok := tfMap["timecode_config"]; ok {
		result.TimecodeConfig = expandMediaConvertTimecodeConfig(v.([]interface{}))
	}
	if v, ok := tfMap["timed_metadata_insertion"]; ok {
		result.TimedMetadataInsertion = expandMediaConvertTimedMetadataInsertion(v.([]interface{}))
	}
	fmt.Println(tfMap)
	return result
}

func expandMediaConvertTimedMetadataInsertion(list []interface{}) *mediaconvert.TimedMetadataInsertion {
	result := &mediaconvert.TimedMetadataInsertion{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["id3_insertion"]; ok {
		result.Id3Insertions = expandMediaConvertId3Insertion(v.([]interface{}))
	}
	return result
}

func expandMediaConvertId3Insertion(list []interface{}) []*mediaconvert.Id3Insertion {
	results := []*mediaconvert.Id3Insertion{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.Id3Insertion{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["id3"].(string); ok && v != "" {
			result.Id3 = aws.String(v)
		}
		if v, ok := tfMap["timecode"].(string); ok && v != "" {
			result.Timecode = aws.String(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertTimecodeConfig(list []interface{}) *mediaconvert.TimecodeConfig {
	result := &mediaconvert.TimecodeConfig{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["anchor"].(string); ok && v != "" {
		result.Anchor = aws.String(v)
	}
	if v, ok := tfMap["source"].(string); ok && v != "" {
		result.Source = aws.String(v)
	}
	if v, ok := tfMap["start"].(string); ok && v != "" {
		result.Start = aws.String(v)
	}
	if v, ok := tfMap["timestamp_offset"].(string); ok && v != "" {
		result.TimestampOffset = aws.String(v)
	}
	return result
}

func expandMediaConvertOutputGroup(list []interface{}) []*mediaconvert.OutputGroup {
	results := []*mediaconvert.OutputGroup{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.OutputGroup{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["automated_encoding_settings"]; ok {
			result.AutomatedEncodingSettings = expandMediaConvertMotionAutomatedEncodingSettings(v.([]interface{}))
		}
		if v, ok := tfMap["custom_name"].(string); ok && v != "" {
			result.CustomName = aws.String(v)
		}
		if v, ok := tfMap["name"].(string); ok && v != "" {
			result.Name = aws.String(v)
		}
		if v, ok := tfMap["output_group_settings"]; ok {
			result.OutputGroupSettings = expandMediaConvertMotionOutputGroupSettings(v.([]interface{}))
		}
		if v, ok := tfMap["output"]; ok {
			result.Outputs = expandMediaConvertOutput(v.([]interface{}))
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertOutput(list []interface{}) []*mediaconvert.Output {
	results := []*mediaconvert.Output{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.Output{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["audio_description"]; ok {
			result.AudioDescriptions = expandMediaConvertAudioDescription(v.([]interface{}))
		}
		if v, ok := tfMap["caption_description"]; ok {
			result.CaptionDescriptions = expandMediaConvertCaptionDescription(v.([]interface{}))
		}
		if v, ok := tfMap["container_settings"]; ok {
			result.ContainerSettings = expandMediaConvertContainerSettings(v.([]interface{}))
		}
		if v, ok := tfMap["extension"].(string); ok && v != "" {
			result.Extension = aws.String(v)
		}
		if v, ok := tfMap["name_modifier"].(string); ok && v != "" {
			result.NameModifier = aws.String(v)
		}
		if v, ok := tfMap["output_settings"]; ok {
			result.OutputSettings = expandMediaConvertOutputSettings(v.([]interface{}))
		}
		if v, ok := tfMap["preset"].(string); ok && v != "" {
			result.Preset = aws.String(v)
		}
		if v, ok := tfMap["video_description"]; ok {
			result.VideoDescription = expandMediaConvertVideoDescription(v.([]interface{}))
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertOutputSettings(list []interface{}) *mediaconvert.OutputSettings {
	result := &mediaconvert.OutputSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["hls_settings"]; ok {
		result.HlsSettings = expandMediaConvertHlsSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertHlsSettings(list []interface{}) *mediaconvert.HlsSettings {
	result := &mediaconvert.HlsSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_group_id"].(string); ok && v != "" {
		result.AudioGroupId = aws.String(v)
	}
	if v, ok := tfMap["audio_only_container"].(string); ok && v != "" {
		result.AudioOnlyContainer = aws.String(v)
	}
	if v, ok := tfMap["audio_rendition_sets"].(string); ok && v != "" {
		result.AudioRenditionSets = aws.String(v)
	}
	if v, ok := tfMap["audio_track_type"].(string); ok && v != "" {
		result.AudioTrackType = aws.String(v)
	}
	if v, ok := tfMap["iframe_only_manifest"].(string); ok && v != "" {
		result.IFrameOnlyManifest = aws.String(v)
	}
	if v, ok := tfMap["segment_modifier"].(string); ok && v != "" {
		result.SegmentModifier = aws.String(v)
	}
	return result
}

func expandMediaConvertAudioDescription(list []interface{}) []*mediaconvert.AudioDescription {
	results := []*mediaconvert.AudioDescription{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.AudioDescription{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["audio_channel_tagging_settings"]; ok {
			result.AudioChannelTaggingSettings = expandMediaConvertAudioChannelTagging(v.([]interface{}))
		}
		if v, ok := tfMap["audio_normalization_settings"]; ok {
			result.AudioNormalizationSettings = expandMediaConvertAudioNormalizationSettings(v.([]interface{}))
		}
		if v, ok := tfMap["audio_source_name"].(string); ok && v != "" {
			result.AudioSourceName = aws.String(v)
		}
		if v, ok := tfMap["audio_type"].(int); ok && v != 0 {
			result.AudioType = aws.Int64(int64(v))
		}
		if v, ok := tfMap["audio_type_control"].(string); ok && v != "" {
			result.AudioTypeControl = aws.String(v)
		}
		if v, ok := tfMap["codec_settings"]; ok {
			result.CodecSettings = expandMediaConvertCodecSettings(v.([]interface{}))
		}
		if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
			result.CustomLanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_code"].(string); ok && v != "" {
			result.LanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_code_control"].(string); ok && v != "" {
			result.LanguageCodeControl = aws.String(v)
		}
		if v, ok := tfMap["remix_settings"]; ok {
			result.RemixSettings = expandMediaConvertRemixSettings(v.([]interface{}))
		}
		if v, ok := tfMap["stream_name"].(string); ok && v != "" {
			result.StreamName = aws.String(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertCaptionDescription(list []interface{}) []*mediaconvert.CaptionDescription {
	results := []*mediaconvert.CaptionDescription{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.CaptionDescription{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["caption_selector_name"].(string); ok && v != "" {
			result.CaptionSelectorName = aws.String(v)
		}
		if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
			result.CustomLanguageCode = aws.String(v)
		}
		if v, ok := tfMap["destination_settings"]; ok {
			result.DestinationSettings = expandMediaConvertCaptionDestinationSettings(v.([]interface{}))
		}
		if v, ok := tfMap["language_code"].(string); ok && v != "" {
			result.LanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_description"].(string); ok && v != "" {
			result.LanguageDescription = aws.String(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertCaptionDestinationSettings(list []interface{}) *mediaconvert.CaptionDestinationSettings {
	result := &mediaconvert.CaptionDestinationSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["burnin_destination_settings"]; ok {
		result.BurninDestinationSettings = expandMediaConvertBurninDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["destination_type"].(string); ok && v != "" {
		result.DestinationType = aws.String(v)
	}
	if v, ok := tfMap["dvb_sub_destination_settings"]; ok {
		result.DvbSubDestinationSettings = expandMediaConvertDvbSubDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["embedded_destination_settings "]; ok {
		result.EmbeddedDestinationSettings = expandMediaConvertEmbeddedDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["imsc_destination_settings "]; ok {
		result.ImscDestinationSettings = expandMediaConvertImscDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["scc_destination_settings "]; ok {
		result.SccDestinationSettings = expandMediaConvertSccDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["teletext_destination_settings "]; ok {
		result.TeletextDestinationSettings = expandMediaConvertTeletextDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["ttml_destination_settings "]; ok {
		result.TtmlDestinationSettings = expandMediaConvertTtmlDestinationSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertMotionOutputGroupSettings(list []interface{}) *mediaconvert.OutputGroupSettings {
	result := &mediaconvert.OutputGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["cmaf_group_settings"]; ok {
		result.CmafGroupSettings = expandMediaConvertCmafGroupSettings(v.([]interface{}))
	}
	if v, ok := tfMap["dash_iso_group_settings"]; ok {
		result.DashIsoGroupSettings = expandMediaConvertDashIsoGroupSettings(v.([]interface{}))
	}
	if v, ok := tfMap["file_group_settings"]; ok {
		result.FileGroupSettings = expandMediaConvertFileGroupSettings(v.([]interface{}))
	}
	if v, ok := tfMap["hls_group_settings"]; ok {
		result.HlsGroupSettings = expandMediaConvertHlsGroupSettings(v.([]interface{}))
	}
	if v, ok := tfMap["ms_smooth_group_settings"]; ok {
		result.MsSmoothGroupSettings = expandMediaConvertMsSmoothGroupSettings(v.([]interface{}))
	}
	if v, ok := tfMap["type"].(string); ok && v != "" {
		result.Type = aws.String(v)
	}
	return result
}

func expandMediaConvertCmafGroupSettings(list []interface{}) *mediaconvert.CmafGroupSettings {
	result := &mediaconvert.CmafGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["additional_manifest"]; ok {
		result.AdditionalManifests = expandMediaConvertCmafAdditionalManifest(v.([]interface{}))
	}
	if v, ok := tfMap["base_url"].(string); ok && v != "" {
		result.BaseUrl = aws.String(v)
	}
	if v, ok := tfMap["client_cache"].(string); ok && v != "" {
		result.ClientCache = aws.String(v)
	}
	if v, ok := tfMap["code_specification"].(string); ok && v != "" {
		result.CodecSpecification = aws.String(v)
	}
	if v, ok := tfMap["destination"].(string); ok && v != "" {
		result.Destination = aws.String(v)
	}
	if v, ok := tfMap["destination_settings"]; ok {
		result.DestinationSettings = expandMediaConvertDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["encryption"]; ok {
		result.Encryption = expandMediaConvertCmafEncryptionSettings(v.([]interface{}))
	}
	if v, ok := tfMap["fragment_length"].(int); ok {
		result.FragmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["manifest_compression"].(string); ok && v != "" {
		result.ManifestCompression = aws.String(v)
	}
	if v, ok := tfMap["manifest_duration_format"].(string); ok && v != "" {
		result.ManifestDurationFormat = aws.String(v)
	}
	if v, ok := tfMap["min_buffer_time"].(int); ok {
		result.MinBufferTime = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_final_segment_length"].(float64); ok {
		result.MinFinalSegmentLength = aws.Float64(float64(v))
	}
	if v, ok := tfMap["mpd_profile"].(string); ok && v != "" {
		result.MpdProfile = aws.String(v)
	}
	if v, ok := tfMap["segment_control"].(string); ok && v != "" {
		result.SegmentControl = aws.String(v)
	}
	if v, ok := tfMap["segment_length"].(int); ok {
		result.SegmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stream_inf_resolution"].(string); ok && v != "" {
		result.StreamInfResolution = aws.String(v)
	}
	if v, ok := tfMap["write_dash_manifest"].(string); ok && v != "" {
		result.WriteDashManifest = aws.String(v)
	}
	if v, ok := tfMap["write_hls_manifest"].(string); ok && v != "" {
		result.WriteHlsManifest = aws.String(v)
	}
	if v, ok := tfMap["write_segment_timeline_in_representation"].(string); ok && v != "" {
		result.WriteSegmentTimelineInRepresentation = aws.String(v)
	}
	return result
}

func expandMediaConvertCmafAdditionalManifest(list []interface{}) []*mediaconvert.CmafAdditionalManifest {
	results := []*mediaconvert.CmafAdditionalManifest{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.CmafAdditionalManifest{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["manifest_name_modifier"].(string); ok && v != "" {
			result.ManifestNameModifier = aws.String(v)
		}
		if v, ok := tfMap["selected_outputs"].(*schema.Set); ok && v.Len() > 0 {
			result.SelectedOutputs = expandStringSet(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertCmafEncryptionSettings(list []interface{}) *mediaconvert.CmafEncryptionSettings {
	result := &mediaconvert.CmafEncryptionSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["constant_initialization_vector"].(string); ok && v != "" {
		result.ConstantInitializationVector = aws.String(v)
	}
	if v, ok := tfMap["encryption_method"].(string); ok && v != "" {
		result.EncryptionMethod = aws.String(v)
	}
	if v, ok := tfMap["initialization_vector_in_manifest"].(string); ok && v != "" {
		result.InitializationVectorInManifest = aws.String(v)
	}
	if v, ok := tfMap["speke_key_provider"]; ok {
		result.SpekeKeyProvider = expandMediaConvertSpekeKeyProviderCmaf(v.([]interface{}))
	}
	if v, ok := tfMap["static_key_provider"]; ok {
		result.StaticKeyProvider = expandMediaConvertStaticKeyProvider(v.([]interface{}))
	}
	if v, ok := tfMap["type"].(string); ok && v != "" {
		result.Type = aws.String(v)
	}
	return result
}

func expandMediaConvertDashIsoGroupSettings(list []interface{}) *mediaconvert.DashIsoGroupSettings {
	result := &mediaconvert.DashIsoGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["additional_manifest"]; ok {
		result.AdditionalManifests = expandMediaConvertDashAdditionalManifest(v.([]interface{}))
	}
	if v, ok := tfMap["base_url"].(string); ok && v != "" {
		result.BaseUrl = aws.String(v)
	}
	if v, ok := tfMap["destination"].(string); ok && v != "" {
		result.Destination = aws.String(v)
	}
	if v, ok := tfMap["destination_settings"]; ok {
		result.DestinationSettings = expandMediaConvertDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["encryption"]; ok {
		result.Encryption = expandMediaConvertDashIsoEncryptionSettings(v.([]interface{}))
	}
	if v, ok := tfMap["fragment_length"].(int); ok {
		result.FragmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hbbtv_compliance"].(string); ok && v != "" {
		result.HbbtvCompliance = aws.String(v)
	}
	if v, ok := tfMap["min_buffer_time"].(int); ok {
		result.MinBufferTime = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_final_segment_length"].(float64); ok {
		result.MinFinalSegmentLength = aws.Float64(float64(v))
	}
	if v, ok := tfMap["mpd_profile"].(string); ok && v != "" {
		result.MpdProfile = aws.String(v)
	}
	if v, ok := tfMap["segment_control"].(string); ok && v != "" {
		result.SegmentControl = aws.String(v)
	}
	if v, ok := tfMap["segment_length"].(int); ok {
		result.SegmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["write_segment_timeline_in_representation"].(string); ok && v != "" {
		result.WriteSegmentTimelineInRepresentation = aws.String(v)
	}
	return result
}

func expandMediaConvertDashIsoEncryptionSettings(list []interface{}) *mediaconvert.DashIsoEncryptionSettings {
	result := &mediaconvert.DashIsoEncryptionSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["playback_device_compatibility"].(string); ok && v != "" {
		result.PlaybackDeviceCompatibility = aws.String(v)
	}
	if v, ok := tfMap["speke_key_provider"]; ok {
		result.SpekeKeyProvider = expandMediaConvertSpekeKeyProvider(v.([]interface{}))
	}
	return result
}

func expandMediaConvertDashAdditionalManifest(list []interface{}) []*mediaconvert.DashAdditionalManifest {
	results := []*mediaconvert.DashAdditionalManifest{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.DashAdditionalManifest{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["manifest_name_modifier"].(string); ok && v != "" {
			result.ManifestNameModifier = aws.String(v)
		}
		if v, ok := tfMap["selected_outputs"].(*schema.Set); ok && v.Len() > 0 {
			result.SelectedOutputs = expandStringSet(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertFileGroupSettings(list []interface{}) *mediaconvert.FileGroupSettings {
	result := &mediaconvert.FileGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["destination"].(string); ok && v != "" {
		result.Destination = aws.String(v)
	}
	if v, ok := tfMap["destination_settings"]; ok {
		result.DestinationSettings = expandMediaConvertDestinationSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertHlsGroupSettings(list []interface{}) *mediaconvert.HlsGroupSettings {
	result := &mediaconvert.HlsGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["ad_markers"].(*schema.Set); ok && v.Len() > 0 {
		result.AdMarkers = expandStringSet(v)
	}
	if v, ok := tfMap["additional_manifest"]; ok {
		result.AdditionalManifests = expandMediaConvertHlsAdditionalManifest(v.([]interface{}))
	}
	if v, ok := tfMap["audio_only_header"].(string); ok && v != "" {
		result.AudioOnlyHeader = aws.String(v)
	}
	if v, ok := tfMap["base_url"].(string); ok && v != "" {
		result.BaseUrl = aws.String(v)
	}
	if v, ok := tfMap["caption_language_mapping"]; ok {
		result.CaptionLanguageMappings = expandMediaConvertHlsCaptionLanguageMapping(v.([]interface{}))
	}
	if v, ok := tfMap["caption_language_setting"].(string); ok && v != "" {
		result.CaptionLanguageSetting = aws.String(v)
	}
	if v, ok := tfMap["client_cache"].(string); ok && v != "" {
		result.ClientCache = aws.String(v)
	}
	if v, ok := tfMap["codec_specification"].(string); ok && v != "" {
		result.CodecSpecification = aws.String(v)
	}
	if v, ok := tfMap["destination"].(string); ok && v != "" {
		result.Destination = aws.String(v)
	}
	if v, ok := tfMap["destination_settings"]; ok {
		result.DestinationSettings = expandMediaConvertDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["directory_structure"].(string); ok && v != "" {
		result.DirectoryStructure = aws.String(v)
	}
	if v, ok := tfMap["encryption"]; ok {
		result.Encryption = expandMediaConvertHlsEncryptionSettings(v.([]interface{}))
	}
	if v, ok := tfMap["manifest_compression"].(string); ok && v != "" {
		result.ManifestCompression = aws.String(v)
	}
	if v, ok := tfMap["manifest_duration_format"].(string); ok && v != "" {
		result.ManifestDurationFormat = aws.String(v)
	}
	if v, ok := tfMap["min_final_segment_length"].(float64); ok {
		result.MinFinalSegmentLength = aws.Float64(float64(v))
	}
	if v, ok := tfMap["min_segment_length"].(int); ok {
		result.MinSegmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["output_selection"].(string); ok && v != "" {
		result.OutputSelection = aws.String(v)
	}
	if v, ok := tfMap["program_date_time"].(string); ok && v != "" {
		result.ProgramDateTime = aws.String(v)
	}
	if v, ok := tfMap["program_date_time_period"].(int); ok {
		result.ProgramDateTimePeriod = aws.Int64(int64(v))
	}
	if v, ok := tfMap["segment_control"].(string); ok && v != "" {
		result.SegmentControl = aws.String(v)
	}
	if v, ok := tfMap["segment_length"].(int); ok {
		result.SegmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["segments_per_subdirectory"].(int); ok {
		result.SegmentsPerSubdirectory = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stream_inf_resolution"].(string); ok && v != "" {
		result.StreamInfResolution = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata_id3_frame"].(string); ok && v != "" {
		result.TimedMetadataId3Frame = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata_id3_period"].(int); ok {
		result.TimedMetadataId3Period = aws.Int64(int64(v))
	}
	if v, ok := tfMap["timestamp_delta_milliseconds"].(int); ok {
		result.TimestampDeltaMilliseconds = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertHlsEncryptionSettings(list []interface{}) *mediaconvert.HlsEncryptionSettings {
	result := &mediaconvert.HlsEncryptionSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["constant_initialization_vector"].(string); ok && v != "" {
		result.ConstantInitializationVector = aws.String(v)
	}
	if v, ok := tfMap["encryption_method"].(string); ok && v != "" {
		result.EncryptionMethod = aws.String(v)
	}
	if v, ok := tfMap["initialization_vector_in_manifest"].(string); ok && v != "" {
		result.InitializationVectorInManifest = aws.String(v)
	}
	if v, ok := tfMap["offline_encrypted"].(string); ok && v != "" {
		result.OfflineEncrypted = aws.String(v)
	}
	if v, ok := tfMap["speke_key_provider"]; ok {
		result.SpekeKeyProvider = expandMediaConvertSpekeKeyProvider(v.([]interface{}))
	}
	if v, ok := tfMap["static_key_provider"]; ok {
		result.StaticKeyProvider = expandMediaConvertStaticKeyProvider(v.([]interface{}))
	}
	if v, ok := tfMap["type"].(string); ok && v != "" {
		result.Type = aws.String(v)
	}
	return result
}

func expandMediaConvertStaticKeyProvider(list []interface{}) *mediaconvert.StaticKeyProvider {
	result := &mediaconvert.StaticKeyProvider{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["key_format"].(string); ok && v != "" {
		result.KeyFormat = aws.String(v)
	}
	if v, ok := tfMap["key_format_versions"].(string); ok && v != "" {
		result.KeyFormatVersions = aws.String(v)
	}
	if v, ok := tfMap["static_key_value"].(string); ok && v != "" {
		result.StaticKeyValue = aws.String(v)
	}
	if v, ok := tfMap["url"].(string); ok && v != "" {
		result.Url = aws.String(v)
	}
	return result
}

func expandMediaConvertHlsCaptionLanguageMapping(list []interface{}) []*mediaconvert.HlsCaptionLanguageMapping {
	results := []*mediaconvert.HlsCaptionLanguageMapping{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.HlsCaptionLanguageMapping{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["caption_channel"].(int); ok {
			result.CaptionChannel = aws.Int64(int64(v))
		}
		if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
			result.CustomLanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_code"].(string); ok && v != "" {
			result.LanguageCode = aws.String(v)
		}
		if v, ok := tfMap["language_description"].(string); ok && v != "" {
			result.LanguageDescription = aws.String(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertHlsAdditionalManifest(list []interface{}) []*mediaconvert.HlsAdditionalManifest {
	results := []*mediaconvert.HlsAdditionalManifest{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.HlsAdditionalManifest{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["manifest_name_modifier"].(string); ok && v != "" {
			result.ManifestNameModifier = aws.String(v)
		}
		if v, ok := tfMap["selected_outputs"].(*schema.Set); ok && v.Len() > 0 {
			result.SelectedOutputs = expandStringSet(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertMsSmoothGroupSettings(list []interface{}) *mediaconvert.MsSmoothGroupSettings {
	result := &mediaconvert.MsSmoothGroupSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["additional_manifest"]; ok {
		result.AdditionalManifests = expandMediaConvertMsSmoothAdditionalManifest(v.([]interface{}))
	}
	if v, ok := tfMap["audio_deduplication"].(string); ok && v != "" {
		result.AudioDeduplication = aws.String(v)
	}
	if v, ok := tfMap["destination"].(string); ok && v != "" {
		result.Destination = aws.String(v)
	}
	if v, ok := tfMap["destination_settings"]; ok {
		result.DestinationSettings = expandMediaConvertDestinationSettings(v.([]interface{}))
	}
	if v, ok := tfMap["encryption"]; ok {
		result.Encryption = expandMediaConvertMsSmoothEncryptionSettings(v.([]interface{}))
	}
	if v, ok := tfMap["fragment_length"].(int); ok {
		result.FragmentLength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["manifest_encoding"].(string); ok && v != "" {
		result.ManifestEncoding = aws.String(v)
	}
	return result
}

func expandMediaConvertMsSmoothAdditionalManifest(list []interface{}) []*mediaconvert.MsSmoothAdditionalManifest {
	results := []*mediaconvert.MsSmoothAdditionalManifest{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.MsSmoothAdditionalManifest{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["manifest_name_modifier"].(string); ok && v != "" {
			result.ManifestNameModifier = aws.String(v)
		}
		if v, ok := tfMap["selected_outputs"].(*schema.Set); ok && v.Len() > 0 {
			result.SelectedOutputs = expandStringSet(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertDestinationSettings(list []interface{}) *mediaconvert.DestinationSettings {
	result := &mediaconvert.DestinationSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["s3_settings"]; ok {
		result.S3Settings = expandMediaConvertS3DestinationSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertS3DestinationSettings(list []interface{}) *mediaconvert.S3DestinationSettings {
	result := &mediaconvert.S3DestinationSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["access_control"]; ok {
		result.AccessControl = expandMediaConvertS3DestinationAccessControl(v.([]interface{}))
	}
	if v, ok := tfMap["encryption"]; ok {
		result.Encryption = expandMediaConvertS3EncryptionSettings(v.([]interface{}))
	}
	return result
}
func expandMediaConvertS3DestinationAccessControl(list []interface{}) *mediaconvert.S3DestinationAccessControl {
	result := &mediaconvert.S3DestinationAccessControl{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["canned_acl"].(string); ok && v != "" {
		result.CannedAcl = aws.String(v)
	}
	return result
}

func expandMediaConvertS3EncryptionSettings(list []interface{}) *mediaconvert.S3EncryptionSettings {
	result := &mediaconvert.S3EncryptionSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["encryption_type"].(string); ok && v != "" {
		result.EncryptionType = aws.String(v)
	}
	if v, ok := tfMap["kms_key_arn"].(string); ok && v != "" {
		result.KmsKeyArn = aws.String(v)
	}
	return result
}

func expandMediaConvertMsSmoothEncryptionSettings(list []interface{}) *mediaconvert.MsSmoothEncryptionSettings {
	result := &mediaconvert.MsSmoothEncryptionSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["speke_key_provider"]; ok {
		result.SpekeKeyProvider = expandMediaConvertSpekeKeyProvider(v.([]interface{}))
	}
	return result
}

func expandMediaConvertSpekeKeyProvider(list []interface{}) *mediaconvert.SpekeKeyProvider {
	result := &mediaconvert.SpekeKeyProvider{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["certificate_arn"].(string); ok && v != "" {
		result.CertificateArn = aws.String(v)
	}
	if v, ok := tfMap["resource_id"].(string); ok && v != "" {
		result.ResourceId = aws.String(v)
	}
	if v, ok := tfMap["system_ids"].(*schema.Set); ok && v.Len() > 0 {
		result.SystemIds = expandStringSet(v)
	}
	if v, ok := tfMap["url"].(string); ok && v != "" {
		result.Url = aws.String(v)
	}
	return result
}
func expandMediaConvertSpekeKeyProviderCmaf(list []interface{}) *mediaconvert.SpekeKeyProviderCmaf {
	result := &mediaconvert.SpekeKeyProviderCmaf{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["certificate_arn"].(string); ok && v != "" {
		result.CertificateArn = aws.String(v)
	}
	if v, ok := tfMap["dash_signaled_system_ids "].(*schema.Set); ok && v.Len() > 0 {
		result.DashSignaledSystemIds = expandStringSet(v)
	}
	if v, ok := tfMap["hls_signaled_system_ids  "].(*schema.Set); ok && v.Len() > 0 {
		result.HlsSignaledSystemIds = expandStringSet(v)
	}
	if v, ok := tfMap["resource_id"].(string); ok && v != "" {
		result.ResourceId = aws.String(v)
	}
	if v, ok := tfMap["url"].(string); ok && v != "" {
		result.Url = aws.String(v)
	}
	return result
}

func expandMediaConvertMotionAutomatedEncodingSettings(list []interface{}) *mediaconvert.AutomatedEncodingSettings {
	result := &mediaconvert.AutomatedEncodingSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["abr_settings"]; ok {
		result.AbrSettings = expandMediaConvertMotionAutomatedAbrSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertMotionAutomatedAbrSettings(list []interface{}) *mediaconvert.AutomatedAbrSettings {
	result := &mediaconvert.AutomatedAbrSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_abr_bitrate"].(int); ok {
		result.MaxAbrBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_renditions"].(int); ok {
		result.MaxRenditions = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_abr_bitrate"].(int); ok {
		result.MinAbrBitrate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNielsenNonLinearWatermarkSettings(list []interface{}) *mediaconvert.NielsenNonLinearWatermarkSettings {
	result := &mediaconvert.NielsenNonLinearWatermarkSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["active_watermark_process"].(string); ok && v != "" {
		result.ActiveWatermarkProcess = aws.String(v)
	}
	if v, ok := tfMap["adi_filename"].(string); ok && v != "" {
		result.AdiFilename = aws.String(v)
	}
	if v, ok := tfMap["asset_id"].(string); ok && v != "" {
		result.AssetId = aws.String(v)
	}
	if v, ok := tfMap["asset_name"].(string); ok && v != "" {
		result.AssetName = aws.String(v)
	}
	if v, ok := tfMap["cbet_source_id"].(string); ok && v != "" {
		result.CbetSourceId = aws.String(v)
	}
	if v, ok := tfMap["episode_id"].(string); ok && v != "" {
		result.EpisodeId = aws.String(v)
	}
	if v, ok := tfMap["metadata_destination"].(string); ok && v != "" {
		result.MetadataDestination = aws.String(v)
	}
	if v, ok := tfMap["source_id"].(int); ok {
		result.SourceId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["source_watermark_status"].(string); ok && v != "" {
		result.SourceWatermarkStatus = aws.String(v)
	}
	if v, ok := tfMap["tic_server_url"].(string); ok && v != "" {
		result.TicServerUrl = aws.String(v)
	}
	if v, ok := tfMap["unique_tic_per_audio_track"].(string); ok && v != "" {
		result.UniqueTicPerAudioTrack = aws.String(v)
	}

	return result
}

func expandMediaConvertMotionNielsenConfiguration(list []interface{}) *mediaconvert.NielsenConfiguration {
	result := &mediaconvert.NielsenConfiguration{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["breakout_code"].(int); ok {
		result.BreakoutCode = aws.Int64(int64(v))
	}
	if v, ok := tfMap["distributor_id"].(string); ok && v != "" {
		result.DistributorId = aws.String(v)
	}
	return result
}

func expandMediaConvertMotionImageInserter(list []interface{}) *mediaconvert.MotionImageInserter {
	result := &mediaconvert.MotionImageInserter{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate"]; ok {
		result.Framerate = expandMediaConvertMotionImageInsertionFramerate(v.([]interface{}))
	}
	if v, ok := tfMap["input"].(string); ok && v != "" {
		result.Input = aws.String(v)
	}
	if v, ok := tfMap["insertion_mode"].(string); ok && v != "" {
		result.InsertionMode = aws.String(v)
	}
	if v, ok := tfMap["offset"]; ok {
		result.Offset = expandMediaConvertMotionImageInsertionOffset(v.([]interface{}))
	}
	if v, ok := tfMap["playback"].(string); ok && v != "" {
		result.Playback = aws.String(v)
	}
	if v, ok := tfMap["start_time"].(string); ok && v != "" {
		result.StartTime = aws.String(v)
	}
	return result
}

func expandMediaConvertMotionImageInsertionOffset(list []interface{}) *mediaconvert.MotionImageInsertionOffset {
	result := &mediaconvert.MotionImageInsertionOffset{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["image_x"].(int); ok {
		result.ImageX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["image_y"].(int); ok {
		result.ImageY = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertMotionImageInsertionFramerate(list []interface{}) *mediaconvert.MotionImageInsertionFramerate {
	result := &mediaconvert.MotionImageInsertionFramerate{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_denominator"].(int); ok {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertInputTemplate(list []interface{}) []*mediaconvert.InputTemplate {
	results := []*mediaconvert.InputTemplate{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.InputTemplate{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["audio_selector_group"]; ok {
			result.AudioSelectorGroups = expandMediaConvertAudioSelectorGroups(v.([]interface{}))
		}
		if v, ok := tfMap["audio_selector"]; ok {
			result.AudioSelectors = expandMediaConvertAudioSelector(v.([]interface{}))
		}
		if v, ok := tfMap["caption_selector"]; ok {
			result.CaptionSelectors = expandMediaConvertCaptionSelector(v.([]interface{}))
		}
		if v, ok := tfMap["crop"]; ok {
			result.Crop = expandMediaConvertRectangle(v.([]interface{}))
		}
		if v, ok := tfMap["deblock_filter"].(string); ok && v != "" {
			result.DeblockFilter = aws.String(v)
		}
		if v, ok := tfMap["denoise_filter"].(string); ok && v != "" {
			result.DenoiseFilter = aws.String(v)
		}
		if v, ok := tfMap["filter_enable"].(string); ok && v != "" {
			result.FilterEnable = aws.String(v)
		}
		if v, ok := tfMap["filter_strength"].(int); ok {
			result.FilterStrength = aws.Int64(int64(v))
		}
		if v, ok := tfMap["image_inserter"]; ok {
			result.ImageInserter = expandMediaConvertImageInserter(v.([]interface{}))
		}
		if v, ok := tfMap["input_clipping"]; ok {
			result.InputClippings = expandMediaConvertInputClipping(v.([]interface{}))
		}
		if v, ok := tfMap["input_scan_type"].(string); ok && v != "" {
			result.InputScanType = aws.String(v)
		}
		if v, ok := tfMap["position"]; ok {
			result.Position = expandMediaConvertRectangle(v.([]interface{}))
		}
		if v, ok := tfMap["program_number"].(int); ok {
			result.ProgramNumber = aws.Int64(int64(v))
		}
		if v, ok := tfMap["psi_control"].(string); ok && v != "" {
			result.PsiControl = aws.String(v)
		}
		if v, ok := tfMap["timecode_source"].(string); ok && v != "" {
			result.TimecodeSource = aws.String(v)
		}
		if v, ok := tfMap["timecode_start"].(string); ok && v != "" {
			result.TimecodeStart = aws.String(v)
		}
		if v, ok := tfMap["video_selector"]; ok {
			result.VideoSelector = expandMediaConvertVideoSelector(v.([]interface{}))
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertVideoSelector(list []interface{}) *mediaconvert.VideoSelector {
	result := &mediaconvert.VideoSelector{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["alpha_behavior"].(string); ok && v != "" {
		result.AlphaBehavior = aws.String(v)
	}
	if v, ok := tfMap["color_space"].(string); ok && v != "" {
		result.ColorSpace = aws.String(v)
	}
	if v, ok := tfMap["color_space_usage"].(string); ok && v != "" {
		result.ColorSpaceUsage = aws.String(v)
	}
	if v, ok := tfMap["pid"].(int); ok {
		result.Pid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["program_number"].(int); ok {
		result.ProgramNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["rotate"].(string); ok && v != "" {
		result.Rotate = aws.String(v)
	}
	if v, ok := tfMap["hdr10_metadata"]; ok {
		result.Hdr10Metadata = expandMediaConvertHdr10Metadata(v.([]interface{}))
	}
	return result
}

func expandMediaConvertHdr10Metadata(list []interface{}) *mediaconvert.Hdr10Metadata {
	result := &mediaconvert.Hdr10Metadata{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["blue_primary_x"].(int); ok {
		result.BluePrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["blue_primary_y"].(int); ok {
		result.BluePrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["green_primary_x"].(int); ok {
		result.GreenPrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["green_primary_y"].(int); ok {
		result.GreenPrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_content_light_level"].(int); ok {
		result.MaxContentLightLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_frame_average_light_level"].(int); ok {
		result.MaxFrameAverageLightLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_luminance"].(int); ok {
		result.MaxLuminance = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_luminance"].(int); ok {
		result.MinLuminance = aws.Int64(int64(v))
	}
	if v, ok := tfMap["red_primary_x"].(int); ok {
		result.RedPrimaryX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["red_primary_y"].(int); ok {
		result.RedPrimaryY = aws.Int64(int64(v))
	}
	if v, ok := tfMap["white_point_x"].(int); ok {
		result.WhitePointX = aws.Int64(int64(v))
	}
	if v, ok := tfMap["white_point_y"].(int); ok {
		result.WhitePointY = aws.Int64(int64(v))
	}

	return result
}

func expandMediaConvertInputClipping(list []interface{}) []*mediaconvert.InputClipping {
	results := []*mediaconvert.InputClipping{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.InputClipping{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["end_timecode"].(string); ok && v != "" {
			result.EndTimecode = aws.String(v)
		}
		if v, ok := tfMap["start_timecode"].(string); ok && v != "" {
			result.StartTimecode = aws.String(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertImageInserter(list []interface{}) *mediaconvert.ImageInserter {
	result := &mediaconvert.ImageInserter{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["insertable_images"]; ok {
		result.InsertableImages = expandMediaConvertInsertableImage(v.([]interface{}))
	}
	return result
}

func expandMediaConvertInsertableImage(list []interface{}) []*mediaconvert.InsertableImage {
	results := []*mediaconvert.InsertableImage{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.InsertableImage{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["duration"].(int); ok {
			result.Duration = aws.Int64(int64(v))
		}
		if v, ok := tfMap["fade_in"].(int); ok {
			result.FadeIn = aws.Int64(int64(v))
		}
		if v, ok := tfMap["fade_out"].(int); ok {
			result.FadeOut = aws.Int64(int64(v))
		}
		if v, ok := tfMap["height"].(int); ok {
			result.Height = aws.Int64(int64(v))
		}
		if v, ok := tfMap["image_inserter_input"].(string); ok && v != "" {
			result.ImageInserterInput = aws.String(v)
		}
		if v, ok := tfMap["image_x"].(int); ok {
			result.ImageX = aws.Int64(int64(v))
		}
		if v, ok := tfMap["image_y"].(int); ok {
			result.ImageY = aws.Int64(int64(v))
		}
		if v, ok := tfMap["layer"].(int); ok {
			result.Layer = aws.Int64(int64(v))
		}
		if v, ok := tfMap["opacity"].(int); ok {
			result.Opacity = aws.Int64(int64(v))
		}
		if v, ok := tfMap["start_time"].(string); ok && v != "" {
			result.StartTime = aws.String(v)
		}
		if v, ok := tfMap["width"].(int); ok {
			result.Width = aws.Int64(int64(v))
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertCaptionSelector(list []interface{}) map[string]*mediaconvert.CaptionSelector {
	results := map[string]*mediaconvert.CaptionSelector{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.CaptionSelector{}
		tfMap := list[i].(map[string]interface{})
		currentName := ""
		if v, ok := tfMap["name"].(string); ok && v != "" {
			currentName = v
			if v, ok := tfMap["name"].(string); ok && v != "" {
				currentName = v
			}
			if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
				result.CustomLanguageCode = aws.String(v)
			}
			if v, ok := tfMap["language_code"].(string); ok && v != "" {
				result.LanguageCode = aws.String(v)
			}
			if v, ok := tfMap["source_settings"]; ok {
				result.SourceSettings = expandMediaConvertCaptionSourceSettings(v.([]interface{}))
			}
			if len(currentName) > 0 {
				results[currentName] = result
			}
		}
	}
	return results
}

func expandMediaConvertCaptionSourceSettings(list []interface{}) *mediaconvert.CaptionSourceSettings {
	result := &mediaconvert.CaptionSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["ancillary_source_settings"]; ok {
		result.AncillarySourceSettings = expandMediaConvertAncillarySourceSettings(v.([]interface{}))
	}
	if v, ok := tfMap["dvb_sub_source_settings"]; ok {
		result.DvbSubSourceSettings = expandMediaConvertDvbSubSourceSettings(v.([]interface{}))
	}
	if v, ok := tfMap["embedded_source_settings"]; ok {
		result.EmbeddedSourceSettings = expandMediaConvertEmbeddedSourceSettings(v.([]interface{}))
	}
	if v, ok := tfMap["file_source_settings"]; ok {
		result.FileSourceSettings = expandMediaConvertFileSourceSettings(v.([]interface{}))
	}
	if v, ok := tfMap["source_type"].(string); ok && v != "" {
		result.SourceType = aws.String(v)
	}
	if v, ok := tfMap["teletext_source_settings"]; ok {
		result.TeletextSourceSettings = expandMediaConvertTeletextSourceSettings(v.([]interface{}))
	}
	if v, ok := tfMap["track_source_settings"]; ok {
		result.TrackSourceSettings = expandMediaConvertTrackSourceSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertTeletextSourceSettings(list []interface{}) *mediaconvert.TeletextSourceSettings {
	result := &mediaconvert.TeletextSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["page_number"].(string); ok && v != "" {
		result.PageNumber = aws.String(v)
	}
	return result
}

func expandMediaConvertTrackSourceSettings(list []interface{}) *mediaconvert.TrackSourceSettings {
	result := &mediaconvert.TrackSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["track_number"].(int); ok {
		result.TrackNumber = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertFileSourceSettings(list []interface{}) *mediaconvert.FileSourceSettings {
	result := &mediaconvert.FileSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["convert_608_to_708"].(string); ok && v != "" {
		result.Convert608To708 = aws.String(v)
	}
	if v, ok := tfMap["framerate"]; ok {
		result.Framerate = expandMediaConvertCaptionSourceFramerate(v.([]interface{}))
	}
	if v, ok := tfMap["source_file"].(string); ok && v != "" {
		result.SourceFile = aws.String(v)
	}
	if v, ok := tfMap["time_delta"].(int); ok {
		result.TimeDelta = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertCaptionSourceFramerate(list []interface{}) *mediaconvert.CaptionSourceFramerate {
	result := &mediaconvert.CaptionSourceFramerate{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_denominator"].(int); ok {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok {
		result.FramerateNumerator = aws.Int64(int64(v))
	}

	return result
}

func expandMediaConvertEmbeddedSourceSettings(list []interface{}) *mediaconvert.EmbeddedSourceSettings {
	result := &mediaconvert.EmbeddedSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["convert_608_to_708"].(string); ok && v != "" {
		result.Convert608To708 = aws.String(v)
	}
	if v, ok := tfMap["source_608_channel_number"].(int); ok {
		result.Source608ChannelNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["source_608_track_number"].(int); ok {
		result.Source608TrackNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["terminate_captions"].(string); ok && v != "" {
		result.TerminateCaptions = aws.String(v)
	}
	return result
}

func expandMediaConvertDvbSubSourceSettings(list []interface{}) *mediaconvert.DvbSubSourceSettings {
	result := &mediaconvert.DvbSubSourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["pid"].(int); ok {
		result.Pid = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAncillarySourceSettings(list []interface{}) *mediaconvert.AncillarySourceSettings {
	result := &mediaconvert.AncillarySourceSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["convert_608_to_708"].(string); ok && v != "" {
		result.Convert608To708 = aws.String(v)
	}
	if v, ok := tfMap["source_ancillary_channel_number"].(int); ok {
		result.SourceAncillaryChannelNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["terminate_captions"].(string); ok && v != "" {
		result.TerminateCaptions = aws.String(v)
	}
	return result
}

func expandMediaConvertAudioSelector(list []interface{}) map[string]*mediaconvert.AudioSelector {
	results := map[string]*mediaconvert.AudioSelector{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.AudioSelector{}
		tfMap := list[i].(map[string]interface{})
		currentName := ""
		if v, ok := tfMap["name"].(string); ok && v != "" {
			currentName = v
			if v, ok := tfMap["name"].(string); ok && v != "" {
				currentName = v
			}
			if v, ok := tfMap["custom_language_code"].(string); ok && v != "" {
				result.CustomLanguageCode = aws.String(v)
			}
			if v, ok := tfMap["default_selection"].(string); ok && v != "" {
				result.DefaultSelection = aws.String(v)
			}
			if v, ok := tfMap["external_audio_file_input"].(string); ok && v != "" {
				result.ExternalAudioFileInput = aws.String(v)
			}
			if v, ok := tfMap["language_code"].(string); ok && v != "" {
				result.LanguageCode = aws.String(v)
			}
			if v, ok := tfMap["offset"].(int); ok {
				result.Offset = aws.Int64(int64(v))
			}
			if v, ok := tfMap["pids"].(*schema.Set); ok && v.Len() > 0 {
				result.Pids = expandInt64Set(v)
			}
			if v, ok := tfMap["program_selection"].(int); ok {
				result.ProgramSelection = aws.Int64(int64(v))
			}
			if v, ok := tfMap["remix_settings"]; ok {
				result.RemixSettings = expandMediaConvertRemixSettings(v.([]interface{}))
			}
			if v, ok := tfMap["selector_type"].(string); ok && v != "" {
				result.SelectorType = aws.String(v)
			}
			if v, ok := tfMap["tracks"].(*schema.Set); ok && v.Len() > 0 {
				result.Tracks = expandInt64Set(v)
			}
			if len(currentName) > 0 {
				results[currentName] = result
			}
		}
	}
	return results
}

func expandMediaConvertRemixSettings(list []interface{}) *mediaconvert.RemixSettings {
	result := &mediaconvert.RemixSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["channel_mapping"]; ok {
		result.ChannelMapping = expandMediaConvertChannelMapping(v.([]interface{}))
	}
	if v, ok := tfMap["channels_in"].(int); ok {
		result.ChannelsIn = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels_out"].(int); ok {
		result.ChannelsOut = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertChannelMapping(list []interface{}) *mediaconvert.ChannelMapping {
	result := &mediaconvert.ChannelMapping{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["output_channel"]; ok {
		result.OutputChannels = expandMediaConvertOutputChannelMapping(v.([]interface{}))
	}
	return result
}

func expandMediaConvertOutputChannelMapping(list []interface{}) []*mediaconvert.OutputChannelMapping {
	results := []*mediaconvert.OutputChannelMapping{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.OutputChannelMapping{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["input_channels"].(*schema.Set); ok && v.Len() > 0 {
			result.InputChannels = expandInt64Set(v)
		}
		if v, ok := tfMap["input_channels_fine_tune"].(*schema.Set); ok && v.Len() > 0 {
			result.InputChannelsFineTune = expandFloat64Set(v)
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertAudioSelectorGroups(list []interface{}) map[string]*mediaconvert.AudioSelectorGroup {
	results := map[string]*mediaconvert.AudioSelectorGroup{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		tmp := &mediaconvert.AudioSelectorGroup{}
		tfMap := list[i].(map[string]interface{})
		currentName := ""
		if v, ok := tfMap["name"].(string); ok && v != "" {
			currentName = v
		}
		if v, ok := tfMap["audio_selector_names"].(*schema.Set); ok && v.Len() > 0 {
			tmp.AudioSelectorNames = expandStringSet(v)
		}
		if len(currentName) > 0 {
			results[currentName] = tmp
		}
	}
	return results
}

func expandMediaConvertAvailBlanking(list []interface{}) *mediaconvert.AvailBlanking {
	result := &mediaconvert.AvailBlanking{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["avail_blanking_image"].(string); ok && v != "" {
		result.AvailBlankingImage = aws.String(v)
	}
	return result
}

func expandMediaConvertEsamSettings(list []interface{}) *mediaconvert.EsamSettings {
	result := &mediaconvert.EsamSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["manifest_confirm_condition_notification"]; ok {
		result.ManifestConfirmConditionNotification = expandMediaConvertEsamManifestConfirmConditionNotification(v.([]interface{}))
	}
	if v, ok := tfMap["signal_processing_notification"]; ok {
		result.SignalProcessingNotification = expandMediaConvertEsamSignalProcessingNotification(v.([]interface{}))
	}
	if v, ok := tfMap["response_signal_preroll"].(int); ok {
		result.ResponseSignalPreroll = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertEsamManifestConfirmConditionNotification(list []interface{}) *mediaconvert.EsamManifestConfirmConditionNotification {
	result := &mediaconvert.EsamManifestConfirmConditionNotification{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["mcc_xml"].(string); ok && v != "" {
		result.MccXml = aws.String(v)
	}
	return result
}

func expandMediaConvertEsamSignalProcessingNotification(list []interface{}) *mediaconvert.EsamSignalProcessingNotification {
	result := &mediaconvert.EsamSignalProcessingNotification{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["scc_xml"].(string); ok && v != "" {
		result.SccXml = aws.String(v)
	}
	return result
}

func expandMediaConvertJobTemplateAccelerationSettings(list []interface{}) *mediaconvert.AccelerationSettings {
	result := &mediaconvert.AccelerationSettings{}
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["mode"].(string); ok && v != "" {
		result.Mode = aws.String(v)
	}
	return result
}

func expandMediaConvertJobTemplateHopDestinations(list []interface{}) []*mediaconvert.HopDestination {
	results := make([]*mediaconvert.HopDestination, 0)
	if len(list) == 0 || list[0] == nil {
		return nil
	}
	for i := 0; i < len(list); i++ {
		result := &mediaconvert.HopDestination{}
		tfMap := list[i].(map[string]interface{})
		if v, ok := tfMap["priority"].(int); ok {
			result.Priority = aws.Int64(int64(v))
		}
		if v, ok := tfMap["queue"].(string); ok && v != "" {
			result.Queue = aws.String(v)
		}
		if v, ok := tfMap["wait_minutes"].(int); ok {
			result.WaitMinutes = aws.Int64(int64(v))
		}
		results = append(results, result)
	}
	return results
}

func expandMediaConvertRectangle(list []interface{}) *mediaconvert.Rectangle {
	result := &mediaconvert.Rectangle{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["height"].(int); ok {
		result.Height = aws.Int64(int64(v))
	}
	if v, ok := tfMap["width"].(int); ok {
		result.Width = aws.Int64(int64(v))
	}
	if v, ok := tfMap["x"].(int); ok {
		result.X = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y"].(int); ok {
		result.Y = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDvbSubDestinationSettings(list []interface{}) *mediaconvert.DvbSubDestinationSettings {
	result := &mediaconvert.DvbSubDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["alignment"].(string); ok && v != "" {
		result.Alignment = aws.String(v)
	}
	if v, ok := tfMap["background_color"].(string); ok && v != "" {
		result.BackgroundColor = aws.String(v)
	}
	if v, ok := tfMap["background_opacity"].(int); ok {
		result.BackgroundOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_color"].(string); ok && v != "" {
		result.FontColor = aws.String(v)
	}
	if v, ok := tfMap["font_opacity"].(int); ok {
		result.FontOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_resolution"].(int); ok {
		result.FontResolution = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_script"].(string); ok && v != "" {
		result.FontScript = aws.String(v)
	}
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["outline_color"].(string); ok && v != "" {
		result.OutlineColor = aws.String(v)
	}
	if v, ok := tfMap["outline_size"].(int); ok {
		result.OutlineSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_color"].(string); ok && v != "" {
		result.ShadowColor = aws.String(v)
	}
	if v, ok := tfMap["shadow_opacity"].(int); ok {
		result.ShadowOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_x_offset"].(int); ok {
		result.ShadowXOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_y_offset"].(int); ok {
		result.ShadowYOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["subtitling_type"].(string); ok && v != "" {
		result.SubtitlingType = aws.String(v)
	}
	if v, ok := tfMap["teletext_spacing"].(string); ok && v != "" {
		result.TeletextSpacing = aws.String(v)
	}
	if v, ok := tfMap["x_position"].(int); ok {
		result.XPosition = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y_position"].(int); ok {
		result.YPosition = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertEmbeddedDestinationSettings(list []interface{}) *mediaconvert.EmbeddedDestinationSettings {
	result := &mediaconvert.EmbeddedDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["destination_608_channel_number"].(int); ok {
		result.Destination608ChannelNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["destination_708_service_number"].(int); ok {
		result.Destination708ServiceNumber = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertImscDestinationSettings(list []interface{}) *mediaconvert.ImscDestinationSettings {
	result := &mediaconvert.ImscDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["style_passthrough"].(string); ok && v != "" {
		result.StylePassthrough = aws.String(v)
	}
	return result
}

func expandMediaConvertSccDestinationSettings(list []interface{}) *mediaconvert.SccDestinationSettings {
	result := &mediaconvert.SccDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate"].(string); ok && v != "" {
		result.Framerate = aws.String(v)
	}
	return result
}

func expandMediaConvertTeletextDestinationSettings(list []interface{}) *mediaconvert.TeletextDestinationSettings {
	result := &mediaconvert.TeletextDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["page_number"].(string); ok && v != "" {
		result.PageNumber = aws.String(v)
	}
	result.PageTypes = expandStringSet(tfMap["page_types"].(*schema.Set))
	return result
}

func expandMediaConvertTtmlDestinationSettings(list []interface{}) *mediaconvert.TtmlDestinationSettings {
	result := &mediaconvert.TtmlDestinationSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["style_passthrough"].(string); ok && v != "" {
		result.StylePassthrough = aws.String(v)
	}
	return result
}

func expandMediaConvertBurninDestinationSettings(list []interface{}) *mediaconvert.BurninDestinationSettings {
	result := &mediaconvert.BurninDestinationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["alignment"].(string); ok && v != "" {
		result.Alignment = aws.String(v)
	}
	if v, ok := tfMap["background_color"].(string); ok && v != "" {
		result.BackgroundColor = aws.String(v)
	}
	if v, ok := tfMap["background_opacity"].(int); ok {
		result.BackgroundOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_color"].(string); ok && v != "" {
		result.FontColor = aws.String(v)
	}
	if v, ok := tfMap["font_opacity"].(int); ok {
		result.FontOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_resolution"].(int); ok {
		result.FontResolution = aws.Int64(int64(v))
	}
	if v, ok := tfMap["font_script"].(string); ok && v != "" {
		result.FontScript = aws.String(v)
	}
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["outline_color"].(string); ok && v != "" {
		result.OutlineColor = aws.String(v)
	}
	if v, ok := tfMap["outline_size"].(int); ok {
		result.OutlineSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_color"].(string); ok && v != "" {
		result.ShadowColor = aws.String(v)
	}
	if v, ok := tfMap["shadow_opacity"].(int); ok {
		result.ShadowOpacity = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_x_offset"].(int); ok {
		result.ShadowXOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["shadow_y_offset"].(int); ok {
		result.ShadowYOffset = aws.Int64(int64(v))
	}
	if v, ok := tfMap["teletext_spacing"].(string); ok && v != "" {
		result.TeletextSpacing = aws.String(v)
	}
	if v, ok := tfMap["x_position"].(int); ok {
		result.XPosition = aws.Int64(int64(v))
	}
	if v, ok := tfMap["y_position"].(int); ok {
		result.YPosition = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAudioChannelTagging(list []interface{}) *mediaconvert.AudioChannelTaggingSettings {
	result := &mediaconvert.AudioChannelTaggingSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["channel_tag"].(string); ok && v != "" {
		result.ChannelTag = aws.String(v)
	}
	return result
}

func expandMediaConvertAudioNormalizationSettings(list []interface{}) *mediaconvert.AudioNormalizationSettings {
	result := &mediaconvert.AudioNormalizationSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["algorithm"].(string); ok && v != "" {
		result.Algorithm = aws.String(v)
	}
	if v, ok := tfMap["algorithm_control"].(string); ok && v != "" {
		result.AlgorithmControl = aws.String(v)
	}
	if v, ok := tfMap["correction_gate_level"].(int); ok {
		result.CorrectionGateLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["loudness_logging"].(string); ok && v != "" {
		result.LoudnessLogging = aws.String(v)
	}
	if v, ok := tfMap["peak_calculation"].(string); ok && v != "" {
		result.PeakCalculation = aws.String(v)
	}
	if v, ok := tfMap["target_lkfs"].(float64); ok {
		result.TargetLkfs = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertCodecSettings(list []interface{}) *mediaconvert.AudioCodecSettings {
	result := &mediaconvert.AudioCodecSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["codec"].(string); ok && v != "" {
		result.Codec = aws.String(v)
	}
	if v, ok := tfMap["aac_settings"]; ok {
		result.AacSettings = expandMediaConvertAacSettings(v.([]interface{}))
	}
	if v, ok := tfMap["ac3_settings"]; ok {
		result.Ac3Settings = expandMediaConvertAc3Settings(v.([]interface{}))
	}
	if v, ok := tfMap["aiff_settings"]; ok {
		result.AiffSettings = expandMediaConvertAiffSettings(v.([]interface{}))
	}
	if v, ok := tfMap["eac3_atmos_settings"]; ok {
		result.Eac3AtmosSettings = expandMediaConvertEac3AtmosSettings(v.([]interface{}))
	}
	if v, ok := tfMap["eac3_settings"]; ok {
		result.Eac3Settings = expandMediaConvertEac3Settings(v.([]interface{}))
	}
	if v, ok := tfMap["mp2_settings"]; ok {
		result.Mp2Settings = expandMediaConvertMp2Settings(v.([]interface{}))
	}
	if v, ok := tfMap["mp3_settings"]; ok {
		result.Mp3Settings = expandMediaConvertMp3Settings(v.([]interface{}))
	}
	if v, ok := tfMap["opus_settings"]; ok {
		result.OpusSettings = expandMediaConvertOpusSettings(v.([]interface{}))
	}
	if v, ok := tfMap["vorbis_settings"]; ok {
		result.VorbisSettings = expandMediaConvertVorbisSettings(v.([]interface{}))
	}
	if v, ok := tfMap["wav_settings"]; ok {
		result.WavSettings = expandMediaConvertWavSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertAacSettings(list []interface{}) *mediaconvert.AacSettings {
	result := &mediaconvert.AacSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_description_broadcaster_mix"].(string); ok && v != "" {
		result.AudioDescriptionBroadcasterMix = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok && v != 0 {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["raw_format"].(string); ok && v != "" {
		result.RawFormat = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok && v != 0 {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["specification"].(string); ok && v != "" {
		result.Specification = aws.String(v)
	}
	if v, ok := tfMap["vbr_quality"].(string); ok && v != "" {
		result.VbrQuality = aws.String(v)
	}
	return result
}

func expandMediaConvertAc3Settings(list []interface{}) *mediaconvert.Ac3Settings {
	result := &mediaconvert.Ac3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dialnorm"].(int); ok {
		result.Dialnorm = aws.Int64(int64(v))
	}
	if v, ok := tfMap["dynamic_range_compression_profile"].(string); ok && v != "" {
		result.DynamicRangeCompressionProfile = aws.String(v)
	}
	if v, ok := tfMap["lfe_filter"].(string); ok && v != "" {
		result.LfeFilter = aws.String(v)
	}
	if v, ok := tfMap["metadata_control"].(string); ok && v != "" {
		result.MetadataControl = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAiffSettings(list []interface{}) *mediaconvert.AiffSettings {
	result := &mediaconvert.AiffSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitdepth"].(int); ok {
		result.BitDepth = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertEac3AtmosSettings(list []interface{}) *mediaconvert.Eac3AtmosSettings {
	result := &mediaconvert.Eac3AtmosSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dialogue_intelligence"].(string); ok && v != "" {
		result.DialogueIntelligence = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_line"].(string); ok && v != "" {
		result.DynamicRangeCompressionLine = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_rf"].(string); ok && v != "" {
		result.DynamicRangeCompressionRf = aws.String(v)
	}
	if v, ok := tfMap["lo_ro_center_mix_level"].(float64); ok {
		result.LoRoCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lo_ro_surround_mix_level"].(float64); ok {
		result.LoRoSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_center_mix_level"].(float64); ok {
		result.LtRtCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_surround_mix_level"].(float64); ok {
		result.LtRtSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["metering_mode"].(string); ok && v != "" {
		result.MeteringMode = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["speech_threshold"].(int); ok {
		result.SpeechThreshold = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stereo_downmix"].(string); ok && v != "" {
		result.StereoDownmix = aws.String(v)
	}
	if v, ok := tfMap["surround_ex_mode"].(string); ok && v != "" {
		result.SurroundExMode = aws.String(v)
	}
	return result
}

func expandMediaConvertEac3Settings(list []interface{}) *mediaconvert.Eac3Settings {
	result := &mediaconvert.Eac3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["attenuation_control"].(string); ok && v != "" {
		result.AttenuationControl = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["bitstream_mode"].(string); ok && v != "" {
		result.BitstreamMode = aws.String(v)
	}
	if v, ok := tfMap["coding_mode"].(string); ok && v != "" {
		result.CodingMode = aws.String(v)
	}
	if v, ok := tfMap["dc_filter"].(string); ok && v != "" {
		result.DcFilter = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Dialnorm = aws.Int64(int64(v))
	}
	if v, ok := tfMap["dynamic_range_compression_line"].(string); ok && v != "" {
		result.DynamicRangeCompressionLine = aws.String(v)
	}
	if v, ok := tfMap["dynamic_range_compression_rf"].(string); ok && v != "" {
		result.DynamicRangeCompressionRf = aws.String(v)
	}
	if v, ok := tfMap["lfe_control"].(string); ok && v != "" {
		result.LfeControl = aws.String(v)
	}
	if v, ok := tfMap["lfe_filter"].(string); ok && v != "" {
		result.LfeFilter = aws.String(v)
	}
	if v, ok := tfMap["lo_ro_center_mix_level"].(float64); ok {
		result.LoRoCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lo_ro_surround_mix_level"].(float64); ok {
		result.LoRoSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_center_mix_level"].(float64); ok {
		result.LtRtCenterMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["lt_rt_surround_mix_level"].(float64); ok {
		result.LtRtSurroundMixLevel = aws.Float64(float64(v))
	}
	if v, ok := tfMap["metadata_control"].(string); ok && v != "" {
		result.MetadataControl = aws.String(v)
	}
	if v, ok := tfMap["passthrough_control"].(string); ok && v != "" {
		result.PassthroughControl = aws.String(v)
	}
	if v, ok := tfMap["phase_control"].(string); ok && v != "" {
		result.PhaseControl = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["stereo_downmix"].(string); ok && v != "" {
		result.StereoDownmix = aws.String(v)
	}
	if v, ok := tfMap["surround_ex_mode"].(string); ok && v != "" {
		result.SurroundExMode = aws.String(v)
	}
	if v, ok := tfMap["surround_mode"].(string); ok && v != "" {
		result.SurroundMode = aws.String(v)
	}
	return result
}

func expandMediaConvertMp2Settings(list []interface{}) *mediaconvert.Mp2Settings {
	result := &mediaconvert.Mp2Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertMp3Settings(list []interface{}) *mediaconvert.Mp3Settings {
	result := &mediaconvert.Mp3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["vbr_quality"].(int); ok {
		result.VbrQuality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertOpusSettings(list []interface{}) *mediaconvert.OpusSettings {
	result := &mediaconvert.OpusSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertVorbisSettings(list []interface{}) *mediaconvert.VorbisSettings {
	result := &mediaconvert.VorbisSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["vbr_quality"].(int); ok {
		result.VbrQuality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertWavSettings(list []interface{}) *mediaconvert.WavSettings {
	result := &mediaconvert.WavSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitdepth"].(int); ok {
		result.BitDepth = aws.Int64(int64(v))
	}
	if v, ok := tfMap["channels"].(int); ok {
		result.Channels = aws.Int64(int64(v))
	}
	if v, ok := tfMap["format"].(string); ok && v != "" {
		result.Format = aws.String(v)
	}
	if v, ok := tfMap["sample_rate"].(int); ok {
		result.SampleRate = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertContainerSettings(list []interface{}) *mediaconvert.ContainerSettings {
	result := &mediaconvert.ContainerSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	containerSettingsMap := list[0].(map[string]interface{})
	if v, ok := containerSettingsMap["cmfc_settings"]; ok {
		result.CmfcSettings = expandMediaConvertCmfcSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["container"].(string); ok && v != "" {
		result.Container = aws.String(v)
	}
	if v, ok := containerSettingsMap["f4v_settings"]; ok {
		result.F4vSettings = expandMediaConvertF4vSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["m2ts_settings"]; ok {
		result.M2tsSettings = expandMediaConvertM2tsSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["m3u8_settings"]; ok {
		result.M3u8Settings = expandMediaConvertM3u8Settings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mov_settings"]; ok {
		result.MovSettings = expandMediaConvertMovSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mp4_settings"]; ok {
		result.Mp4Settings = expandMediaConvertMp4Settings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mpd_settings"]; ok {
		result.MpdSettings = expandMediaConvertMpdSettings(v.([]interface{}))
	}
	if v, ok := containerSettingsMap["mxf_settings"]; ok {
		result.MxfSettings = expandMediaConvertMxfSettings(v.([]interface{}))
	}

	return result
}

func expandMediaConvertMxfSettings(list []interface{}) *mediaconvert.MxfSettings {
	result := &mediaconvert.MxfSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["afd_signaling"].(string); ok && v != "" {
		result.AfdSignaling = aws.String(v)
	}
	if v, ok := tfMap["profile"].(string); ok && v != "" {
		result.Profile = aws.String(v)
	}
	return result
}

func expandMediaConvertMpdSettings(list []interface{}) *mediaconvert.MpdSettings {
	result := &mediaconvert.MpdSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["accessibility_caption_hints"].(string); ok && v != "" {
		result.AccessibilityCaptionHints = aws.String(v)
	}
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["caption_container_type"].(string); ok && v != "" {
		result.CaptionContainerType = aws.String(v)
	}
	if v, ok := tfMap["scte_35_esam"].(string); ok && v != "" {
		result.Scte35Esam = aws.String(v)
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	return result
}

func expandMediaConvertMp4Settings(list []interface{}) *mediaconvert.Mp4Settings {
	result := &mediaconvert.Mp4Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["cslg_atom"].(string); ok && v != "" {
		result.CslgAtom = aws.String(v)
	}
	if v, ok := tfMap["ctts_version"].(int); ok {
		result.CttsVersion = aws.Int64(int64(v))
	}
	if v, ok := tfMap["free_space_box"].(string); ok && v != "" {
		result.FreeSpaceBox = aws.String(v)
	}
	if v, ok := tfMap["moov_placement"].(string); ok && v != "" {
		result.MoovPlacement = aws.String(v)
	}
	if v, ok := tfMap["mp4_major_brand"].(string); ok && v != "" {
		result.Mp4MajorBrand = aws.String(v)
	}
	return result
}

func expandMediaConvertMovSettings(list []interface{}) *mediaconvert.MovSettings {
	result := &mediaconvert.MovSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["clap_atom"].(string); ok && v != "" {
		result.ClapAtom = aws.String(v)
	}
	if v, ok := tfMap["cslg_atom"].(string); ok && v != "" {
		result.CslgAtom = aws.String(v)
	}
	if v, ok := tfMap["mpeg2_fourcc_control"].(string); ok && v != "" {
		result.Mpeg2FourCCControl = aws.String(v)
	}
	if v, ok := tfMap["padding_control"].(string); ok && v != "" {
		result.PaddingControl = aws.String(v)
	}
	if v, ok := tfMap["reference"].(string); ok && v != "" {
		result.Reference = aws.String(v)
	}
	return result
}

func expandMediaConvertM3u8Settings(list []interface{}) *mediaconvert.M3u8Settings {
	result := &mediaconvert.M3u8Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["audio_frames_per_pes"].(int); ok {
		result.AudioFramesPerPes = aws.Int64(int64(v))
	}
	if v, ok := tfMap["audio_pids"].(*schema.Set); ok && v.Len() > 0 {
		result.AudioPids = expandInt64Set(v)
	}
	if v, ok := tfMap["nielsen_id3"].(string); ok && v != "" {
		result.NielsenId3 = aws.String(v)
	}
	if v, ok := tfMap["pat_interval"].(int); ok {
		result.PatInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pcr_control"].(string); ok && v != "" {
		result.PcrControl = aws.String(v)
	}
	if v, ok := tfMap["pcr_pid"].(int); ok {
		result.PcrPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_interval"].(int); ok {
		result.PmtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_pid"].(int); ok {
		result.PmtPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["private_metadata_pid"].(int); ok {
		result.PrivateMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["program_number"].(int); ok {
		result.ProgramNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_pid"].(int); ok {
		result.Scte35Pid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata"].(string); ok && v != "" {
		result.TimedMetadata = aws.String(v)
	}
	if v, ok := tfMap["timed_metadata_pid"].(int); ok {
		result.TimedMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["transport_stream_id"].(int); ok {
		result.TransportStreamId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["video_pid"].(int); ok {
		result.VideoPid = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertM2tsSettings(list []interface{}) *mediaconvert.M2tsSettings {
	result := &mediaconvert.M2tsSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_buffer_model"].(string); ok && v != "" {
		result.AudioBufferModel = aws.String(v)
	}
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["audio_frames_per_pes"].(int); ok {
		result.AudioFramesPerPes = aws.Int64(int64(v))
	}
	if v, ok := tfMap["protocols"].(*schema.Set); ok && v.Len() > 0 {
		result.AudioPids = expandInt64Set(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["buffer_model"].(string); ok && v != "" {
		result.BufferModel = aws.String(v)
	}
	if v, ok := tfMap["dvb_nit_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbNitSettings = expandMediaConvertDvbNitSettings(v.List())
	}
	if v, ok := tfMap["dvb_sdt_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbSdtSettings = expandMediaConvertDvbSdtSettings(v.List())
	}
	if v, ok := tfMap["dvb_sub_pids"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbSubPids = expandInt64Set(v)
	}
	if v, ok := tfMap["dvb_tdt_settings"].(*schema.Set); ok && v.Len() > 0 {
		result.DvbTdtSettings = expandMediaConvertDvbTdtSettings(v.List())
	}
	if v, ok := tfMap["dvb_teletext_pid"].(int); ok {
		result.DvbTeletextPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["ebp_audio_interval"].(string); ok && v != "" {
		result.EbpAudioInterval = aws.String(tfMap["ebp_audio_interval"].(string))
	}
	if v, ok := tfMap["ebp_placement"].(string); ok && v != "" {
		result.EbpPlacement = aws.String(v)
	}
	if v, ok := tfMap["es_rate_in_pes"].(string); ok && v != "" {
		result.EsRateInPes = aws.String(v)
	}
	if v, ok := tfMap["force_ts_video_ebp_order"].(string); ok && v != "" {
		result.ForceTsVideoEbpOrder = aws.String(v)
	}
	if v, ok := tfMap["fragment_time"].(float64); ok {
		result.FragmentTime = aws.Float64(float64(v))
	}
	if v, ok := tfMap["max_pcr_interval"].(int); ok {
		result.MaxPcrInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_ebp_interval"].(int); ok {
		result.MinEbpInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["nielsen_id3"].(string); ok && v != "" {
		result.NielsenId3 = aws.String(v)
	}
	if v, ok := tfMap["null_packet_bitrate"].(float64); ok {
		result.NullPacketBitrate = aws.Float64(float64(v))
	}
	if v, ok := tfMap["pat_interval"].(int); ok {
		result.PatInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pcr_control"].(string); ok && v != "" {
		result.PcrControl = aws.String(v)
	}
	if v, ok := tfMap["pcr_pid"].(int); ok {
		result.PcrPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_interval"].(int); ok {
		result.PmtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["pmt_pid"].(int); ok {
		result.PmtPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["private_metadata_pid"].(int); ok {
		result.PrivateMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["program_number"].(int); ok {
		result.ProgramNumber = aws.Int64(int64(v))
	}
	if v, ok := tfMap["rate_mode"].(string); ok && v != "" {
		result.RateMode = aws.String(v)
	}
	if v, ok := tfMap["scte_35_esam"].(*schema.Set); ok && v.Len() > 0 {
		result.Scte35Esam = expandMediaConvertM2tsScte35Esam(v.List())
	}
	if v, ok := tfMap["scte_35_pid"].(int); ok {
		result.Scte35Pid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["scte_35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	if v, ok := tfMap["segmentation_markers"].(string); ok && v != "" {
		result.SegmentationMarkers = aws.String(v)
	}
	if v, ok := tfMap["segmentation_style"].(string); ok && v != "" {
		result.SegmentationStyle = aws.String(v)
	}
	if v, ok := tfMap["segmentation_time"].(float64); ok {
		result.SegmentationTime = aws.Float64(float64(v))
	}
	if v, ok := tfMap["timed_metadata_pid"].(int); ok {
		result.TimedMetadataPid = aws.Int64(int64(v))
	}
	if v, ok := tfMap["transport_stream_id"].(int); ok {
		result.TransportStreamId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["video_pid"].(int); ok {
		result.VideoPid = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertF4vSettings(list []interface{}) *mediaconvert.F4vSettings {
	result := &mediaconvert.F4vSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["moov_placement"].(string); ok && v != "" {
		result.MoovPlacement = aws.String(tfMap["moov_placement"].(string))
	}
	return result
}

func expandMediaConvertCmfcSettings(list []interface{}) *mediaconvert.CmfcSettings {
	result := &mediaconvert.CmfcSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["audio_duration"].(string); ok && v != "" {
		result.AudioDuration = aws.String(v)
	}
	if v, ok := tfMap["scte35_esam"].(string); ok && v != "" {
		result.Scte35Esam = aws.String(v)
	}
	if v, ok := tfMap["scte35_source"].(string); ok && v != "" {
		result.Scte35Source = aws.String(v)
	}
	return result
}

func expandMediaConvertDvbNitSettings(list []interface{}) *mediaconvert.DvbNitSettings {
	result := &mediaconvert.DvbNitSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["network_id"].(int); ok {
		result.NetworkId = aws.Int64(int64(v))
	}
	if v, ok := tfMap["network_name"].(string); ok && v != "" {
		result.NetworkName = aws.String(v)
	}
	if v, ok := tfMap["nit_interval"].(int); ok {
		result.NitInterval = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertDvbSdtSettings(list []interface{}) *mediaconvert.DvbSdtSettings {
	result := &mediaconvert.DvbSdtSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["output_sdt"].(string); ok && v != "" {
		result.OutputSdt = aws.String(v)
	}
	if v, ok := tfMap["sdt_interval"].(int); ok {
		result.SdtInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["service_name"].(string); ok && v != "" {
		result.ServiceName = aws.String(v)
	}
	if v, ok := tfMap["service_provider_name"].(string); ok && v != "" {
		result.ServiceProviderName = aws.String(v)
	}

	return result
}

func expandMediaConvertDvbTdtSettings(list []interface{}) *mediaconvert.DvbTdtSettings {
	result := &mediaconvert.DvbTdtSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["tdt_interval"].(int); ok {
		result.TdtInterval = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertM2tsScte35Esam(list []interface{}) *mediaconvert.M2tsScte35Esam {
	result := &mediaconvert.M2tsScte35Esam{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["scte_35_esam_pid"].(int); ok {
		result.Scte35EsamPid = aws.Int64(int64(v))
	}
	return result
}

func flattenMediaConvertAccelerationSettings(cfg *mediaconvert.AccelerationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"mode": aws.StringValue(cfg.Mode),
	}
	return []interface{}{m}
}

func flattenMediaConvertHopDestinations(cfg []*mediaconvert.HopDestination) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"priority":     aws.Int64Value(cfg[i].Priority),
			"queue":        aws.StringValue(cfg[i].Queue),
			"wait_minutes": aws.Int64Value(cfg[i].WaitMinutes),
		}
		results = append(results, m)
	}
	return results
}

//###############################################################
//###############################################################
//###############################################################
//###############################################################

func flattenMediaConvertJobTemplateSettings(cfg *mediaconvert.JobTemplateSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"ad_avail_offset":              aws.Int64Value(cfg.AdAvailOffset),
		"avail_blanking":               flattenMediaConvertAvailBlanking(cfg.AvailBlanking),
		"esam":                         flattenMediaConvertEsamSettings(cfg.Esam),
		"input":                        flattenMediaConvertInputTemplate(cfg.Inputs),
		"motion_image_inserter":        flattenMediaConvertMotionImageInserter(cfg.MotionImageInserter),
		"nielsen_configuration":        flattenMediaConvertNielsenConfiguration(cfg.NielsenConfiguration),
		"nielsen_non_linear_watermark": flattenMediaConvertNielsenNonLinearWatermarkSettings(cfg.NielsenNonLinearWatermark),
		"output_group":                 flattenMediaConvertOutputGroup(cfg.OutputGroups),
		"timecode_config":              flattenMediaConvertTimecodeConfig(cfg.TimecodeConfig),
		"timed_metadata_insertion":     flattenMediaConvertTimedMetadataInsertion(cfg.TimedMetadataInsertion),
	}
	return []interface{}{m}
}

func flattenMediaConvertTimecodeConfig(cfg *mediaconvert.TimecodeConfig) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"anchor":           aws.StringValue(cfg.Anchor),
		"source":           aws.StringValue(cfg.Source),
		"start":            aws.StringValue(cfg.Start),
		"timestamp_offset": aws.StringValue(cfg.TimestampOffset),
	}
	return []interface{}{m}
}

func flattenMediaConvertTimedMetadataInsertion(cfg *mediaconvert.TimedMetadataInsertion) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"id3_insertion": flattenMediaConvertId3Insertion(cfg.Id3Insertions),
	}
	return []interface{}{m}
}

func flattenMediaConvertId3Insertion(cfg []*mediaconvert.Id3Insertion) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"id3":      aws.StringValue(cfg[i].Id3),
			"timecode": aws.StringValue(cfg[i].Timecode),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertInputTemplate(cfg []*mediaconvert.InputTemplate) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"audio_selector_group": flattenMediaConvertAudioSelectorGroup(cfg[i].AudioSelectorGroups),
			"audio_selector":       flattenMediaConvertAudioSelector(cfg[i].AudioSelectors),
			"caption_selector":     flattenMediaConvertCaptionSelector(cfg[i].CaptionSelectors),
			"crop":                 flattenMediaConvertRectangle(cfg[i].Crop),
			"deblock_filter":       aws.StringValue(cfg[i].DeblockFilter),
			"denoise_filter":       aws.StringValue(cfg[i].DenoiseFilter),
			"filter_enable":        aws.StringValue(cfg[i].FilterEnable),
			"filter_strength":      aws.Int64Value(cfg[i].FilterStrength),
			"image_inserter":       flattenMediaConvertImageInserter(cfg[i].ImageInserter),
			"input_clipping":       flattenMediaConvertInputClipping(cfg[i].InputClippings),
			"input_scan_type":      aws.StringValue(cfg[i].InputScanType),
			"position":             flattenMediaConvertRectangle(cfg[i].Position),
			"program_number":       aws.Int64Value(cfg[i].ProgramNumber),
			"psi_control":          aws.StringValue(cfg[i].PsiControl),
			"timecode_source":      aws.StringValue(cfg[i].TimecodeSource),
			"timecode_start":       aws.StringValue(cfg[i].TimecodeStart),
			"video_selector":       flattenMediaConvertVideoSelector(cfg[i].VideoSelector),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertAudioSelectorGroup(cfg map[string]*mediaconvert.AudioSelectorGroup) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for k, v := range cfg {
		m := map[string]interface{}{
			"name":                 k,
			"audio_selector_names": flattenStringSet(v.AudioSelectorNames),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertEsamSettings(cfg *mediaconvert.EsamSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"manifest_confirm_condition_notification": flattenMediaConvertEsamManifestConfirmConditionNotification(cfg.ManifestConfirmConditionNotification),
		"signal_processing_notification":          flattenMediaConvertEsamSignalProcessingNotification(cfg.SignalProcessingNotification),
		"response_signal_preroll":                 aws.Int64Value(cfg.ResponseSignalPreroll),
	}
	return []interface{}{m}
}

func flattenMediaConvertEsamSignalProcessingNotification(cfg *mediaconvert.EsamSignalProcessingNotification) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"scc_xml": aws.StringValue(cfg.SccXml),
	}
	return []interface{}{m}
}

func flattenMediaConvertEsamManifestConfirmConditionNotification(cfg *mediaconvert.EsamManifestConfirmConditionNotification) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"mcc_xml": aws.StringValue(cfg.MccXml),
	}
	return []interface{}{m}
}

func flattenMediaConvertAvailBlanking(cfg *mediaconvert.AvailBlanking) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"avail_blanking_image": aws.StringValue(cfg.AvailBlankingImage),
	}
	return []interface{}{m}
}

func flattenMediaConvertAudioSelector(cfg map[string]*mediaconvert.AudioSelector) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for k, v := range cfg {
		m := map[string]interface{}{
			"name":                      k,
			"custom_language_code":      aws.StringValue(v.CustomLanguageCode),
			"default_selection":         aws.StringValue(v.DefaultSelection),
			"external_audio_file_input": aws.StringValue(v.ExternalAudioFileInput),
			"language_code":             aws.StringValue(v.LanguageCode),
			"offset":                    aws.Int64Value(v.Offset),
			"pids":                      flattenInt64Set(v.Pids),
			"program_selection":         aws.Int64Value(v.ProgramSelection),
			"remix_settings":            flattenMediaConvertRemixSettings(v.RemixSettings),
			"selector_type":             aws.StringValue(v.SelectorType),
			"tracks":                    flattenInt64Set(v.Tracks),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertCaptionSelector(cfg map[string]*mediaconvert.CaptionSelector) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for k, v := range cfg {
		m := map[string]interface{}{
			"name":                 k,
			"custom_language_code": aws.StringValue(v.CustomLanguageCode),
			"language_code":        aws.StringValue(v.LanguageCode),
			"source_settings":      flattenMediaConvertCaptionSourceSettings(v.SourceSettings),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertCaptionSourceSettings(cfg *mediaconvert.CaptionSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"ancillary_source_settings": flattenMediaConvertAncillarySourceSettings(cfg.AncillarySourceSettings),
		"dvb_sub_source_settings":   flattenMediaConvertDvbSubSourceSettings(cfg.DvbSubSourceSettings),
		"embedded_source_settings":  flattenMediaConvertEmbeddedSourceSettings(cfg.EmbeddedSourceSettings),
		"file_source_settings":      flattenMediaConvertFileSourceSettings(cfg.FileSourceSettings),
		"source_type":               aws.StringValue(cfg.SourceType),
		"teletext_source_settings":  flattenMediaConvertTeletextSourceSettings(cfg.TeletextSourceSettings),
		"track_source_settings":     flattenMediaConvertTrackSourceSettings(cfg.TrackSourceSettings),
	}
	return []interface{}{m}
}
func flattenMediaConvertAncillarySourceSettings(cfg *mediaconvert.AncillarySourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"convert_608_to_708":              aws.StringValue(cfg.Convert608To708),
		"source_ancillary_channel_number": aws.Int64Value(cfg.SourceAncillaryChannelNumber),
		"terminate_captions":              aws.StringValue(cfg.TerminateCaptions),
	}
	return []interface{}{m}
}

func flattenMediaConvertDvbSubSourceSettings(cfg *mediaconvert.DvbSubSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"pid": aws.Int64Value(cfg.Pid),
	}
	return []interface{}{m}
}

func flattenMediaConvertEmbeddedSourceSettings(cfg *mediaconvert.EmbeddedSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"convert_608_to_708":        aws.StringValue(cfg.Convert608To708),
		"source_608_channel_number": aws.Int64Value(cfg.Source608ChannelNumber),
		"source_608_track_number":   aws.Int64Value(cfg.Source608TrackNumber),
		"terminate_captions":        aws.StringValue(cfg.TerminateCaptions),
	}
	return []interface{}{m}
}

func flattenMediaConvertCaptionSourceFramerate(cfg *mediaconvert.CaptionSourceFramerate) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate_denominator": aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":   aws.Int64Value(cfg.FramerateNumerator),
	}
	return []interface{}{m}
}

func flattenMediaConvertFileSourceSettings(cfg *mediaconvert.FileSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"convert_608_to_708": aws.StringValue(cfg.Convert608To708),
		"framerate":          flattenMediaConvertCaptionSourceFramerate(cfg.Framerate),
		"source_file":        aws.StringValue(cfg.SourceFile),
		"time_delta":         aws.Int64Value(cfg.TimeDelta),
	}
	return []interface{}{m}
}

func flattenMediaConvertTeletextSourceSettings(cfg *mediaconvert.TeletextSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"page_number": aws.StringValue(cfg.PageNumber),
	}
	return []interface{}{m}
}

func flattenMediaConvertTrackSourceSettings(cfg *mediaconvert.TrackSourceSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"track_number": aws.Int64Value(cfg.TrackNumber),
	}
	return []interface{}{m}
}

func flattenMediaConvertImageInserter(cfg *mediaconvert.ImageInserter) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"insertable_images": flattenMediaConvertInsertableImage(cfg.InsertableImages),
	}
	return []interface{}{m}
}

func flattenMediaConvertInsertableImage(cfg []*mediaconvert.InsertableImage) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"duration":             aws.Int64Value(cfg[i].Duration),
			"fade_in":              aws.Int64Value(cfg[i].FadeIn),
			"fade_out":             aws.Int64Value(cfg[i].FadeOut),
			"height":               aws.Int64Value(cfg[i].Height),
			"image_inserter_input": aws.StringValue(cfg[i].ImageInserterInput),
			"image_x":              aws.Int64Value(cfg[i].ImageX),
			"image_y":              aws.Int64Value(cfg[i].ImageY),
			"layer":                aws.Int64Value(cfg[i].Layer),
			"opacity":              aws.Int64Value(cfg[i].Opacity),
			"start_time":           aws.StringValue(cfg[i].StartTime),
			"width":                aws.Int64Value(cfg[i].Width),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertInputClipping(cfg []*mediaconvert.InputClipping) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"end_timecode":   aws.StringValue(cfg[i].EndTimecode),
			"start_timecode": aws.StringValue(cfg[i].StartTimecode),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertVideoSelector(cfg *mediaconvert.VideoSelector) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"alpha_behavior":    aws.StringValue(cfg.AlphaBehavior),
		"color_space":       aws.StringValue(cfg.ColorSpace),
		"color_space_usage": aws.StringValue(cfg.ColorSpaceUsage),
		"pid":               aws.Int64Value(cfg.Pid),
		"program_number":    aws.Int64Value(cfg.ProgramNumber),
		"rotate":            aws.StringValue(cfg.Rotate),
		"hdr10_metadata":    flattenMediaConvertHdr10Metadata(cfg.Hdr10Metadata),
	}
	return []interface{}{m}
}

func flattenMediaConvertHdr10Metadata(cfg *mediaconvert.Hdr10Metadata) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"blue_primary_x":                aws.Int64Value(cfg.BluePrimaryX),
		"blue_primary_y":                aws.Int64Value(cfg.BluePrimaryY),
		"green_primary_x":               aws.Int64Value(cfg.GreenPrimaryX),
		"green_primary_y":               aws.Int64Value(cfg.GreenPrimaryY),
		"max_content_light_level":       aws.Int64Value(cfg.MaxContentLightLevel),
		"max_frame_average_light_level": aws.Int64Value(cfg.MaxFrameAverageLightLevel),
		"max_luminance":                 aws.Int64Value(cfg.MaxLuminance),
		"min_luminance":                 aws.Int64Value(cfg.MinLuminance),
		"red_primary_x":                 aws.Int64Value(cfg.RedPrimaryX),
		"red_primary_y":                 aws.Int64Value(cfg.RedPrimaryY),
		"white_point_x":                 aws.Int64Value(cfg.WhitePointX),
		"white_point_y":                 aws.Int64Value(cfg.WhitePointY),
	}
	return []interface{}{m}
}

func flattenMediaConvertMotionImageInserter(cfg *mediaconvert.MotionImageInserter) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate":      flattenMediaConvertMotionImageInsertionFramerate(cfg.Framerate),
		"input":          aws.StringValue(cfg.Input),
		"insertion_mode": aws.StringValue(cfg.InsertionMode),
		"offset":         flattenMediaConvertMotionImageInsertionOffset(cfg.Offset),
		"playback":       aws.StringValue(cfg.Playback),
		"start_time":     aws.StringValue(cfg.StartTime),
	}
	return []interface{}{m}
}

func flattenMediaConvertMotionImageInsertionFramerate(cfg *mediaconvert.MotionImageInsertionFramerate) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"framerate_denominator": aws.Int64Value(cfg.FramerateDenominator),
		"framerate_numerator":   aws.Int64Value(cfg.FramerateNumerator),
	}
	return []interface{}{m}
}

func flattenMediaConvertMotionImageInsertionOffset(cfg *mediaconvert.MotionImageInsertionOffset) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"image_x": aws.Int64Value(cfg.ImageX),
		"image_y": aws.Int64Value(cfg.ImageY),
	}
	return []interface{}{m}
}

func flattenMediaConvertNielsenConfiguration(cfg *mediaconvert.NielsenConfiguration) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"breakout_code":  aws.Int64Value(cfg.BreakoutCode),
		"distributor_id": aws.StringValue(cfg.DistributorId),
	}
	return []interface{}{m}
}

func flattenMediaConvertNielsenNonLinearWatermarkSettings(cfg *mediaconvert.NielsenNonLinearWatermarkSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"active_watermark_process":   aws.StringValue(cfg.ActiveWatermarkProcess),
		"adi_filename":               aws.StringValue(cfg.AdiFilename),
		"asset_id":                   aws.StringValue(cfg.AssetId),
		"asset_name":                 aws.StringValue(cfg.AssetName),
		"cbet_source_id":             aws.StringValue(cfg.CbetSourceId),
		"episode_id":                 aws.StringValue(cfg.EpisodeId),
		"metadata_destination":       aws.StringValue(cfg.MetadataDestination),
		"source_id":                  aws.Int64Value(cfg.SourceId),
		"source_watermark_status":    aws.StringValue(cfg.SourceWatermarkStatus),
		"tic_server_url":             aws.StringValue(cfg.TicServerUrl),
		"unique_tic_per_audio_track": aws.StringValue(cfg.UniqueTicPerAudioTrack),
	}
	return []interface{}{m}
}

func flattenMediaConvertOutputGroup(cfg []*mediaconvert.OutputGroup) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"automated_encoding_settings": flattenMediaConvertAutomatedEncodingSettings(cfg[i].AutomatedEncodingSettings),
			"custom_name":                 aws.StringValue(cfg[i].CustomName),
			"name":                        aws.StringValue(cfg[i].Name),
			"output_group_settings":       flattenMediaConvertOutputGroupSettings(cfg[i].OutputGroupSettings),
			"output":                      flattenMediaConvertOutput(cfg[i].Outputs),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertAutomatedEncodingSettings(cfg *mediaconvert.AutomatedEncodingSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"abr_settings": flattenMediaConvertAutomatedAbrSettings(cfg.AbrSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertAutomatedAbrSettings(cfg *mediaconvert.AutomatedAbrSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"max_abr_bitrate": aws.Int64Value(cfg.MaxAbrBitrate),
		"max_renditions":  aws.Int64Value(cfg.MaxRenditions),
		"min_abr_bitrate": aws.Int64Value(cfg.MinAbrBitrate),
	}
	return []interface{}{m}
}

func flattenMediaConvertOutputGroupSettings(cfg *mediaconvert.OutputGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"cmaf_group_settings":      flattenMediaConvertCmafGroupSettings(cfg.CmafGroupSettings),
		"dash_iso_group_settings":  flattenMediaConvertDashIsoGroupSettings(cfg.DashIsoGroupSettings),
		"file_group_settings":      flattenMediaConvertFileGroupSettings(cfg.FileGroupSettings),
		"hls_group_settings":       flattenMediaConvertHlsGroupSettings(cfg.HlsGroupSettings),
		"ms_smooth_group_settings": flattenMediaConvertMsSmoothGroupSettings(cfg.MsSmoothGroupSettings),
		"type":                     aws.StringValue(cfg.Type),
	}
	return []interface{}{m}
}

func flattenMediaConvertMsSmoothGroupSettings(cfg *mediaconvert.MsSmoothGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"additional_manifest":  flattenMediaConvertMsSmoothAdditionalManifest(cfg.AdditionalManifests),
		"audio_deduplication":  aws.StringValue(cfg.AudioDeduplication),
		"destination":          aws.StringValue(cfg.Destination),
		"destination_settings": flattenMediaConvertDestinationSettings(cfg.DestinationSettings),
		"encryption":           flattenMediaConvertMsSmoothEncryptionSettings(cfg.Encryption),
		"fragment_length":      aws.Int64Value(cfg.FragmentLength),
		"manifest_encoding":    aws.StringValue(cfg.ManifestEncoding),
	}
	return []interface{}{m}
}

func flattenMediaConvertMsSmoothEncryptionSettings(cfg *mediaconvert.MsSmoothEncryptionSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"speke_key_provider": flattenMediaConvertSpekeKeyProvider(cfg.SpekeKeyProvider),
	}
	return []interface{}{m}
}

func flattenMediaConvertMsSmoothAdditionalManifest(cfg []*mediaconvert.MsSmoothAdditionalManifest) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"manifest_name_modifier": aws.StringValue(cfg[i].ManifestNameModifier),
			"selected_outputs":       flattenStringSet(cfg[i].SelectedOutputs),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertHlsGroupSettings(cfg *mediaconvert.HlsGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"ad_markers":                   flattenStringSet(cfg.AdMarkers),
		"additional_manifest":          flattenMediaConvertHlsAdditionalManifest(cfg.AdditionalManifests),
		"audio_only_header":            aws.StringValue(cfg.AudioOnlyHeader),
		"base_url":                     aws.StringValue(cfg.BaseUrl),
		"caption_language_mapping":     flattenMediaConvertHlsCaptionLanguageMapping(cfg.CaptionLanguageMappings),
		"client_cache":                 aws.StringValue(cfg.ClientCache),
		"codec_specification":          aws.StringValue(cfg.CodecSpecification),
		"destination":                  aws.StringValue(cfg.Destination),
		"destination_settings":         flattenMediaConvertDestinationSettings(cfg.DestinationSettings),
		"directory_structure":          aws.StringValue(cfg.DirectoryStructure),
		"encryption":                   flattenMediaConvertHlsEncryptionSettings(cfg.Encryption),
		"manifest_compression":         aws.StringValue(cfg.ManifestCompression),
		"manifest_duration_format":     aws.StringValue(cfg.ManifestDurationFormat),
		"min_final_segment_length":     aws.Float64Value(cfg.MinFinalSegmentLength),
		"min_segment_length":           aws.Int64Value(cfg.MinSegmentLength),
		"output_selection":             aws.StringValue(cfg.OutputSelection),
		"program_date_time":            aws.StringValue(cfg.ProgramDateTime),
		"program_date_time_period":     aws.Int64Value(cfg.ProgramDateTimePeriod),
		"segment_control":              aws.StringValue(cfg.SegmentControl),
		"segment_length":               aws.Int64Value(cfg.SegmentLength),
		"segments_per_subdirectory":    aws.Int64Value(cfg.SegmentsPerSubdirectory),
		"stream_inf_resolution":        aws.StringValue(cfg.StreamInfResolution),
		"timed_metadata_id3_frame":     aws.StringValue(cfg.TimedMetadataId3Frame),
		"timed_metadata_id3_period":    aws.Int64Value(cfg.TimedMetadataId3Period),
		"timestamp_delta_milliseconds": aws.Int64Value(cfg.TimestampDeltaMilliseconds),
	}
	return []interface{}{m}
}

func flattenMediaConvertHlsEncryptionSettings(cfg *mediaconvert.HlsEncryptionSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"constant_initialization_vector":    aws.StringValue(cfg.ConstantInitializationVector),
		"encryption_method":                 aws.StringValue(cfg.EncryptionMethod),
		"initialization_vector_in_manifest": aws.StringValue(cfg.InitializationVectorInManifest),
		"offline_encrypted":                 aws.StringValue(cfg.OfflineEncrypted),
		"speke_key_provider":                flattenMediaConvertSpekeKeyProvider(cfg.SpekeKeyProvider),
		"static_key_provider":               flattenMediaConvertStaticKeyProvider(cfg.StaticKeyProvider),
		"type":                              aws.StringValue(cfg.Type),
	}
	return []interface{}{m}
}

func flattenMediaConvertHlsCaptionLanguageMapping(cfg []*mediaconvert.HlsCaptionLanguageMapping) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"caption_channel":      aws.Int64Value(cfg[i].CaptionChannel),
			"custom_language_code": aws.StringValue(cfg[i].CustomLanguageCode),
			"language_code":        aws.StringValue(cfg[i].LanguageCode),
			"language_description": aws.StringValue(cfg[i].LanguageDescription),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertHlsAdditionalManifest(cfg []*mediaconvert.HlsAdditionalManifest) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"manifest_name_modifier": aws.StringValue(cfg[i].ManifestNameModifier),
			"selected_outputs":       flattenStringSet(cfg[i].SelectedOutputs),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertFileGroupSettings(cfg *mediaconvert.FileGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"destination":          aws.StringValue(cfg.Destination),
		"destination_settings": flattenMediaConvertDestinationSettings(cfg.DestinationSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertDashIsoGroupSettings(cfg *mediaconvert.DashIsoGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"additional_manifest":      flattenMediaConvertDashAdditionalManifest(cfg.AdditionalManifests),
		"base_url":                 aws.StringValue(cfg.BaseUrl),
		"destination":              aws.StringValue(cfg.Destination),
		"destination_settings":     flattenMediaConvertDestinationSettings(cfg.DestinationSettings),
		"encryption":               flattenMediaConvertDashIsoEncryptionSettings(cfg.Encryption),
		"fragment_length":          aws.Int64Value(cfg.FragmentLength),
		"hbbtv_compliance":         aws.StringValue(cfg.HbbtvCompliance),
		"min_buffer_time":          aws.Int64Value(cfg.MinBufferTime),
		"min_final_segment_length": aws.Float64Value(cfg.MinFinalSegmentLength),
		"mpd_profile":              aws.StringValue(cfg.MpdProfile),
		"segment_control":          aws.StringValue(cfg.SegmentControl),
		"segment_length":           aws.Int64Value(cfg.SegmentLength),
		"write_segment_timeline_in_representation": aws.StringValue(cfg.WriteSegmentTimelineInRepresentation),
	}
	return []interface{}{m}
}

func flattenMediaConvertDashIsoEncryptionSettings(cfg *mediaconvert.DashIsoEncryptionSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"playback_device_compatibility": aws.StringValue(cfg.PlaybackDeviceCompatibility),
		"speke_key_provider":            flattenMediaConvertSpekeKeyProvider(cfg.SpekeKeyProvider),
	}
	return []interface{}{m}
}

func flattenMediaConvertSpekeKeyProvider(cfg *mediaconvert.SpekeKeyProvider) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"certificate_arn": aws.StringValue(cfg.CertificateArn),
		"resource_id":     aws.StringValue(cfg.ResourceId),
		"system_ids":      flattenStringSet(cfg.SystemIds),
		"url":             aws.StringValue(cfg.Url),
	}
	return []interface{}{m}
}

func flattenMediaConvertDashAdditionalManifest(cfg []*mediaconvert.DashAdditionalManifest) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"manifest_name_modifier": aws.StringValue(cfg[i].ManifestNameModifier),
			"selected_outputs":       flattenStringSet(cfg[i].SelectedOutputs),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertCmafGroupSettings(cfg *mediaconvert.CmafGroupSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"additional_manifest":                      flattenMediaConvertCmafAdditionalManifest(cfg.AdditionalManifests),
		"base_url":                                 aws.StringValue(cfg.BaseUrl),
		"client_cache":                             aws.StringValue(cfg.ClientCache),
		"code_specification":                       aws.StringValue(cfg.CodecSpecification),
		"destination":                              aws.StringValue(cfg.Destination),
		"destination_settings":                     flattenMediaConvertDestinationSettings(cfg.DestinationSettings),
		"encryption":                               flattenMediaConvertCmafEncryptionSettings(cfg.Encryption),
		"fragment_length":                          aws.Int64Value(cfg.FragmentLength),
		"manifest_compression":                     aws.StringValue(cfg.ManifestCompression),
		"manifest_duration_format":                 aws.StringValue(cfg.ManifestDurationFormat),
		"min_buffer_time":                          aws.Int64Value(cfg.MinBufferTime),
		"min_final_segment_length":                 aws.Float64Value(cfg.MinFinalSegmentLength),
		"mpd_profile":                              aws.StringValue(cfg.MpdProfile),
		"segment_control":                          aws.StringValue(cfg.SegmentControl),
		"segment_length":                           aws.Int64Value(cfg.SegmentLength),
		"stream_inf_resolution":                    aws.StringValue(cfg.StreamInfResolution),
		"write_dash_manifest":                      aws.StringValue(cfg.WriteDashManifest),
		"write_hls_manifest":                       aws.StringValue(cfg.WriteHlsManifest),
		"write_segment_timeline_in_representation": aws.StringValue(cfg.WriteSegmentTimelineInRepresentation),
	}
	return []interface{}{m}
}

func flattenMediaConvertCmafEncryptionSettings(cfg *mediaconvert.CmafEncryptionSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"constant_initialization_vector":    aws.StringValue(cfg.ConstantInitializationVector),
		"encryption_method":                 aws.StringValue(cfg.EncryptionMethod),
		"initialization_vector_in_manifest": aws.StringValue(cfg.InitializationVectorInManifest),
		"speke_key_provider":                flattenMediaConvertSpekeKeyProviderCmaf(cfg.SpekeKeyProvider),
		"static_key_provider":               flattenMediaConvertStaticKeyProvider(cfg.StaticKeyProvider),
		"type":                              aws.StringValue(cfg.Type),
	}
	return []interface{}{m}
}

func flattenMediaConvertSpekeKeyProviderCmaf(cfg *mediaconvert.SpekeKeyProviderCmaf) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"certificate_arn":          aws.StringValue(cfg.CertificateArn),
		"dash_signaled_system_ids": flattenStringSet(cfg.DashSignaledSystemIds),
		"hls_signaled_system_ids":  flattenStringSet(cfg.HlsSignaledSystemIds),
		"resource_id":              aws.StringValue(cfg.ResourceId),
		"url":                      aws.StringValue(cfg.Url),
	}
	return []interface{}{m}
}

func flattenMediaConvertStaticKeyProviders(cfg []*mediaconvert.StaticKeyProvider) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"key_format":          aws.StringValue(cfg[i].KeyFormat),
			"key_format_versions": aws.StringValue(cfg[i].KeyFormatVersions),
			"static_key_value":    aws.StringValue(cfg[i].StaticKeyValue),
			"url":                 aws.StringValue(cfg[i].Url),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertStaticKeyProvider(cfg *mediaconvert.StaticKeyProvider) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"key_format":          aws.StringValue(cfg.KeyFormat),
		"key_format_versions": aws.StringValue(cfg.KeyFormatVersions),
		"static_key_value":    aws.StringValue(cfg.StaticKeyValue),
		"url":                 aws.StringValue(cfg.Url),
	}
	return []interface{}{m}
}

func flattenMediaConvertCmafAdditionalManifest(cfg []*mediaconvert.CmafAdditionalManifest) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"manifest_name_modifier": aws.StringValue(cfg[i].ManifestNameModifier),
			"selected_outputs":       flattenStringSet(cfg[i].SelectedOutputs),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertDestinationSettings(cfg *mediaconvert.DestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"s3_settings": flattenMediaConvertS3DestinationSettings(cfg.S3Settings),
	}
	return []interface{}{m}
}

func flattenMediaConvertS3DestinationSettings(cfg *mediaconvert.S3DestinationSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"access_control": flattenMediaConvertS3DestinationAccessControl(cfg.AccessControl),
		"encryption":     flattenMediaConvertS3EncryptionSettings(cfg.Encryption),
	}
	return []interface{}{m}
}

func flattenMediaConvertS3DestinationAccessControl(cfg *mediaconvert.S3DestinationAccessControl) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"canned_acl": aws.StringValue(cfg.CannedAcl),
	}
	return []interface{}{m}
}

func flattenMediaConvertS3EncryptionSettings(cfg *mediaconvert.S3EncryptionSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"encryption_type": aws.StringValue(cfg.EncryptionType),
		"kms_key_arn":     aws.StringValue(cfg.KmsKeyArn),
	}
	return []interface{}{m}
}

func flattenMediaConvertOutput(cfg []*mediaconvert.Output) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"audio_description":   flattenMediaConvertAudioDescription(cfg[i].AudioDescriptions),
			"caption_description": flattenMediaConvertCaptionDescription(cfg[i].CaptionDescriptions),
			"container_settings":  flattenMediaConvertContainerSettings(cfg[i].ContainerSettings),
			"extension":           aws.StringValue(cfg[i].Extension),
			"name_modifier":       aws.StringValue(cfg[i].NameModifier),
			"output_settings":     flattenMediaConvertOutputSettings(cfg[i].OutputSettings),
			"preset":              aws.StringValue(cfg[i].Preset),
			"video_description":   flattenMediaConvertVideoDescription(cfg[i].VideoDescription),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertCaptionDescription(cfg []*mediaconvert.CaptionDescription) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	results := make([]interface{}, 0, len(cfg))
	for i := 0; i < len(cfg); i++ {
		m := map[string]interface{}{
			"caption_selector_name": aws.StringValue(cfg[i].CaptionSelectorName),
			"custom_language_code":  aws.StringValue(cfg[i].CustomLanguageCode),
			"destination_settings":  flattenMediaConvertCaptionDestinationSettings(cfg[i].DestinationSettings),
			"language_code":         aws.StringValue(cfg[i].LanguageCode),
			"language_description":  aws.StringValue(cfg[i].LanguageDescription),
		}
		results = append(results, m)
	}
	return results
}

func flattenMediaConvertOutputSettings(cfg *mediaconvert.OutputSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"hls_settings": flattenMediaConvertHlsSettings(cfg.HlsSettings),
	}
	return []interface{}{m}
}

func flattenMediaConvertHlsSettings(cfg *mediaconvert.HlsSettings) []interface{} {
	if cfg == nil {
		return []interface{}{}
	}
	m := map[string]interface{}{
		"audio_group_id":       aws.StringValue(cfg.AudioGroupId),
		"audio_only_container": aws.StringValue(cfg.AudioOnlyContainer),
		"audio_rendition_sets": aws.StringValue(cfg.AudioRenditionSets),
		"audio_track_type":     aws.StringValue(cfg.AudioTrackType),
		"iframe_only_manifest": aws.StringValue(cfg.IFrameOnlyManifest),
		"segment_modifier":     aws.StringValue(cfg.SegmentModifier),
	}
	return []interface{}{m}
}

// func flattenMediaConvert---(cfg *mediaconvert.---) []interface{} {
// 	if cfg == nil {
// 		return []interface{}{}
// 	}
// 	m := map[string]interface{}{
// 		"breakout_code":  aws.Int64Value(cfg.BreakoutCode),
// 		"distributor_id": aws.StringValue(cfg.DistributorId),
// 	}
// 	return []interface{}{m}
// }
