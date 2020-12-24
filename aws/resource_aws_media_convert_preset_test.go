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

func TestAccAwsMediaConvertPreset_base(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-base")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_Basic(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "settings.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "settings.*", map[string]string{
						"video_description.#":                                                                     "1",
						"audio_description.#":                                                                     "1",
						"container_settings.#":                                                                    "1",
						"video_description.0.timecode_insertion":                                                  "DISABLED",
						"audio_description.0.codec_settings.0.aac_settings.#":                                     "1",
						"audio_description.0.codec_settings.0.aac_settings.0.codec_profile":                       "LC",
						"video_description.0.codec_settings.0.h264_settings.0.quality_tuning_level":               "SINGLE_PASS",
						"video_description.0.codec_settings.0.h264_settings.0.qvbr_settings.0.qvbr_quality_level": "9",
					}),
				),
			},
			{
				ResourceName: resourceName,
				ImportState:  true,
				//ImportStateIdFunc: testAccAWSMediaConvertPresetImportStateIdFunc(resourceName),
				ImportStateVerify: true,
			},
			{
				Config: testAccMediaConvertPresetConfig_BasicUpdate(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
					resource.TestCheckResourceAttr(resourceName, "settings.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs(resourceName, "settings.*", map[string]string{
						"video_description.#":                                                                     "1",
						"audio_description.#":                                                                     "1",
						"container_settings.#":                                                                    "1",
						"video_description.0.timecode_insertion":                                                  "PIC_TIMING_SEI",
						"audio_description.0.codec_settings.0.aac_settings.#":                                     "1",
						"audio_description.0.codec_settings.0.aac_settings.0.codec_profile":                       "HEV1",
						"video_description.0.codec_settings.0.h264_settings.0.quality_tuning_level":               "MULTI_PASS_HQ",
						"video_description.0.codec_settings.0.h264_settings.0.qvbr_settings.0.qvbr_quality_level": "7",
					}),
				),
			},
		},
	})
}

