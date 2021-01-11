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
