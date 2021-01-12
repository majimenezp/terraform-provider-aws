package aws

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccAwsMediaConvertJobTemplate_base(t *testing.T) {
	var jobTemplate mediaconvert.JobTemplate
	resourceName := "aws_media_convert_job_template.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-base")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertJobTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertJobTemplateConfig_Basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertJobTemplateExists(resourceName, &jobTemplate),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`jobTemplates/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status_update_interval", "SECONDS_60"),
					resource.TestCheckResourceAttr(resourceName, "settings.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "settings.*", map[string]string{
						"input.#":                                               "1",
						"output_group.#":                                        "1",
						"input.0.audio_selector.#":                              "1",
						"input.0.audio_selector_group.#":                        "1",
						"input.0.caption_selector.#":                            "1",
						"input.0.psi_control":                                   "USE_PSI",
						"input.0.timecode_source":                               "EMBEDDED",
						"input.0.audio_selector_group.0.name":                   "Audio Selector Group 1",
						"input.0.audio_selector_group.0.audio_selector_names.#": "1",
						"input.0.audio_selector.0.name":                         "Audio Selector 1",
						"input.0.audio_selector.0.default_selection":            "DEFAULT",
						"input.0.caption_selector.0.name":                       "Captions Selector 1",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.convert_608_to_708":        "DISABLED",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.source_608_channel_number": "1",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.terminate_captions":        "END_OF_INPUT",
						"input.0.caption_selector.0.source_settings.0.source_type":                                          "EMBEDDED",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: testAccMediaConvertJobTemplateConfig_BasicUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertJobTemplateExists(resourceName, &jobTemplate),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status_update_interval", "SECONDS_300"),
					resource.TestCheckResourceAttr(resourceName, "settings.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "settings.*", map[string]string{
						"input.#":                                               "1",
						"output_group.#":                                        "1",
						"input.0.audio_selector.#":                              "1",
						"input.0.audio_selector_group.#":                        "1",
						"input.0.caption_selector.#":                            "1",
						"input.0.psi_control":                                   "IGNORE_PSI",
						"input.0.timecode_source":                               "ZEROBASED",
						"input.0.audio_selector_group.0.name":                   "Audio Selector Group 1",
						"input.0.audio_selector_group.0.audio_selector_names.#": "1",
						"input.0.audio_selector.0.name":                         "Audio Selector 1",
						"input.0.audio_selector.0.default_selection":            "DEFAULT",
						"input.0.caption_selector.0.name":                       "Captions Selector 1",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.convert_608_to_708":        "DISABLED",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.source_608_channel_number": "1",
						"input.0.caption_selector.0.source_settings.0.embedded_source_settings.0.terminate_captions":        "END_OF_INPUT",
						"input.0.caption_selector.0.source_settings.0.source_type":                                          "EMBEDDED",
					}),
				),
			},
		},
	})
}

func TestAccAwsMediaConvertJobTemplate_BasicMP4(t *testing.T) {
	var jobTemplate mediaconvert.JobTemplate
	resourceName := "aws_media_convert_job_template.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-base")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertJobTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertJobTemplateConfig_BasicMP4(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertJobTemplateExists(resourceName, &jobTemplate),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`jobTemplates/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "status_update_interval", "SECONDS_60"),
					resource.TestCheckResourceAttr(resourceName, "settings.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "settings.*", map[string]string{
						"input.#":                                    "1",
						"output_group.#":                             "1",
						"input.0.audio_selector.#":                   "1",
						"input.0.video_selector.#":                   "1",
						"input.0.psi_control":                        "USE_PSI",
						"input.0.timecode_source":                    "EMBEDDED",
						"input.0.audio_selector.0.name":              "Audio Selector 1",
						"input.0.audio_selector.0.default_selection": "DEFAULT",
						"input.0.audio_selector.0.program_selection": "1",
						"input.0.video_selector.0.alpha_behavior":    "DISCARD",
						"input.0.video_selector.0.color_space":       "FOLLOW",
						"input.0.video_selector.0.rotate":            "DEGREE_0",
						"output_group.0.output.#":                    "2",
						"output_group.0.output.0.preset":             "MP4",
						"output_group.0.output.1.preset":             "Framecapture",
					}),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccMediaConvertJobTemplateConfig_Basic(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_job_template" "test" {
		name = %[1]q
		description = "test job template"
		acceleration_settings {
			mode = "DISABLED"
		}
		priority = 0
		settings {
			ad_avail_offset = 0
			input {
					audio_selector_group {
						name = "Audio Selector Group 1"
						audio_selector_names = ["Audio Selector 1"]
					}
					
					audio_selector {
							name = "Audio Selector 1"
							default_selection = "DEFAULT"
							offset = 0
							program_selection = 1
					}
					
					caption_selector {
							name = "Captions Selector 1"
							source_settings {
								embedded_source_settings {
									convert_608_to_708 = "DISABLED"
									source_608_channel_number = 1
									source_608_track_number = 1
									terminate_captions = "END_OF_INPUT"
								}
								source_type = "EMBEDDED"
							}
					}
										        
					deblock_filter = "DISABLED"
					denoise_filter = "DISABLED"
					filter_enable = "AUTO"
					filter_strength = 0
					psi_control = "USE_PSI"
					timecode_source = "EMBEDDED"
					video_selector {
						alpha_behavior = "DISCARD"
						color_space = "FOLLOW"
						rotate = "DEGREE_0"
					}
			}
			output_group {
				name = "CMAF"
				output_group_settings {
					type = "CMAF_GROUP_SETTINGS"
					cmaf_group_settings {
						client_cache = "ENABLED"
						code_specification = "RFC_4281"
						destination_settings {
							s3_settings {
								encryption {
									encryption_type = "SERVER_SIDE_ENCRYPTION_S3"
								}
							}
						}
						fragment_length = 2
						manifest_compression = "NONE"
						manifest_duration_format = "INTEGER"
						min_final_segment_length = 0.0
						mpd_profile = "MAIN_PROFILE"
						segment_control = "SEGMENTED_FILES"
						segment_length = 30
						stream_inf_resolution = "INCLUDE"
						write_dash_manifest = "ENABLED"
						write_hls_manifest = "ENABLED"
						write_segment_timeline_in_representation = "ENABLED"
					}
				}
				output {
					name_modifier = "240p"
					preset = "240p"
				}
				output {
					preset = "Audio"
				}
				output {
					name_modifier = "360p"
					preset = "360p"
				}
				output {
					name_modifier = "432p"
					preset = "432p"
				}
				output {
					name_modifier = "480p"
					preset = "480p"
				}
				output {
					name_modifier = "576p"
					preset = "576p"
				}
				output {
					name_modifier = "720p"
					preset = "720p"
				}
				output {
					name_modifier = "1080p"
					preset = "1080p"
				}
				output {
					caption_description {
						caption_selector_name = "Captions Selector 1"
						destination_settings {
							destination_type = "WEBVTT"
						}
					}
					container_settings {
						container = "CMFC"
					}
				}
			}
		}
		status_update_interval = "SECONDS_60"
	}
	`, rName)
}

func testAccMediaConvertJobTemplateConfig_BasicUpdate(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_job_template" "test" {
		name = %[1]q
		description = "test job template"
		acceleration_settings {
			mode = "DISABLED"
		}
		priority = 0
		settings {
			ad_avail_offset = 0
			input {
					audio_selector_group {
						name = "Audio Selector Group 1"
						audio_selector_names = ["Audio Selector 1"]
					}
					
					audio_selector {
							name = "Audio Selector 1"
							default_selection = "DEFAULT"
							offset = 0
							program_selection = 1
					}
					
					caption_selector {
							name = "Captions Selector 1"
							source_settings {
								embedded_source_settings {
									convert_608_to_708 = "DISABLED"
									source_608_channel_number = 1
									source_608_track_number = 1
									terminate_captions = "END_OF_INPUT"
								}
								source_type = "EMBEDDED"
							}
					}
										        
					deblock_filter = "DISABLED"
					denoise_filter = "DISABLED"
					filter_enable = "AUTO"
					filter_strength = 0
					psi_control = "IGNORE_PSI"
					timecode_source = "ZEROBASED"
					video_selector {
						alpha_behavior = "DISCARD"
						color_space = "FOLLOW"
						rotate = "DEGREE_0"
					}
			}
			output_group {
				name = "CMAF"
				output_group_settings {
					type = "CMAF_GROUP_SETTINGS"
					cmaf_group_settings {
						client_cache = "ENABLED"
						code_specification = "RFC_4281"
						destination_settings {
							s3_settings {
								encryption {
									encryption_type = "SERVER_SIDE_ENCRYPTION_S3"
								}
							}
						}
						fragment_length = 2
						manifest_compression = "NONE"
						manifest_duration_format = "INTEGER"
						min_final_segment_length = 0.0
						mpd_profile = "MAIN_PROFILE"
						segment_control = "SEGMENTED_FILES"
						segment_length = 30
						stream_inf_resolution = "INCLUDE"
						write_dash_manifest = "ENABLED"
						write_hls_manifest = "ENABLED"
						write_segment_timeline_in_representation = "ENABLED"
					}
				}
				output {
					name_modifier = "240p"
					preset = "240p"
				}
				output {
					preset = "Audio"
				}
				output {
					name_modifier = "360p"
					preset = "360p"
				}
				output {
					name_modifier = "432p"
					preset = "432p"
				}
				output {
					name_modifier = "480p"
					preset = "480p"
				}
				output {
					name_modifier = "576p"
					preset = "576p"
				}
				output {
					name_modifier = "720p"
					preset = "720p"
				}
				output {
					name_modifier = "1080p"
					preset = "1080p"
				}
				output {
					caption_description {
						caption_selector_name = "Captions Selector 1"
						destination_settings {
							destination_type = "WEBVTT"
						}
					}
					container_settings {
						container = "CMFC"
					}
				}
			}
		}
		status_update_interval = "SECONDS_300"
	}
	`, rName)
}