func TestAccAwsMediaConvertPreset_Framecaptureshort(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-framecaptureshort")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_Framecaptureshort(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_Framecapture(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-framecapture")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_Framecapture(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_Audio(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-audio")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_Audio(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_720p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-720p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_720p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_576p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-576p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_576p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_480p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-480p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_480p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_432p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-432p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_432p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_360p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-360p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_360p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_240p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-240p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_240p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func TestAccAwsMediaConvertPreset_1080p(t *testing.T) {
	var preset mediaconvert.Preset
	resourceName := "aws_media_convert_preset.test"
	rName := acctest.RandomWithPrefix("tf-acc-test-1080p")
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t); testAccPreCheckAWSMediaConvert(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckAwsMediaConvertPresetDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccMediaConvertPresetConfig_1080p(rName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckAwsMediaConvertPresetExists(resourceName, &preset),
					testAccMatchResourceAttrRegionalARN(resourceName, "arn", "mediaconvert", regexp.MustCompile(`presets/.+`)),
					resource.TestCheckResourceAttr(resourceName, "name", rName),
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

func testAccAWSMediaConvertPresetImportStateIdFunc(resourceName string) resource.ImportStateIdFunc {
	return func(s *terraform.State) (string, error) {
		rs, ok := s.RootModule().Resources[resourceName]
		if !ok {
			return "", fmt.Errorf("Not found: %s", resourceName)
		}

		return fmt.Sprintf("%s/%s", rs.Primary.Attributes["rest_api_id"], rs.Primary.ID), nil
	}
}

func testAccMediaConvertPresetConfig_Framecaptureshort(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"	
		settings {
			container_settings {
				container = "RAW"				
			}
			video_description {
				anti_alias = "ENABLED"
				codec_settings {
					codec = "FRAME_CAPTURE"
					frame_capture_settings {
						framerate_denominator = 1
						framerate_numerator = 3
						max_captures = 5
						quality = 80
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 480
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 640
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_Framecapture(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"	
		settings {
			container_settings {
				container = "RAW"				
			}
			video_description {
				anti_alias = "ENABLED"
				codec_settings {
					codec = "FRAME_CAPTURE"
					frame_capture_settings {
						framerate_denominator = 3
						framerate_numerator = 1
						max_captures = 5
						quality = 80
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 480
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 640
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_Audio(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"	
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			audio_description {
				audio_source_name = "Audio Selector 1"
				audio_type_control = "FOLLOW_INPUT"
				codec_settings {
					codec = "AAC"
					aac_settings {
						audio_description_broadcaster_mix = "NORMAL"
						bitrate = 96000
						codec_profile = "LC"
						coding_mode = "CODING_MODE_2_0"
						rate_control_mode = "CBR"						
						raw_format = "NONE"
						sample_rate = 48000
						specification = "MPEG4"						
					}
				}				
				language_code_control = "FOLLOW_INPUT"				
			}
		}	
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_720p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "720p (1280x720)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 3000000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 8
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 720
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 1280
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_576p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "576p (1024x576)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 1500000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 9
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 576
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 1024
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_480p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "480p (848x480)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 1000000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 8
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 480
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 848
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_432p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "432p (768x432)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 850000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 9
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 432
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 768
			}
		}
		
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_360p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "360p (640x360)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 700000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 9
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}					
				}
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 360
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 640
			}
		}	
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_240p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "240p (424x240)"
		settings {
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 240
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 424
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 350000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 9
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}				
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_1080p(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = "custom preset"
		description = "1080p (1920x1080)"
		settings {
			video_description {
				afd_signaling = "NONE"
				anti_alias = "ENABLED"
				color_metadata = "INSERT"
				drop_frame_timecode = "ENABLED"
				height = 1080
				respond_to_afd = "NONE"
				scaling_behavior = "DEFAULT"
				sharpness = 50
				timecode_insertion = "DISABLED"
				width = 1920
				codec_settings {
					codec = "H_264"
					h264_settings {
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						codec_profile = "MAIN"
						dynamic_sub_gop = "STATIC"
						entropy_encoding = "CABAC"
						field_encoding = "PAFF"
						flicker_adaptive_quantization = "DISABLED"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						gop_b_reference = "DISABLED"
						gop_closed_cadence = 1
						gop_size = 90.0
						gop_size_units = "FRAMES"
						interlace_mode = "PROGRESSIVE"
						max_bitrate = 6000000
						min_i_interval = 0
						number_b_frames_between_reference_frames = 2
						number_reference_frames = 3
						par_control = "INITIALIZE_FROM_SOURCE"
						quality_tuning_level = "SINGLE_PASS_HQ"
						qvbr_settings {
							qvbr_quality_level = 8
							qvbr_quality_level_fine_tune = 0.0
						}
						rate_control_mode = "QVBR"
						repeat_pps = "DISABLED"
						scene_change_detect = "ENABLED"
						slices = 1
						slow_pal = "DISABLED"
						softness = 0
						spatial_adaptive_quantization = "ENABLED"
						syntax = "DEFAULT"
						telecine = "NONE"
						temporal_adaptive_quantization = "ENABLED"
						unregistered_sei_timecode = "DISABLED"
					}
				}
			}
			container_settings {
				container = "CMFC"
				cmfc_settings {
					scte35_esam = "NONE"
					scte35_source = "NONE"
				}
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_Basic(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = ""
		settings {
			video_description {
				scaling_behavior = "DEFAULT"
				timecode_insertion = "DISABLED"
				anti_alias = "ENABLED"
				sharpness = 50
				afd_signaling = "NONE"
				drop_frame_timecode = "ENABLED"
				respond_to_afd = "NONE"
				color_metadata = "INSERT"
				codec_settings {
					codec = "H_264"
					h264_settings {
						interlace_mode = "PROGRESSIVE"
						number_reference_frames = 3
						syntax = "DEFAULT"
						softness = 0
						gop_closed_cadence = 1
						gop_size = 90
						slices = 1
						gop_b_reference = "DISABLED"
						max_bitrate = 5000000
						slow_pal = "DISABLED"
						spatial_adaptive_quantization = "ENABLED"
						temporal_adaptive_quantization = "ENABLED"
						flicker_adaptive_quantization = "DISABLED"
						entropy_encoding = "CABAC"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						rate_control_mode = "QVBR"
						qvbr_settings {
							qvbr_quality_level = 9
							qvbr_quality_level_fine_tune = 0
						}
						codec_profile = "MAIN"
						telecine = "NONE"
						min_i_interval = 0
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						field_encoding = "PAFF"
						scene_change_detect = "ENABLED"
						quality_tuning_level = "SINGLE_PASS"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						unregistered_sei_timecode = "DISABLED"
						gop_size_units = "FRAMES"
						par_control = "INITIALIZE_FROM_SOURCE"
						number_b_frames_between_reference_frames = 2
						repeat_pps = "DISABLED"
						dynamic_sub_gop = "STATIC"
					}
				}
			}
			audio_description {
				audio_type_control = "FOLLOW_INPUT"
				audio_source_name = "Audio Selector 1"
				language_code_control = "FOLLOW_INPUT"
				codec_settings {
					codec = "AAC"
					aac_settings {
						audio_description_broadcaster_mix = "NORMAL"
						bitrate = 96000
						rate_control_mode = "CBR"
						coding_mode = "CODING_MODE_2_0"
						raw_format = "NONE"
						sample_rate = 48000
						specification = "MPEG4"
						codec_profile = "LC"
					}
				}
			}		
			container_settings {
				container = "MP4"
				mp4_settings {
					cslg_atom = "INCLUDE"
					ctts_version = 0
					free_space_box = "EXCLUDE"
					moov_placement = "PROGRESSIVE_DOWNLOAD"
				}
			}
		}
	}
	`, rName)
}

func testAccMediaConvertPresetConfig_BasicUpdate(rName string) string {
	return fmt.Sprintf(`
	resource "aws_media_convert_preset" "test" {
		name = %[1]q
		category = ""
		settings {
			video_description {
				scaling_behavior = "DEFAULT"
				timecode_insertion = "PIC_TIMING_SEI"
				anti_alias = "ENABLED"
				sharpness = 50
				afd_signaling = "NONE"
				drop_frame_timecode = "ENABLED"
				respond_to_afd = "NONE"
				color_metadata = "INSERT"
				codec_settings {
					codec = "H_264"
					h264_settings {
						interlace_mode = "FOLLOW_TOP_FIELD"
						number_reference_frames = 3
						syntax = "DEFAULT"
						softness = 0
						gop_closed_cadence = 1
						gop_size = 90
						slices = 1
						gop_b_reference = "ENABLED"
						max_bitrate = 5000000
						slow_pal = "DISABLED"
						spatial_adaptive_quantization = "ENABLED"
						temporal_adaptive_quantization = "ENABLED"
						flicker_adaptive_quantization = "DISABLED"
						entropy_encoding = "CABAC"
						framerate_control = "INITIALIZE_FROM_SOURCE"
						rate_control_mode = "QVBR"
						qvbr_settings {
							qvbr_quality_level = 7
							qvbr_quality_level_fine_tune = 0
						}
						codec_profile = "MAIN"
						telecine = "NONE"
						min_i_interval = 0
						adaptive_quantization = "HIGH"
						codec_level = "AUTO"
						field_encoding = "PAFF"
						scene_change_detect = "ENABLED"
						quality_tuning_level = "MULTI_PASS_HQ"
						framerate_conversion_algorithm = "DUPLICATE_DROP"
						unregistered_sei_timecode = "DISABLED"
						gop_size_units = "FRAMES"
						par_control = "INITIALIZE_FROM_SOURCE"
						number_b_frames_between_reference_frames = 2
						repeat_pps = "DISABLED"
						dynamic_sub_gop = "STATIC"
					}
				}
			}
			audio_description {
				audio_type_control = "FOLLOW_INPUT"
				audio_source_name = "Audio Selector 1"
				language_code_control = "FOLLOW_INPUT"
				codec_settings {
					codec = "AAC"
					aac_settings {
						audio_description_broadcaster_mix = "NORMAL"
						bitrate = 96000
						rate_control_mode = "CBR"
						coding_mode = "CODING_MODE_2_0"
						raw_format = "NONE"
						sample_rate = 48000
						specification = "MPEG4"
						codec_profile = "HEV1"
					}
				}
			}		
			container_settings {
				container = "MP4"
				mp4_settings {
					cslg_atom = "INCLUDE"
					ctts_version = 0
					free_space_box = "EXCLUDE"
					moov_placement = "PROGRESSIVE_DOWNLOAD"
				}
			}
		}
	}
	`, rName)
}

func testAccCheckAwsMediaConvertPresetDestroy(s *terraform.State) error {
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "aws_media_convert_preset" {
			continue
		}
		conn, err := getAwsMediaConvertAccountClient(testAccProvider.Meta().(*AWSClient))
		if err != nil {
			return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
		}

		_, err = conn.GetPreset(&mediaconvert.GetPresetInput{
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

func testAccCheckAwsMediaConvertPresetExists(n string, preset *mediaconvert.Preset) resource.TestCheckFunc {
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

		resp, err := conn.GetPreset(&mediaconvert.GetPresetInput{
			Name: aws.String(rs.Primary.ID),
		})
		if err != nil {
			return fmt.Errorf("Error getting preset: %s", err)
		}

		*preset = *resp.Preset
		return nil
	}
}
