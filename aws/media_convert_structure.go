package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/mediaconvert"
)

func expandMediaConvertReservationPlanSettings(config map[string]interface{}) *mediaconvert.ReservationPlanSettings {
	reservationPlanSettings := &mediaconvert.ReservationPlanSettings{}

	if v, ok := config["commitment"]; ok {
		reservationPlanSettings.Commitment = aws.String(v.(string))
	}

	if v, ok := config["renewal_type"]; ok {
		reservationPlanSettings.RenewalType = aws.String(v.(string))
	}

	if v, ok := config["reserved_slots"]; ok {
		reservationPlanSettings.ReservedSlots = aws.Int64(int64(v.(int)))
	}

	return reservationPlanSettings
}

func flattenMediaConvertReservationPlan(reservationPlan *mediaconvert.ReservationPlan) []interface{} {
	if reservationPlan == nil {
		return []interface{}{}
	}

	m := map[string]interface{}{
		"commitment":     aws.StringValue(reservationPlan.Commitment),
		"renewal_type":   aws.StringValue(reservationPlan.RenewalType),
		"reserved_slots": aws.Int64Value(reservationPlan.ReservedSlots),
	}

	return []interface{}{m}
}

func expandMediaConvertNexGuardFileMarkerSettings(list []interface{}) *mediaconvert.NexGuardFileMarkerSettings {
	result := &mediaconvert.NexGuardFileMarkerSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["license"].(string); ok && v != "" {
		result.License = aws.String(v)
	}
	if v, ok := tfMap["payload"].(int); ok {
		result.Payload = aws.Int64(int64(v))
	}
	if v, ok := tfMap["preset"].(string); ok && v != "" {
		result.Preset = aws.String(v)
	}
	if v, ok := tfMap["strength"].(string); ok && v != "" {
		result.Strength = aws.String(v)
	}
	return result
}

