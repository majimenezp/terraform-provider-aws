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
						"input": {
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
