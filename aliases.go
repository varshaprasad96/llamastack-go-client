// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package llamastackclient

import (
	"github.com/varshaprasad96/llamastack-go-client/internal/apierror"
	"github.com/varshaprasad96/llamastack-go-client/packages/param"
	"github.com/varshaprasad96/llamastack-go-client/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// A image content item
//
// This is an alias to an internal type.
type ImageContentItem = shared.ImageContentItem

// Image as a base64 encoded string or an URL
//
// This is an alias to an internal type.
type ImageContentItemImage = shared.ImageContentItemImage

// Discriminator type of the content item. Always "image"
//
// This is an alias to an internal type.
type ImageContentItemType = shared.ImageContentItemType

// Equals "text"
const ImageContentItemTypeText = shared.ImageContentItemTypeText

// Equals "image"
const ImageContentItemTypeImage = shared.ImageContentItemTypeImage

// Equals "tool_call"
const ImageContentItemTypeToolCall = shared.ImageContentItemTypeToolCall

// A image content item
//
// This is an alias to an internal type.
type ImageContentItemParam = shared.ImageContentItemParam

// Image as a base64 encoded string or an URL
//
// This is an alias to an internal type.
type ImageContentItemImageParam = shared.ImageContentItemImageParam

// A scoring result for a single row.
//
// This is an alias to an internal type.
type ScoringResult = shared.ScoringResult

// This is an alias to an internal type.
type ScoringResultAggregatedResultUnion = shared.ScoringResultAggregatedResultUnion

// This is an alias to an internal type.
type ScoringResultScoreRowUnion = shared.ScoringResultScoreRowUnion

// A text content item
//
// This is an alias to an internal type.
type TextContentItem = shared.TextContentItem

// Discriminator type of the content item. Always "text"
//
// This is an alias to an internal type.
type TextContentItemType = shared.TextContentItemType

// Equals "text"
const TextContentItemTypeText = shared.TextContentItemTypeText

// Equals "image"
const TextContentItemTypeImage = shared.TextContentItemTypeImage

// Equals "tool_call"
const TextContentItemTypeToolCall = shared.TextContentItemTypeToolCall

// A text content item
//
// This is an alias to an internal type.
type TextContentItemParam = shared.TextContentItemParam