func expandMediaConvertDolbyVisionLevel6Metadata(list []interface{}) *mediaconvert.DolbyVisionLevel6Metadata {
	result := &mediaconvert.DolbyVisionLevel6Metadata{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_cll"].(int); ok {
		result.MaxCll = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_fall"].(int); ok {
		result.MaxFall = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerTemporalFilterSettings(list []interface{}) *mediaconvert.NoiseReducerTemporalFilterSettings {
	result := &mediaconvert.NoiseReducerTemporalFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["aggressive_mode"].(int); ok {
		result.AggressiveMode = aws.Int64(int64(v))
	}
	if v, ok := tfMap["post_temporal_sharpening"].(string); ok && v != "" {
		result.PostTemporalSharpening = aws.String(v)
	}
	if v, ok := tfMap["speed"].(int); ok {
		result.Speed = aws.Int64(int64(v))
	}
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerFilterSettings(list []interface{}) *mediaconvert.NoiseReducerFilterSettings {
	result := &mediaconvert.NoiseReducerFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertNoiseReducerSpatialFilterSettings(list []interface{}) *mediaconvert.NoiseReducerSpatialFilterSettings {
	result := &mediaconvert.NoiseReducerSpatialFilterSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["post_filter_sharpen_strength"].(int); ok {
		result.PostFilterSharpenStrength = aws.Int64(int64(v))
	}
	if v, ok := tfMap["speed"].(int); ok {
		result.Speed = aws.Int64(int64(v))
	}
	if v, ok := tfMap["strength"].(int); ok {
		result.Strength = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertVideoPreprocessor(list []interface{}) *mediaconvert.VideoPreprocessor {
	result := &mediaconvert.VideoPreprocessor{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["color_corrector"]; ok {
		result.ColorCorrector = expandMediaConvertColorCorrector(v.([]interface{}))
	}
	if v, ok := tfMap["deinterlacer"]; ok {
		result.Deinterlacer = expandMediaConvertDeinterlacer(v.([]interface{}))
	}
	if v, ok := tfMap["dolby_vision"]; ok {
		result.DolbyVision = expandMediaConvertDolbyVision(v.([]interface{}))
	}
	if v, ok := tfMap["image_inserter"]; ok {
		result.ImageInserter = expandMediaConvertImageInserter(v.([]interface{}))
	}
	if v, ok := tfMap["image_inserter"]; ok {
		result.ImageInserter = expandMediaConvertImageInserter(v.([]interface{}))
	}
	if v, ok := tfMap["noise_reducer"]; ok {
		result.NoiseReducer = expandMediaConvertNoiseReducer(v.([]interface{}))
	}
	if v, ok := tfMap["partner_watermaking"]; ok {
		result.PartnerWatermarking = expandMediaConvertPartnerWatermarking(v.([]interface{}))
	}
	if v, ok := tfMap["timecode_burnin"]; ok {
		result.TimecodeBurnin = expandMediaConvertTimecodeBurnin(v.([]interface{}))
	}
	return result
}

func expandMediaConvertNoiseReducer(list []interface{}) *mediaconvert.NoiseReducer {
	result := &mediaconvert.NoiseReducer{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["filter"].(string); ok && v != "" {
		result.Filter = aws.String(v)
	}
	if v, ok := tfMap["filter_settings"]; ok {
		result.FilterSettings = expandMediaConvertNoiseReducerFilterSettings(v.([]interface{}))
	}
	if v, ok := tfMap["spatial_filter_settings"]; ok {
		result.SpatialFilterSettings = expandMediaConvertNoiseReducerSpatialFilterSettings(v.([]interface{}))
	}
	if v, ok := tfMap["temporal_filter_settings"]; ok {
		result.TemporalFilterSettings = expandMediaConvertNoiseReducerTemporalFilterSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertPartnerWatermarking(list []interface{}) *mediaconvert.PartnerWatermarking {
	result := &mediaconvert.PartnerWatermarking{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["nexguard_file_marker_settings"]; ok {
		result.NexguardFileMarkerSettings = expandMediaConvertNexGuardFileMarkerSettings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertTimecodeBurnin(list []interface{}) *mediaconvert.TimecodeBurnin {
	result := &mediaconvert.TimecodeBurnin{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["font_size"].(int); ok {
		result.FontSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["position"].(string); ok && v != "" {
		result.Position = aws.String(v)
	}
	if v, ok := tfMap["prefix"].(string); ok && v != "" {
		result.Prefix = aws.String(v)
	}
	return result
}

func expandMediaConvertDolbyVision(list []interface{}) *mediaconvert.DolbyVision {
	result := &mediaconvert.DolbyVision{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})

	if v, ok := tfMap["l6_metadata"]; ok {
		result.L6Metadata = expandMediaConvertDolbyVisionLevel6Metadata(v.([]interface{}))
	}
	if v, ok := tfMap["l6_mode"].(string); ok && v != "" {
		result.L6Mode = aws.String(v)
	}
	if v, ok := tfMap["profile"].(string); ok && v != "" {
		result.Profile = aws.String(v)
	}
	return result
}

func expandMediaConvertDeinterlacer(list []interface{}) *mediaconvert.Deinterlacer {
	result := &mediaconvert.Deinterlacer{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["algorithm"].(string); ok && v != "" {
		result.Algorithm = aws.String(v)
	}
	if v, ok := tfMap["control"].(string); ok && v != "" {
		result.Control = aws.String(v)
	}
	if v, ok := tfMap["mode"].(string); ok && v != "" {
		result.Mode = aws.String(v)
	}
	return result
}

func expandMediaConvertColorCorrector(list []interface{}) *mediaconvert.ColorCorrector {
	result := &mediaconvert.ColorCorrector{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["brightness"].(int); ok {
		result.Brightness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["color_space_conversion"].(string); ok && v != "" {
		result.ColorSpaceConversion = aws.String(v)
	}
	if v, ok := tfMap["contrast"].(int); ok {
		result.Contrast = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hdr10_metadata"]; ok {
		result.Hdr10Metadata = expandMediaConvertHdr10Metadata(v.([]interface{}))
	}
	if v, ok := tfMap["hue"].(int); ok {
		result.Hue = aws.Int64(int64(v))
	}
	if v, ok := tfMap["saturation"].(int); ok {
		result.Saturation = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertVideoCodecSettings(list []interface{}) *mediaconvert.VideoCodecSettings {
	result := &mediaconvert.VideoCodecSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["av1_settings"]; ok {
		result.Av1Settings = expandMediaConvertAv1Settings(v.([]interface{}))
	}
	if v, ok := tfMap["avc_intra_settings"]; ok {
		result.AvcIntraSettings = expandMediaConvertAvcIntraSettings(v.([]interface{}))
	}
	if v, ok := tfMap["codec"].(string); ok && v != "" {
		result.Codec = aws.String(v)
	}
	if v, ok := tfMap["frame_capture_settings"]; ok {
		result.FrameCaptureSettings = expandMediaConvertFrameCaptureSettings(v.([]interface{}))
	}
	if v, ok := tfMap["h264_settings"]; ok {
		result.H264Settings = expandMediaConvertH264Settings(v.([]interface{}))
	}
	if v, ok := tfMap["h265_settings"]; ok {
		result.H265Settings = expandMediaConvertH265Settings(v.([]interface{}))
	}
	if v, ok := tfMap["mpeg2_settings"]; ok {
		result.Mpeg2Settings = expandMediaConvertMpeg2Settings(v.([]interface{}))
	}
	if v, ok := tfMap["prores_settings "]; ok {
		result.ProresSettings = expandMediaConvertProresSettings(v.([]interface{}))
	}
	if v, ok := tfMap["vc3_settings"]; ok {
		result.Vc3Settings = expandMediaConvertVc3Settings(v.([]interface{}))
	}
	if v, ok := tfMap["vp8_settings"]; ok {
		result.Vp8Settings = expandMediaConvertVp8Settings(v.([]interface{}))
	}
	if v, ok := tfMap["vp9_settings"]; ok {
		result.Vp9Settings = expandMediaConvertVp9Settings(v.([]interface{}))
	}
	return result
}

func expandMediaConvertVideoDescription(list []interface{}) *mediaconvert.VideoDescription {
	result := &mediaconvert.VideoDescription{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["afd_signaling"].(string); ok && v != "" {
		result.AfdSignaling = aws.String(v)
	}
	if v, ok := tfMap["anti_alias"].(string); ok && v != "" {
		result.AntiAlias = aws.String(v)
	}
	if v, ok := tfMap["codec_settings"]; ok {
		result.CodecSettings = expandMediaConvertVideoCodecSettings(v.([]interface{}))
	}
	if v, ok := tfMap["color_metadata"].(string); ok && v != "" {
		result.ColorMetadata = aws.String(v)
	}
	if v, ok := tfMap["crop"]; ok {
		result.Crop = expandMediaConvertRectangle(v.([]interface{}))
	}
	if v, ok := tfMap["drop_frame_timecode"].(string); ok && v != "" {
		result.DropFrameTimecode = aws.String(v)
	}
	if v, ok := tfMap["fixed_afd"].(int); ok && v != 0 {
		result.FixedAfd = aws.Int64(int64(v))
	}
	if v, ok := tfMap["height"].(int); ok && v != 0 {
		result.Height = aws.Int64(int64(v))
	}
	if v, ok := tfMap["position"]; ok {
		result.Position = expandMediaConvertRectangle(v.([]interface{}))
	}
	if v, ok := tfMap["respond_to_afd"].(string); ok && v != "" {
		result.RespondToAfd = aws.String(v)
	}
	if v, ok := tfMap["scaling_behavior"].(string); ok && v != "" {
		result.ScalingBehavior = aws.String(v)
	}
	if v, ok := tfMap["sharpness"].(int); ok {
		result.Sharpness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["timecode_insertion"].(string); ok && v != "" {
		result.TimecodeInsertion = aws.String(v)
	}
	if v, ok := tfMap["video_preprocessors"]; ok {
		result.VideoPreprocessors = expandMediaConvertVideoPreprocessor(v.([]interface{}))
	}
	if v, ok := tfMap["width"].(int); ok && v != 0 {
		result.Width = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertAv1Settings(list []interface{}) *mediaconvert.Av1Settings {
	result := &mediaconvert.Av1Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertAv1QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	return result
}

func expandMediaConvertAvcIntraSettings(list []interface{}) *mediaconvert.AvcIntraSettings {
	result := &mediaconvert.AvcIntraSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["avc_intra_class"].(string); ok && v != "" {
		result.AvcIntraClass = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	return result
}

func expandMediaConvertFrameCaptureSettings(list []interface{}) *mediaconvert.FrameCaptureSettings {
	result := &mediaconvert.FrameCaptureSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_captures"].(int); ok {
		result.MaxCaptures = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality"].(int); ok {
		result.Quality = aws.Int64(int64(v))
	}
	return result
}

func expandMediaConvertH264Settings(list []interface{}) *mediaconvert.H264Settings {
	result := &mediaconvert.H264Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok && v != 0 {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["entropy_encoding"].(string); ok && v != "" {
		result.EntropyEncoding = aws.String(v)
	}
	if v, ok := tfMap["field_encoding"].(string); ok && v != "" {
		result.FieldEncoding = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_b_reference"].(string); ok && v != "" {
		result.GopBReference = aws.String(v)
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_reference_frames"].(int); ok {
		result.NumberReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertH264QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["repeat_pps"].(string); ok && v != "" {
		result.RepeatPps = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["softness"].(int); ok {
		result.Softness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["syntax"].(string); ok && v != "" {
		result.Syntax = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["unregistered_sei_timecode"].(string); ok && v != "" {
		result.UnregisteredSeiTimecode = aws.String(v)
	}
	return result
}

func expandMediaConvertH265Settings(list []interface{}) *mediaconvert.H265Settings {
	result := &mediaconvert.H265Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["alternate_transfer_function_sei"].(string); ok && v != "" {
		result.AlternateTransferFunctionSei = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["flicker_adaptive_quantization"].(string); ok && v != "" {
		result.FlickerAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_b_reference"].(string); ok && v != "" {
		result.GopBReference = aws.String(v)
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_reference_frames"].(int); ok {
		result.NumberReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["qvbr_settings"]; ok {
		result.QvbrSettings = expandMediaConvertH265QvbrSettings(v.([]interface{}))
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["sample_adaptive_offset_filter_mode"].(string); ok && v != "" {
		result.SampleAdaptiveOffsetFilterMode = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slices"].(int); ok {
		result.Slices = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["temporal_ids"].(string); ok && v != "" {
		result.TemporalIds = aws.String(v)
	}
	if v, ok := tfMap["tiles"].(string); ok && v != "" {
		result.Tiles = aws.String(v)
	}
	if v, ok := tfMap["unregistered_sei_timecode"].(string); ok && v != "" {
		result.UnregisteredSeiTimecode = aws.String(v)
	}
	if v, ok := tfMap["write_mp4_packaging_type"].(string); ok && v != "" {
		result.WriteMp4PackagingType = aws.String(v)
	}
	return result
}

func expandMediaConvertMpeg2Settings(list []interface{}) *mediaconvert.Mpeg2Settings {
	result := &mediaconvert.Mpeg2Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["adaptive_quantization"].(string); ok && v != "" {
		result.AdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["codec_level"].(string); ok && v != "" {
		result.CodecLevel = aws.String(v)
	}
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["dynamic_sub_gop"].(string); ok && v != "" {
		result.DynamicSubGop = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_closed_cadence"].(int); ok {
		result.GopClosedCadence = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["gop_size_units"].(string); ok && v != "" {
		result.GopSizeUnits = aws.String(v)
	}
	if v, ok := tfMap["hrd_buffer_initial_fill_percentage"].(int); ok && v != 0 {
		result.HrdBufferInitialFillPercentage = aws.Int64(int64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["intra_dc_precision"].(string); ok && v != "" {
		result.IntraDcPrecision = aws.String(v)
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["min_i_interval"].(int); ok {
		result.MinIInterval = aws.Int64(int64(v))
	}
	if v, ok := tfMap["number_b_frames_between_reference_frames"].(int); ok {
		result.NumberBFramesBetweenReferenceFrames = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	if v, ok := tfMap["scene_change_detect"].(string); ok && v != "" {
		result.SceneChangeDetect = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["softness"].(int); ok {
		result.Softness = aws.Int64(int64(v))
	}
	if v, ok := tfMap["spatial_adaptive_quantization"].(string); ok && v != "" {
		result.SpatialAdaptiveQuantization = aws.String(v)
	}
	if v, ok := tfMap["syntax"].(string); ok && v != "" {
		result.Syntax = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["temporal_adaptive_quantization"].(string); ok && v != "" {
		result.TemporalAdaptiveQuantization = aws.String(v)
	}
	return result
}

func expandMediaConvertProresSettings(list []interface{}) *mediaconvert.ProresSettings {
	result := &mediaconvert.ProresSettings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["codec_profile"].(string); ok && v != "" {
		result.CodecProfile = aws.String(v)
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	return result
}

func expandMediaConvertVc3Settings(list []interface{}) *mediaconvert.Vc3Settings {
	result := &mediaconvert.Vc3Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["interlace_mode"].(string); ok && v != "" {
		result.InterlaceMode = aws.String(v)
	}
	if v, ok := tfMap["slow_pal"].(string); ok && v != "" {
		result.SlowPal = aws.String(v)
	}
	if v, ok := tfMap["telecine"].(string); ok && v != "" {
		result.Telecine = aws.String(v)
	}
	if v, ok := tfMap["vc3_class"].(string); ok && v != "" {
		result.Vc3Class = aws.String(v)
	}
	return result
}

func expandMediaConvertVp8Settings(list []interface{}) *mediaconvert.Vp8Settings {
	result := &mediaconvert.Vp8Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	return result
}

func expandMediaConvertVp9Settings(list []interface{}) *mediaconvert.Vp9Settings {
	result := &mediaconvert.Vp9Settings{}
	if list == nil || len(list) == 0 {
		return nil
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["bitrate"].(int); ok {
		result.Bitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_control"].(string); ok && v != "" {
		result.FramerateControl = aws.String(v)
	}
	if v, ok := tfMap["framerate_conversion_algorithm"].(string); ok && v != "" {
		result.FramerateConversionAlgorithm = aws.String(v)
	}
	if v, ok := tfMap["framerate_denominator"].(int); ok && v != 0 {
		result.FramerateDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["framerate_numerator"].(int); ok && v != 0 {
		result.FramerateNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["gop_size"].(float64); ok {
		result.GopSize = aws.Float64(float64(v))
	}
	if v, ok := tfMap["hrd_buffer_size"].(int); ok && v != 0 {
		result.HrdBufferSize = aws.Int64(int64(v))
	}
	if v, ok := tfMap["max_bitrate"].(int); ok {
		result.MaxBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_control"].(string); ok && v != "" {
		result.ParControl = aws.String(v)
	}
	if v, ok := tfMap["par_denominator"].(int); ok && v != 0 {
		result.ParDenominator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["par_numerator"].(int); ok && v != 0 {
		result.ParNumerator = aws.Int64(int64(v))
	}
	if v, ok := tfMap["quality_tuning_level"].(string); ok && v != "" {
		result.QualityTuningLevel = aws.String(v)
	}
	if v, ok := tfMap["rate_control_mode"].(string); ok && v != "" {
		result.RateControlMode = aws.String(v)
	}
	return result
}

func expandMediaConvertAv1QvbrSettings(list []interface{}) *mediaconvert.Av1QvbrSettings {
	result := &mediaconvert.Av1QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertH264QvbrSettings(list []interface{}) *mediaconvert.H264QvbrSettings {
	result := &mediaconvert.H264QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_average_bitrate"].(int); ok && v != 0 {
		result.MaxAverageBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}

func expandMediaConvertH265QvbrSettings(list []interface{}) *mediaconvert.H265QvbrSettings {
	result := &mediaconvert.H265QvbrSettings{}
	if list == nil || len(list) == 0 {
		return result
	}
	tfMap := list[0].(map[string]interface{})
	if v, ok := tfMap["max_average_bitrate"].(int); ok && v != 0 {
		result.MaxAverageBitrate = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level"].(int); ok {
		result.QvbrQualityLevel = aws.Int64(int64(v))
	}
	if v, ok := tfMap["qvbr_quality_level_fine_tune"].(float64); ok {
		result.QvbrQualityLevelFineTune = aws.Float64(float64(v))
	}
	return result
}
