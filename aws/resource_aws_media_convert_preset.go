package aws

import (
	"fmt"

	"github.com/aws/aws-sdk-go/service/mediaconvert"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAwsMediaConvertPreset() *schema.Resource {
	return &schema.Resource{
		Create: resourceAwsMediaConvertPresetCreate,
		Read:   resourceAwsMediaConvertPresetRead,
		Update: resourceAwsMediaConvertPresetUpdate,
		Delete: resourceAwsMediaConvertPresetDelete,

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
			"settings": {
				Type:     schema.TypeList,
				Required: true,
				ForceNew: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"audio_description": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"audio_source_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audio_type": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"audio_type_control": {
										Type:     schema.TypeString,
										Computed: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AudioTypeControlFollowInput,
											mediaconvert.AudioTypeControlUseConfigured,
										}, false),
									},
									"custom_language_code": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"language_code": {
										Type:         schema.TypeString,
										Computed:     true,
										ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
									},
									"language_code_control": {
										Type:     schema.TypeString,
										Computed: true,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.AudioLanguageCodeControlFollowInput,
											mediaconvert.AudioLanguageCodeControlUseConfigured,
										}, false),
									},
									"stream_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"audio_channel_tagging_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channel_tag": {
													Type:     schema.TypeString,
													Required: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioChannelTagL,
														mediaconvert.AudioChannelTagR,
														mediaconvert.AudioChannelTagC,
														mediaconvert.AudioChannelTagLfe,
														mediaconvert.AudioChannelTagLs,
														mediaconvert.AudioChannelTagRs,
														mediaconvert.AudioChannelTagLc,
														mediaconvert.AudioChannelTagRc,
														mediaconvert.AudioChannelTagCs,
														mediaconvert.AudioChannelTagLsd,
														mediaconvert.AudioChannelTagRsd,
														mediaconvert.AudioChannelTagTcs,
														mediaconvert.AudioChannelTagVhl,
														mediaconvert.AudioChannelTagVhc,
														mediaconvert.AudioChannelTagVhr,
													}, false),
												},
											},
										},
									},
									"audio_normalization_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"algorithm": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationAlgorithmItuBs17701,
														mediaconvert.AudioNormalizationAlgorithmItuBs17702,
														mediaconvert.AudioNormalizationAlgorithmItuBs17703,
														mediaconvert.AudioNormalizationAlgorithmItuBs17704,
													}, false),
												},
												"algorithm_control": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationAlgorithmControlCorrectAudio,
														mediaconvert.AudioNormalizationAlgorithmControlMeasureOnly,
													}, false),
												},
												"correction_gate_level": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"loudness_logging": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationLoudnessLoggingLog,
														mediaconvert.AudioNormalizationLoudnessLoggingDontLog,
													}, false),
												},
												"peak_calculation": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioNormalizationPeakCalculationTruePeak,
														mediaconvert.AudioNormalizationPeakCalculationNone,
													}, false),
												},
												"target_lkfs": {
													Type:     schema.TypeFloat,
													Computed: true,
												},
											},
										},
									},
									"codec_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"codec": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.AudioCodecAac,
														mediaconvert.AudioCodecMp2,
														mediaconvert.AudioCodecMp3,
														mediaconvert.AudioCodecWav,
														mediaconvert.AudioCodecAiff,
														mediaconvert.AudioCodecAc3,
														mediaconvert.AudioCodecEac3,
														mediaconvert.AudioCodecEac3Atmos,
														mediaconvert.AudioCodecVorbis,
														mediaconvert.AudioCodecOpus,
														mediaconvert.AudioCodecPassthrough,
													}, false),
												},
												"aac_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"audio_description_broadcaster_mix": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacAudioDescriptionBroadcasterMixBroadcasterMixedAd,
																mediaconvert.AacAudioDescriptionBroadcasterMixNormal,
															}, false),
														},
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(6000),
														},
														"codec_profile": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacCodecProfileLc,
																mediaconvert.AacCodecProfileHev1,
																mediaconvert.AacCodecProfileHev2,
															}, false),
														},
														"coding_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacCodingModeAdReceiverMix,
																mediaconvert.AacCodingModeCodingMode10,
																mediaconvert.AacCodingModeCodingMode11,
																mediaconvert.AacCodingModeCodingMode20,
																mediaconvert.AacCodingModeCodingMode51,
															}, false),
														},
														"rate_control_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacRateControlModeCbr,
																mediaconvert.AacRateControlModeVbr,
															}, false),
														},
														"raw_format": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacRawFormatLatmLoas,
																mediaconvert.AacRawFormatNone,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(8000),
														},
														"specification": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacSpecificationMpeg2,
																mediaconvert.AacSpecificationMpeg4,
															}, false),
														},
														"vbr_quality": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.AacVbrQualityLow,
																mediaconvert.AacVbrQualityMediumLow,
																mediaconvert.AacVbrQualityMediumHigh,
																mediaconvert.AacVbrQualityHigh,
															}, false),
														},
													},
												},
												"ac3_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(64000),
														},
														"bitstream_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Ac3BitstreamModeCompleteMain,
																mediaconvert.Ac3BitstreamModeCommentary,
																mediaconvert.Ac3BitstreamModeDialogue,
																mediaconvert.Ac3BitstreamModeEmergency,
																mediaconvert.Ac3BitstreamModeHearingImpaired,
																mediaconvert.Ac3BitstreamModeMusicAndEffects,
																mediaconvert.Ac3BitstreamModeVisuallyImpaired,
																mediaconvert.Ac3BitstreamModeVoiceOver,
															}, false),
														},
														"coding_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Ac3CodingModeCodingMode10,
																mediaconvert.Ac3CodingModeCodingMode11,
																mediaconvert.Ac3CodingModeCodingMode20,
																mediaconvert.Ac3CodingModeCodingMode32Lfe,
															}, false),
														},
														"dialnorm": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"dynamic_range_compression_profile": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Ac3DynamicRangeCompressionProfileFilmStandard,
																mediaconvert.Ac3DynamicRangeCompressionProfileNone,
															}, false),
														},
														"lfe_filter": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Ac3LfeFilterEnabled,
																mediaconvert.Ac3LfeFilterDisabled,
															}, false),
														},
														"metadata_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Ac3MetadataControlFollowInput,
																mediaconvert.Ac3MetadataControlUseConfigured,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(48000),
														},
													},
												},
												"aiff_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitdepth": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(16),
														},
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(8000),
														},
													},
												},
												"eac3_atmos_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(384000),
														},
														"bitstream_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosBitstreamModeCompleteMain,
															}, false),
														},
														"coding_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosCodingModeCodingMode916,
															}, false),
														},
														"dialogue_intelligence": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosDialogueIntelligenceEnabled,
																mediaconvert.Eac3AtmosDialogueIntelligenceDisabled,
															}, false),
														},
														"dynamic_range_compression_line": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineNone,
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineFilmStandard,
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineFilmLight,
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineMusicStandard,
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineMusicLight,
																mediaconvert.Eac3AtmosDynamicRangeCompressionLineSpeech,
															}, false),
														},
														"dynamic_range_compression_rf": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfNone,
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfFilmStandard,
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfFilmLight,
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfMusicStandard,
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfMusicLight,
																mediaconvert.Eac3AtmosDynamicRangeCompressionRfSpeech,
															}, false),
														},
														"lo_ro_center_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lo_ro_surround_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lt_rt_center_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lt_rt_surround_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"metering_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosMeteringModeLeqA,
																mediaconvert.Eac3AtmosMeteringModeItuBs17701,
																mediaconvert.Eac3AtmosMeteringModeItuBs17702,
																mediaconvert.Eac3AtmosMeteringModeItuBs17703,
																mediaconvert.Eac3AtmosMeteringModeItuBs17704,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(48000),
														},
														"speech_threshold": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"stereo_downmix": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosStereoDownmixNotIndicated,
																mediaconvert.Eac3AtmosStereoDownmixStereo,
																mediaconvert.Eac3AtmosStereoDownmixSurround,
																mediaconvert.Eac3AtmosStereoDownmixDpl2,
															}, false),
														},
														"surround_ex_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AtmosSurroundExModeNotIndicated,
																mediaconvert.Eac3AtmosSurroundExModeEnabled,
																mediaconvert.Eac3AtmosSurroundExModeDisabled,
															}, false),
														},
													},
												},
												"eac3_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"attenuation_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3AttenuationControlAttenuate3Db,
																mediaconvert.Eac3AttenuationControlNone,
															}, false),
														},
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(64000),
														},
														"bitstream_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3BitstreamModeCompleteMain,
																mediaconvert.Eac3BitstreamModeCommentary,
																mediaconvert.Eac3BitstreamModeEmergency,
																mediaconvert.Eac3BitstreamModeHearingImpaired,
																mediaconvert.Eac3BitstreamModeVisuallyImpaired,
															}, false),
														},
														"coding_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3CodingModeCodingMode10,
																mediaconvert.Eac3CodingModeCodingMode20,
																mediaconvert.Eac3CodingModeCodingMode32,
															}, false),
														},
														"dc_filter": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3DcFilterEnabled,
																mediaconvert.Eac3DcFilterDisabled,
															}, false),
														},
														"dialnorm": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"dynamic_range_compression_line": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3DynamicRangeCompressionLineNone,
																mediaconvert.Eac3DynamicRangeCompressionLineFilmStandard,
																mediaconvert.Eac3DynamicRangeCompressionLineFilmLight,
																mediaconvert.Eac3DynamicRangeCompressionLineMusicStandard,
																mediaconvert.Eac3DynamicRangeCompressionLineMusicLight,
																mediaconvert.Eac3DynamicRangeCompressionLineSpeech,
															}, false),
														},
														"dynamic_range_compression_rf": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3DynamicRangeCompressionRfNone,
																mediaconvert.Eac3DynamicRangeCompressionRfFilmStandard,
																mediaconvert.Eac3DynamicRangeCompressionRfFilmLight,
																mediaconvert.Eac3DynamicRangeCompressionRfMusicStandard,
																mediaconvert.Eac3DynamicRangeCompressionRfMusicLight,
																mediaconvert.Eac3DynamicRangeCompressionRfSpeech,
															}, false),
														},
														"lfe_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3LfeControlLfe,
																mediaconvert.Eac3LfeControlNoLfe,
															}, false),
														},
														"lfe_filter": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3LfeFilterEnabled,
																mediaconvert.Eac3LfeFilterDisabled,
															}, false),
														},
														"lo_ro_center_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lo_ro_surround_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lt_rt_center_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"lt_rt_surround_mix_level": {
															Type:     schema.TypeFloat,
															Computed: true,
														},
														"metadata_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3MetadataControlFollowInput,
																mediaconvert.Eac3MetadataControlUseConfigured,
															}, false),
														},
														"passthrough_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3PassthroughControlWhenPossible,
																mediaconvert.Eac3PassthroughControlNoPassthrough,
															}, false),
														},
														"phase_control": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3PhaseControlShift90Degrees,
																mediaconvert.Eac3PhaseControlNoShift,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(48000),
														},
														"stereo_downmix": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3StereoDownmixNotIndicated,
																mediaconvert.Eac3StereoDownmixLoRo,
																mediaconvert.Eac3StereoDownmixLtRt,
																mediaconvert.Eac3StereoDownmixDpl2,
															}, false),
														},
														"surround_ex_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3SurroundExModeNotIndicated,
																mediaconvert.Eac3SurroundExModeEnabled,
																mediaconvert.Eac3SurroundExModeDisabled,
															}, false),
														},
														"surround_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Eac3SurroundModeNotIndicated,
																mediaconvert.Eac3SurroundModeEnabled,
																mediaconvert.Eac3SurroundModeDisabled,
															}, false),
														},
													},
												},
												"mp2_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(32000),
														},
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(32000),
														},
													},
												},
												"mp3_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(16000),
														},
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"rate_control_mode": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.Mp3RateControlModeCbr,
																mediaconvert.Mp3RateControlModeVbr,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(22050),
														},
														"vbr_quality": {
															Type:     schema.TypeInt,
															Computed: true,
														},
													},
												},
												"opus_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitrate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(32000),
														},
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(16000),
														},
													},
												},
												"vorbis_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(22050),
														},
														"vbr_quality": {
															Type:     schema.TypeInt,
															Computed: true,
														},
													},
												},
												"wav_settings": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: map[string]*schema.Schema{
														"bitdepth": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(16),
														},
														"channels": {
															Type:         schema.TypeInt,
															Computed:     true,
															ValidateFunc: validation.IntAtLeast(1),
														},
														"format": {
															Type:     schema.TypeString,
															Computed: true,
															ValidateFunc: validation.StringInSlice([]string{
																mediaconvert.WavFormatRiff,
																mediaconvert.WavFormatRf64,
															}, false),
														},
														"sample_rate": {
															Type:         schema.TypeInt,
															ValidateFunc: validation.IntAtLeast(8000),
														},
													},
												},
											},
										},
									},
									"remix_settings": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"channel_mapping": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"output_channels": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"input_channels": {
																			Type:     schema.TypeList,
																			Computed: true,
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
													Computed:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
												"channels_out": {
													Type:         schema.TypeInt,
													Computed:     true,
													ValidateFunc: validation.IntAtLeast(1),
												},
											},
										},
									},
								},
							},
						},
						"caption_description": {
							Type:     schema.TypeList,
							Required: true,
							ForceNew: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"custom_language_code": {
										Type: schema.TypeString,
									},
									"destination_settings": {
										Type:     schema.TypeList,
										Optional: true,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"burnin_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleAlignmentCentered,
																	mediaconvert.BurninSubtitleAlignmentLeft,
																}, false),
															},
															"background_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleBackgroundColorNone,
																	mediaconvert.BurninSubtitleBackgroundColorBlack,
																	mediaconvert.BurninSubtitleBackgroundColorWhite,
																}, false),
															},
															"background_opacity": {
																Type: schema.TypeInt,
															},
															"font_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleFontColorWhite,
																	mediaconvert.BurninSubtitleFontColorBlack,
																	mediaconvert.BurninSubtitleFontColorYellow,
																	mediaconvert.BurninSubtitleFontColorRed,
																	mediaconvert.BurninSubtitleFontColorGreen,
																	mediaconvert.BurninSubtitleFontColorBlue,
																}, false),
															},
															"font_opacity": {
																Type: schema.TypeInt,
															},
															"font_resolution": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(96),
															},
															"font_script": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.FontScriptAutomatic,
																	mediaconvert.FontScriptHans,
																	mediaconvert.FontScriptHant,
																}, false),
															},
															"font_size": {
																Type: schema.TypeInt,
															},
															"outline_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleOutlineColorBlack,
																	mediaconvert.BurninSubtitleOutlineColorWhite,
																	mediaconvert.BurninSubtitleOutlineColorYellow,
																	mediaconvert.BurninSubtitleOutlineColorRed,
																	mediaconvert.BurninSubtitleOutlineColorGreen,
																	mediaconvert.BurninSubtitleOutlineColorBlue,
																}, false),
															},
															"outline_size": {
																Type: schema.TypeInt,
															},
															"shadow_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleShadowColorNone,
																	mediaconvert.BurninSubtitleShadowColorBlack,
																	mediaconvert.BurninSubtitleShadowColorWhite,
																}, false),
															},
															"shadow_opacity": {
																Type: schema.TypeInt,
															},
															"shadow_x_offset": {
																Type: schema.TypeInt,
															},
															"shadow_y_offset": {
																Type: schema.TypeInt,
															},
															"teletext_spacing": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.BurninSubtitleTeletextSpacingFixedGrid,
																	mediaconvert.BurninSubtitleTeletextSpacingProportional,
																}, false),
															},
															"x_position": {
																Type: schema.TypeInt,
															},
															"y_position": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"destination_type": {
													Type:     schema.TypeString,
													Computed: true,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CaptionDestinationTypeBurnIn,
														mediaconvert.CaptionDestinationTypeDvbSub,
														mediaconvert.CaptionDestinationTypeEmbedded,
														mediaconvert.CaptionDestinationTypeEmbeddedPlusScte20,
														mediaconvert.CaptionDestinationTypeImsc,
														mediaconvert.CaptionDestinationTypeScte20PlusEmbedded,
														mediaconvert.CaptionDestinationTypeScc,
														mediaconvert.CaptionDestinationTypeSrt,
														mediaconvert.CaptionDestinationTypeSmi,
														mediaconvert.CaptionDestinationTypeTeletext,
														mediaconvert.CaptionDestinationTypeTtml,
														mediaconvert.CaptionDestinationTypeWebvtt,
													}, false),
												},
												"dvb_sub_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"alignment": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleAlignmentCentered,
																	mediaconvert.DvbSubtitleAlignmentLeft,
																}, false),
															},
															"background_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleBackgroundColorNone,
																	mediaconvert.DvbSubtitleBackgroundColorBlack,
																	mediaconvert.DvbSubtitleBackgroundColorWhite,
																}, false),
															},
															"background_opacity": {
																Type: schema.TypeInt,
															},
															"font_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleFontColorWhite,
																	mediaconvert.DvbSubtitleFontColorBlack,
																	mediaconvert.DvbSubtitleFontColorYellow,
																	mediaconvert.DvbSubtitleFontColorRed,
																	mediaconvert.DvbSubtitleFontColorGreen,
																	mediaconvert.DvbSubtitleFontColorBlue,
																}, false),
															},
															"font_opacity": {
																Type: schema.TypeInt,
															},
															"font_resolution": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(96),
															},
															"font_script": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.FontScriptAutomatic,
																	mediaconvert.FontScriptHans,
																	mediaconvert.FontScriptHant,
																}, false),
															},
															"font_size": {
																Type: schema.TypeInt,
															},
															"outline_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleOutlineColorBlack,
																	mediaconvert.DvbSubtitleOutlineColorWhite,
																	mediaconvert.DvbSubtitleOutlineColorYellow,
																	mediaconvert.DvbSubtitleOutlineColorRed,
																	mediaconvert.DvbSubtitleOutlineColorGreen,
																	mediaconvert.DvbSubtitleOutlineColorBlue,
																}, false),
															},
															"outline_size": {
																Type: schema.TypeInt,
															},
															"shadow_color": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleShadowColorNone,
																	mediaconvert.DvbSubtitleShadowColorBlack,
																	mediaconvert.DvbSubtitleShadowColorWhite,
																}, false),
															},
															"shadow_opacity": {
																Type: schema.TypeInt,
															},
															"shadow_x_offset": {
																Type: schema.TypeInt,
															},
															"shadow_y_offset": {
																Type: schema.TypeInt,
															},
															"subtitling_type": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitlingTypeHearingImpaired,
																	mediaconvert.DvbSubtitlingTypeStandard,
																}, false),
															},
															"teletext_spacing": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.DvbSubtitleTeletextSpacingFixedGrid,
																	mediaconvert.DvbSubtitleTeletextSpacingProportional,
																}, false),
															},
															"x_position": {
																Type: schema.TypeInt,
															},
															"y_position": {
																Type: schema.TypeInt,
															},
														},
													},
												},
												"embedded_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"destination_608_channel_number": {
																Type:         schema.TypeInt,
																Computed:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
															"destination_708_service_number": {
																Type:         schema.TypeInt,
																Computed:     true,
																ValidateFunc: validation.IntAtLeast(1),
															},
														},
													},
												},
												"imsc_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_passthrough": {
																Type:     schema.TypeString,
																Computed: true,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.ImscStylePassthroughEnabled,
																	mediaconvert.ImscStylePassthroughDisabled,
																}, false),
															},
														},
													},
												},
												"scc_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"framerate": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.SccDestinationFramerateFramerate2397,
																	mediaconvert.SccDestinationFramerateFramerate24,
																	mediaconvert.SccDestinationFramerateFramerate25,
																	mediaconvert.SccDestinationFramerateFramerate2997Dropframe,
																	mediaconvert.SccDestinationFramerateFramerate2997NonDropframe,
																}, false),
															},
														},
													},
												},
												"teletext_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"page_number": {
																Type:         schema.TypeString,
																ValidateFunc: validation.StringLenBetween(3, 256),
															},
															"page_types": {
																Type: schema.TypeSet,
																Elem: &schema.Schema{Type: schema.TypeString},
																Set:  schema.HashString,
															},
														},
													},
												},
												"ttml_destination_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"style_passthrough": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.TtmlStylePassthroughEnabled,
																	mediaconvert.TtmlStylePassthroughDisabled,
																}, false),
															},
														},
													},
												},
											},
										},
									},
									"language_code": {
										Type:         schema.TypeString,
										ValidateFunc: validation.StringInSlice(mediaconvert.LanguageCode_Values(), false),
									},
									"language_description": {
										Type: schema.TypeString,
									},
								},
							},
						},
						"container_settings": {
							Type:     schema.TypeList,
							Required: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cmfc_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcAudioDurationDefaultCodecDuration,
														mediaconvert.CmfcAudioDurationMatchVideoDuration,
													}, false),
												},
												"scte35_esam": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcScte35EsamInsert,
														mediaconvert.CmfcScte35EsamNone,
													}, false),
												},
												"scte35_source ": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcScte35SourcePassthrough,
														mediaconvert.CmfcScte35SourceNone,
													}, false),
												},
											},
										},
									},
									"container": {
										Type: schema.TypeString,
										ValidateFunc: validation.StringInSlice([]string{
											mediaconvert.ContainerTypeF4v,
											mediaconvert.ContainerTypeIsmv,
											mediaconvert.ContainerTypeM2ts,
											mediaconvert.ContainerTypeM3u8,
											mediaconvert.ContainerTypeCmfc,
											mediaconvert.ContainerTypeMov,
											mediaconvert.ContainerTypeMp4,
											mediaconvert.ContainerTypeMpd,
											mediaconvert.ContainerTypeMxf,
											mediaconvert.ContainerTypeWebm,
											mediaconvert.ContainerTypeRaw,
										}, false),
									},
									"f4v_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"moov_placement": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.F4vMoovPlacementProgressiveDownload,
														mediaconvert.F4vMoovPlacementNormal,
													}, false),
												},
											},
										},
									},
									"m2ts_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsAudioDurationDefaultCodecDuration,
														mediaconvert.M2tsAudioDurationMatchVideoDuration,
													}, false),
												},
												"audio_frames_per_pes ": {
													Type: schema.TypeInt,
												},
												"audio_pids": {
													Type: schema.TypeSet,
													Elem: &schema.Schema{Type: schema.TypeInt},
													Set:  schema.HashString,
												},
												"bitrate ": {
													Type: schema.TypeInt,
												},
												"buffer_model ": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsBufferModelMultiplex,
														mediaconvert.M2tsBufferModelNone,
													}, false),
												},
												"dvb_nit_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"network_id": {
																Type: schema.TypeInt,
															},
															"network_name": {
																Type:         schema.TypeString,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
															"nit_interval": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(25),
															},
														},
													},
												},
												"dvb_sdt_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"output_sdt": {
																Type: schema.TypeString,
																ValidateFunc: validation.StringInSlice([]string{
																	mediaconvert.OutputSdtSdtFollow,
																	mediaconvert.OutputSdtSdtFollowIfPresent,
																	mediaconvert.OutputSdtSdtManual,
																	mediaconvert.OutputSdtSdtNone,
																}, false),
															},
															"sdt_interval": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(25),
															},
															"service_name": {
																Type:         schema.TypeString,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
															"service_provider_name": {
																Type:         schema.TypeString,
																ValidateFunc: validation.StringLenBetween(1, 256),
															},
														},
													},
												},
												"dvb_sub_pids": {
													Type: schema.TypeSet,
													Elem: &schema.Schema{Type: schema.TypeInt},
													Set:  schema.HashString,
												},
												"dvb_tdt_settings": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"tdt_interval": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(1000),
															},
														},
													},
												},
												"dvb_teletext_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      499,
												},
												"ebp_audio_interval": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEbpAudioIntervalVideoAndFixedIntervals,
														mediaconvert.M2tsEbpAudioIntervalVideoInterval,
													}, false),
												},
												"ebp_placement": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEbpPlacementVideoAndAudioPids,
														mediaconvert.M2tsEbpPlacementVideoPid,
													}, false),
												},
												"es_rate_in_pes": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsEsRateInPesInclude,
														mediaconvert.M2tsEsRateInPesExclude,
													}, false),
												},
												"force_ts_video_ebp_order": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsForceTsVideoEbpOrderForce,
														mediaconvert.M2tsForceTsVideoEbpOrderDefault,
													}, false),
												},
												"fragment_time": {
													Type: schema.TypeFloat,
												},
												"max_pcr_interval": {
													Type: schema.TypeInt,
												},
												"min_pcr_interval": {
													Type: schema.TypeInt,
												},
												"nielsen_id3": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsNielsenId3Insert,
														mediaconvert.M2tsNielsenId3None,
													}, false),
												},
												"null_packet_bitrate": {
													Type: schema.TypeFloat,
												},
												"pat_interval": {
													Type: schema.TypeInt,
												},
												"pcr_control": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsPcrControlPcrEveryPesPacket,
														mediaconvert.M2tsPcrControlConfiguredPcrPeriod,
													}, false),
												},
												"pcr_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"pmt_interval": {
													Type: schema.TypeInt,
												},
												"pmt_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      48,
												},
												"private_metadata_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
													Default:      503,
												},
												"program_number": {
													Type:    schema.TypeInt,
													Default: 1,
												},
												"rate_mode": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsRateModeVbr,
														mediaconvert.M2tsRateModeCbr,
													}, false),
												},
												"scte_35_esam": {
													Type:     schema.TypeList,
													MaxItems: 1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"scte_35_esam_pid": {
																Type:         schema.TypeInt,
																ValidateFunc: validation.IntAtLeast(32),
															},
														},
													},
												},
												"scte_35_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"scte_35_source": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsScte35SourcePassthrough,
														mediaconvert.M2tsScte35SourceNone,
													}, false),
												},
												"segmentation_markers": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsSegmentationMarkersNone,
														mediaconvert.M2tsSegmentationMarkersRaiSegstart,
														mediaconvert.M2tsSegmentationMarkersRaiAdapt,
														mediaconvert.M2tsSegmentationMarkersPsiSegstart,
														mediaconvert.M2tsSegmentationMarkersEbp,
														mediaconvert.M2tsSegmentationMarkersEbpLegacy,
													}, false),
												},
												"segmentation_style": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M2tsSegmentationStyleMaintainCadence,
														mediaconvert.M2tsSegmentationStyleResetCadence,
													}, false),
												},
												"segmentation_time": {
													Type: schema.TypeFloat,
												},
												"timed_metadata_pid": {
													Type:    schema.TypeInt,
													Default: 32,
												},
												"transport_stream_id": {
													Type: schema.TypeInt,
												},
												"video_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
											},
										},
									},
									"m3u8_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8AudioDurationDefaultCodecDuration,
														mediaconvert.M3u8AudioDurationMatchVideoDuration,
													}, false)},
												"audio_frames_per_pes": {
													Type: schema.TypeInt,
												},
												"audio_pids": {
													Type: schema.TypeSet,
													Elem: &schema.Schema{Type: schema.TypeInt},
													Set:  schema.HashString,
												},
												"nielsen_id3": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8NielsenId3Insert,
														mediaconvert.M3u8NielsenId3None,
													}, false),
												},
												"pat_interval": {
													Type: schema.TypeInt,
												},
												"pcr_control": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8PcrControlPcrEveryPesPacket,
														mediaconvert.M3u8PcrControlConfiguredPcrPeriod,
													}, false),
												},
												"pcr_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"pmt_interval": {
													Type: schema.TypeInt,
												},
												"pmt_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"private_metadata_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"program_number": {
													Type: schema.TypeInt,
												},
												"scte_35_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"scte_35_source": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.M3u8Scte35SourcePassthrough,
														mediaconvert.M3u8Scte35SourceNone,
													}, false),
												},
												"timed_metadata": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.TimedMetadataPassthrough,
														mediaconvert.TimedMetadataNone,
													}, false),
												},
												"timed_metadata_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
												"transport_stream_id": {
													Type: schema.TypeInt,
												},
												"video_pid": {
													Type:         schema.TypeInt,
													ValidateFunc: validation.IntAtLeast(32),
												},
											},
										},
									},
									"mov_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"clap_atom": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovClapAtomInclude,
														mediaconvert.MovClapAtomExclude,
													}, false),
												},
												"cslg_atom": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovCslgAtomInclude,
														mediaconvert.MovCslgAtomExclude,
													}, false),
												},
												"mpeg2_fourcc_control": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovMpeg2FourCCControlXdcam,
														mediaconvert.MovMpeg2FourCCControlMpeg,
													}, false),
												},
												"padding_control": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovPaddingControlOmneon,
														mediaconvert.MovPaddingControlNone,
													}, false),
												},
												"reference": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MovReferenceSelfContained,
														mediaconvert.MovReferenceExternal,
													}, false),
												},
											},
										},
									},
									"mp4_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"audio_duration": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.CmfcAudioDurationDefaultCodecDuration,
														mediaconvert.CmfcAudioDurationMatchVideoDuration,
													}, false),
												},
												"cslg_atom": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4CslgAtomInclude,
														mediaconvert.Mp4CslgAtomExclude,
													}, false),
												},
												"ctts_version ": {
													Type:    schema.TypeInt,
													Default: 0,
												},
												"free_space_box": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4FreeSpaceBoxInclude,
														mediaconvert.Mp4FreeSpaceBoxExclude,
													}, false),
												},
												"moov_placement": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.Mp4MoovPlacementProgressiveDownload,
														mediaconvert.Mp4MoovPlacementNormal,
													}, false),
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
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"accessibility_caption_hints ": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdAccessibilityCaptionHintsInclude,
														mediaconvert.MpdAccessibilityCaptionHintsExclude,
													}, false),
												},
												"audio_duration": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdAudioDurationDefaultCodecDuration,
														mediaconvert.MpdAudioDurationMatchVideoDuration,
													}, false)},
												"caption_container_type": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdCaptionContainerTypeRaw,
														mediaconvert.MpdCaptionContainerTypeFragmentedMp4,
													}, false),
												},
												"scte_35_esam": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdScte35EsamInsert,
														mediaconvert.MpdScte35EsamNone,
													}, false),
												},
												"scte_35_source": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MpdScte35SourcePassthrough,
														mediaconvert.MpdScte35SourceNone,
													}, false),
												},
											},
										},
									},
									"mxf_settings": {
										Type:     schema.TypeList,
										MaxItems: 1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"afd_signaling": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MxfAfdSignalingNoCopy,
														mediaconvert.MxfAfdSignalingCopyFromVideo,
													}, false),
												},
												"profile": {
													Type: schema.TypeString,
													ValidateFunc: validation.StringInSlice([]string{
														mediaconvert.MxfProfileD10,
														mediaconvert.MxfProfileXdcam,
														mediaconvert.MxfProfileOp1a,
													}, false),
												},
											},
										},
									},
								},
							},
						},
						//https://docs.aws.amazon.com/sdk-for-go/api/service/mediaconvert/#VideoDescription
						"video_description": {
							Type:     schema.TypeList,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
					},
				},
			},
			"tags": tagsSchema(),
		},
	}
}

func resourceAwsMediaConvertPresetCreate(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	// createOpts := &mediaconvert.CreatePresetInput{
	// 	Category:    aws.String(d.Get("category").(string)),
	// 	Description: aws.String(d.Get("description").(string)),
	// 	Name:        aws.String(d.Get("name").(string)),
	// 	//Settings
	// 	Tags: keyvaluetags.New(d.Get("tags").(map[string]interface{})).IgnoreAws().MediaconvertTags(),
	// }
	return resourceAwsMediaConvertPresetRead(d, meta)
}

func resourceAwsMediaConvertPresetRead(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}

	//ignoreTagsConfig := meta.(*AWSClient).IgnoreTagsConfig
	return nil
}

func resourceAwsMediaConvertPresetUpdate(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	return resourceAwsMediaConvertQueueRead(d, meta)
}

func resourceAwsMediaConvertPresetDelete(d *schema.ResourceData, meta interface{}) error {
	_, err := getAwsMediaConvertAccountClient(meta.(*AWSClient))
	if err != nil {
		return fmt.Errorf("Error getting Media Convert Account Client: %s", err)
	}
	return nil
}
