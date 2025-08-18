// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/varshaprasad96/llamastack-go-client"
	"github.com/varshaprasad96/llamastack-go-client/internal/testutil"
	"github.com/varshaprasad96/llamastack-go-client/option"
)

func TestPostTrainingCancel(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	err := client.PostTraining.Cancel(context.TODO(), llamastackclient.PostTrainingCancelParams{
		JobUuid: "job_uuid",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostTrainingFineTuneSupervisedWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.FineTuneSupervised(context.TODO(), llamastackclient.PostTrainingFineTuneSupervisedParams{
		HyperparamSearchConfig: map[string]llamastackclient.PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.PostTrainingFineTuneSupervisedParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.TrainingConfigParam{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.TrainingConfigDataConfigParam{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.TrainingConfigEfficiencyConfigParam{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.TrainingConfigOptimizerConfigParam{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
		AlgorithmConfig: llamastackclient.PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion{
			OfLoRa: &llamastackclient.PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa{
				Alpha:             0,
				ApplyLoraToMlp:    true,
				ApplyLoraToOutput: true,
				LoraAttnModules:   []string{"string"},
				Rank:              0,
				QuantizeBase:      llamastackclient.Bool(true),
				UseDora:           llamastackclient.Bool(true),
			},
		},
		CheckpointDir: llamastackclient.String("checkpoint_dir"),
		Model:         llamastackclient.String("model"),
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostTrainingGetArtifacts(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.GetArtifacts(context.TODO(), llamastackclient.PostTrainingGetArtifactsParams{
		JobUuid: "job_uuid",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostTrainingGetStatus(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.GetStatus(context.TODO(), llamastackclient.PostTrainingGetStatusParams{
		JobUuid: "job_uuid",
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostTrainingListJobs(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.ListJobs(context.TODO())
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPostTrainingOptimizePreferencesWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.OptimizePreferences(context.TODO(), llamastackclient.PostTrainingOptimizePreferencesParams{
		AlgorithmConfig: llamastackclient.PostTrainingOptimizePreferencesParamsAlgorithmConfig{
			Beta:     0,
			LossType: "sigmoid",
		},
		FinetunedModel: "finetuned_model",
		HyperparamSearchConfig: map[string]llamastackclient.PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackclient.PostTrainingOptimizePreferencesParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackclient.Bool(true),
			},
		},
		TrainingConfig: llamastackclient.TrainingConfigParam{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackclient.TrainingConfigDataConfigParam{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackclient.Bool(true),
				TrainOnInput:        llamastackclient.Bool(true),
				ValidationDatasetID: llamastackclient.String("validation_dataset_id"),
			},
			Dtype: llamastackclient.String("dtype"),
			EfficiencyConfig: llamastackclient.TrainingConfigEfficiencyConfigParam{
				EnableActivationCheckpointing: llamastackclient.Bool(true),
				EnableActivationOffloading:    llamastackclient.Bool(true),
				FsdpCPUOffload:                llamastackclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackclient.Bool(true),
			},
			MaxValidationSteps: llamastackclient.Int(0),
			OptimizerConfig: llamastackclient.TrainingConfigOptimizerConfigParam{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
	})
	if err != nil {
		var apierr *llamastackclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