func testAccMediaConvertJobTemplateConfig_BasicMP4(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_job_template" "test" {
		name = %[1]q
		description = "test MP4 job template"
		acceleration_settings {
			mode = "DISABLED"
		}
		priority = 0
		queue = "arn:aws:mediaconvert:us-east-1:583161073698:queues/jive-cloud-queue"
		status_update_interval = "SECONDS_60"
		settings {
			ad_avail_offset = 0
			input {					
					audio_selector {
						name = "Audio Selector 1"
						default_selection = "DEFAULT"
						offset = 0
						program_selection = 1
					}										        
					deblock_filter = "DISABLED"
					denoise_filter = "DISABLED"
					filter_enable = "AUTO"
					filter_strength = 0
					psi_control = "USE_PSI"
					timecode_source = "EMBEDDED"
					video_selector {
						alpha_behavior = "DISCARD"
						color_space = "FOLLOW"
						rotate = "DEGREE_0"
					}
			}
			output_group {
				name = "File Group"
				output_group_settings {
					type = "FILE_GROUP_SETTINGS"
					file_group_settings {
						destination_settings {
							s3_settings {
								encryption {
									encryption_type = "SERVER_SIDE_ENCRYPTION_S3"
								}
							}
						}
					}					
				}
				output {
					preset = "MP4"
				}
				output {
					preset = "Framecapture"
				}
			}
		}
		
	}
	`, rName)
}

func testAccMediaConvertJobTemplateConfig_Complex(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_job_template" "test" {
		name = %[1]q
		description = "test complex  job template"
		acceleration_settings {
			mode = "PREFERRED"
		}
		priority = 0
		queue = "arn:aws:mediaconvert:us-east-1:583161073698:queues/Default"
		status_update_interval = "SECONDS_420"
		hop_destinations {
			wait_minutes = 10
			queue = "arn:aws:mediaconvert:us-east-1:583161073698:queues/jive-cloud-queue"
			priority = 0
		}
		settings {
			ad_avail_offset = 0
			timecode_config {
				source = "ZEROBASED"
			}
			output_group {
				name = "File Group"
				custom_name = "test group 01"
				output {
					preset = "MP4"
					name_modifier = "H264Video"
				}
				output {
					preset = "Framecapture"
					name_modifier = "JPEGVideo"
				}
				output {
					name_modifier = "Test1"
					container_settings {
						container = "MP4"
						mp4_settings {
							cslg_atom = "INCLUDE"
							ctts_version = 0
							free_space_box = "EXCLUDE"
							moov_placement = "PROGRESSIVE_DOWNLOAD"
							audio_duration = "DEFAULT_CODEC_DURATION"
						}
					}
					video_description {
						width = 320
						scaling_behavior = "STRETCH_TO_OUTPUT"
						height = 240
						timecode_insertion = "DISABLED"
						anti_alias = "ENABLED"
						sharpness = 50
						drop_frame_timecode = "ENABLED"
						color_metadata = "INSERT"
						video_preprocessors {
							color_corrector {
								brightness = 60
								color_space_conversion = "FORCE_601"
								contrast = 55
								hue = 0
								saturation = 50
							}
							deinterlacer {
								algorithm = "BLEND"
								control = "FORCE_ALL_FRAMES"
								mode = "ADAPTIVE"
							}
							noise_reducer {
								filter = "MEAN"
								filter_settings {
									strength = 1
								}
							}
							timecode_burnin {
								font_size = 16
								position = "MIDDLE_CENTER"
								prefix = "test"
							}
						}
						codec_settings {
							codec = "AV1"
							av1_settings {
								gop_size = 81
								number_b_frames_between_reference_frames = 15
								slices = 1
								rate_control_mode = "QVBR"
								max_bitrate = 6000000
								adaptive_quantization = "HIGHER"
								spatial_adaptive_quantization = "ENABLED"
								framerate_control = "SPECIFIED"
								framerate_conversion_algorithm = "INTERPOLATE"
								framerate_numerator = 24000
								framerate_denominator = 1001
								qvbr_settings {
									qvbr_quality_level = 3
									qvbr_quality_level_fine_tune = 0.33
								}
							}
						}
					}
					audio_description {
						audio_type_control = "FOLLOW_INPUT"
						audio_source_name = "Audio Selector 1"
						stream_name = "main"
						language_code_control = "FOLLOW_INPUT"
						language_code = "ENG"
						audio_normalization_settings {
							algorithm = "ITU_BS_1770_1"
							algorithm_control = "CORRECT_AUDIO"
							correction_gate_level = -70
							loudness_logging = "LOG"
							peak_calculation = "NONE"
						}
						codec_settings {
							codec = "AC3"
							ac3_settings {
								bitrate = 128000
								bitstream_mode = "EMERGENCY"
								coding_mode = "CODING_MODE_1_1"
								dynamic_range_compression_profile = "FILM_STANDARD"
								metadata_control = "FOLLOW_INPUT"
								sample_rate = 48000
							} 
						}
					}
				}
				output_group_settings {
					type = "FILE_GROUP_SETTINGS"
					file_group_settings {
						destination_settings {
							s3_settings {
								encryption {
									encryption_type = "SERVER_SIDE_ENCRYPTION_S3"
								}
							}
							access_control {
								canned_acl = "AUTHENTICATED_READ"
							}
						}
					}					
				}
				
			}
			motion_image_inserter {
				input = "s3://jive-video-upload-us-east-1-iac/testimage_000.png"
				start_time = "00:00:05:00"
				playback = "ONCE"
				framerate {
					framerate_denominator = 1
					framerate_numerator = 1
				}
				offset {
					image_x = 0
					image_y = 0
				}
			}
			input {					
					audio_selector {
						name = "Audio Selector 1"
						default_selection = "DEFAULT"
						offset = 0
						program_selection = 1
					}										        
					deblock_filter = "DISABLED"
					denoise_filter = "DISABLED"
					filter_enable = "AUTO"
					filter_strength = 0
					psi_control = "USE_PSI"
					timecode_source = "EMBEDDED"
					video_selector {
						alpha_behavior = "DISCARD"
						color_space = "FOLLOW"
						rotate = "DEGREE_0"
					}
			}
			
		}		
	}
	`, rName)
}

func testAccCheckAwsMediaConvertJobTemplateDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_media_convert_job_template" {
			continue
		}
		conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		if err != nil {
			return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		}

		_, err = conn.GetJobTemplate(&mediaconvert.GetJobTemplateInput{
			Name: aws.String(rs.Primary.ID),
		})
		if err != nil {
			if isAWSErr(err, mediaconvert.ErrCodeNotFoundException, "") {
				continue
			}
			return err
		}
	}
	return nil
}

func testAccCheckAwsMediaConvertJobTemplateExists(n string, jobTemplate *mediaconvert.JobTemplate) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No Preset id is set")
		}

		conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		if err != nil {
			return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		}

		resp, err := conn.GetJobTemplate(&mediaconvert.GetJobTemplateInput{
			Name: aws.String(rs.Primary.ID),
		})
		if err != nil {
			return fmt.Errorf("Error getting job template: %s", err)
		}

		*jobTemplate = *resp.JobTemplate
		return nil
	}
}
