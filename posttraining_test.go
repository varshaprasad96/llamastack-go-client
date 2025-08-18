// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackgoclient_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/llamastack-go-client-go"
	"github.com/stainless-sdks/llamastack-go-client-go/internal/testutil"
	"github.com/stainless-sdks/llamastack-go-client-go/option"
)

func TestPostTrainingFineTuneSupervisedWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.FineTuneSupervised(context.TODO(), llamastackgoclient.PostTrainingFineTuneSupervisedParams{
		HyperparamSearchConfig: map[string]llamastackgoclient.PostTrainingFineTuneSupervisedParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackgoclient.PostTrainingFineTuneSupervisedParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		TrainingConfig: llamastackgoclient.TrainingConfigParam{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackgoclient.TrainingConfigDataConfigParam{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackgoclient.Bool(true),
				TrainOnInput:        llamastackgoclient.Bool(true),
				ValidationDatasetID: llamastackgoclient.String("validation_dataset_id"),
			},
			Dtype: llamastackgoclient.String("dtype"),
			EfficiencyConfig: llamastackgoclient.TrainingConfigEfficiencyConfigParam{
				EnableActivationCheckpointing: llamastackgoclient.Bool(true),
				EnableActivationOffloading:    llamastackgoclient.Bool(true),
				FsdpCPUOffload:                llamastackgoclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackgoclient.Bool(true),
			},
			MaxValidationSteps: llamastackgoclient.Int(0),
			OptimizerConfig: llamastackgoclient.TrainingConfigOptimizerConfigParam{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
		AlgorithmConfig: llamastackgoclient.PostTrainingFineTuneSupervisedParamsAlgorithmConfigUnion{
			OfLoRa: &llamastackgoclient.PostTrainingFineTuneSupervisedParamsAlgorithmConfigLoRa{
				Alpha:             0,
				ApplyLoraToMlp:    true,
				ApplyLoraToOutput: true,
				LoraAttnModules:   []string{"string"},
				Rank:              0,
				QuantizeBase:      llamastackgoclient.Bool(true),
				UseDora:           llamastackgoclient.Bool(true),
			},
		},
		CheckpointDir: llamastackgoclient.String("checkpoint_dir"),
		Model:         llamastackgoclient.String("model"),
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.ListJobs(context.TODO())
	if err != nil {
		var apierr *llamastackgoclient.Error
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
	client := llamastackgoclient.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.PostTraining.OptimizePreferences(context.TODO(), llamastackgoclient.PostTrainingOptimizePreferencesParams{
		AlgorithmConfig: llamastackgoclient.PostTrainingOptimizePreferencesParamsAlgorithmConfig{
			Beta:     0,
			LossType: "sigmoid",
		},
		FinetunedModel: "finetuned_model",
		HyperparamSearchConfig: map[string]llamastackgoclient.PostTrainingOptimizePreferencesParamsHyperparamSearchConfigUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		JobUuid: "job_uuid",
		LoggerConfig: map[string]llamastackgoclient.PostTrainingOptimizePreferencesParamsLoggerConfigUnion{
			"foo": {
				OfBool: llamastackgoclient.Bool(true),
			},
		},
		TrainingConfig: llamastackgoclient.TrainingConfigParam{
			GradientAccumulationSteps: 0,
			MaxStepsPerEpoch:          0,
			NEpochs:                   0,
			DataConfig: llamastackgoclient.TrainingConfigDataConfigParam{
				BatchSize:           0,
				DataFormat:          "instruct",
				DatasetID:           "dataset_id",
				Shuffle:             true,
				Packed:              llamastackgoclient.Bool(true),
				TrainOnInput:        llamastackgoclient.Bool(true),
				ValidationDatasetID: llamastackgoclient.String("validation_dataset_id"),
			},
			Dtype: llamastackgoclient.String("dtype"),
			EfficiencyConfig: llamastackgoclient.TrainingConfigEfficiencyConfigParam{
				EnableActivationCheckpointing: llamastackgoclient.Bool(true),
				EnableActivationOffloading:    llamastackgoclient.Bool(true),
				FsdpCPUOffload:                llamastackgoclient.Bool(true),
				MemoryEfficientFsdpWrap:       llamastackgoclient.Bool(true),
			},
			MaxValidationSteps: llamastackgoclient.Int(0),
			OptimizerConfig: llamastackgoclient.TrainingConfigOptimizerConfigParam{
				Lr:             0,
				NumWarmupSteps: 0,
				OptimizerType:  "adam",
				WeightDecay:    0,
			},
		},
	})
	if err != nil {
		var apierr *llamastackgoclient.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
