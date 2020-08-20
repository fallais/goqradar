package goqradar

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

//------------------------------------------------------------------------------
// Structures
//------------------------------------------------------------------------------

// Schemas is a QRadar schemas
type Schemas struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	SampleQuery struct {
		Fields []struct {
			ArgumentFields []struct {
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				LocalizedName  string `json:"localized_name"`
				SemanticType   string `json:"semantic_type"`
			} `json:"argument_fields"`
			ContextualType string `json:"contextual_type"`
			DataType       string `json:"data_type"`
			Filter         struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				LeftFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					LeftFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters []string `json:"parameters"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters  []string `json:"parameters"`
					RightFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters []string `json:"parameters"`
					} `json:"right_filter"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters  []string `json:"parameters"`
				RightFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					LeftFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters []string `json:"parameters"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters  []string `json:"parameters"`
					RightFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters []string `json:"parameters"`
					} `json:"right_filter"`
				} `json:"right_filter"`
			} `json:"filter"`
			Function struct {
				Aggregate     bool `json:"aggregate"`
				ArgumentTypes []struct {
					Position int    `json:"position"`
					Type     string `json:"type"`
				} `json:"argument_types"`
				LocalizedName  string `json:"localized_name"`
				Name           string `json:"name"`
				ReturnDataType string `json:"return_data_type"`
			} `json:"function"`
			LocalizedName string `json:"localized_name"`
			SemanticType  string `json:"semantic_type"`
		} `json:"fields"`
		Filters []struct {
			Argument struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         string `json:"filter"`
				Function       struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"argument"`
			LeftFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				LeftFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters []string `json:"parameters"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters  []string `json:"parameters"`
				RightFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters []string `json:"parameters"`
				} `json:"right_filter"`
			} `json:"left_filter"`
			Operator struct {
				Connective    string `json:"connective"`
				DataType      string `json:"data_type"`
				LocalizedName string `json:"localized_name"`
				Name          string `json:"name"`
			} `json:"operator"`
			Parameters  []string `json:"parameters"`
			RightFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				LeftFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters []string `json:"parameters"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters  []string `json:"parameters"`
				RightFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters []string `json:"parameters"`
				} `json:"right_filter"`
			} `json:"right_filter"`
		} `json:"filters"`
		Range struct {
			Limit  int `json:"limit"`
			Offset int `json:"offset"`
		} `json:"range"`
		Sorts []struct {
			Direction string `json:"direction"`
			Field     struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					LeftFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						LeftFilter struct {
							Argument struct {
								ArgumentFields []struct {
									ContextualType string `json:"contextual_type"`
									DataType       string `json:"data_type"`
									LocalizedName  string `json:"localized_name"`
									SemanticType   string `json:"semantic_type"`
								} `json:"argument_fields"`
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								Filter         string `json:"filter"`
								Function       struct {
									Aggregate     bool `json:"aggregate"`
									ArgumentTypes []struct {
										Position int    `json:"position"`
										Type     string `json:"type"`
									} `json:"argument_types"`
									LocalizedName  string `json:"localized_name"`
									Name           string `json:"name"`
									ReturnDataType string `json:"return_data_type"`
								} `json:"function"`
								LocalizedName string `json:"localized_name"`
								SemanticType  string `json:"semantic_type"`
							} `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameters []string `json:"parameters"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters  []string `json:"parameters"`
						RightFilter struct {
							Argument struct {
								ArgumentFields []struct {
									ContextualType string `json:"contextual_type"`
									DataType       string `json:"data_type"`
									LocalizedName  string `json:"localized_name"`
									SemanticType   string `json:"semantic_type"`
								} `json:"argument_fields"`
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								Filter         string `json:"filter"`
								Function       struct {
									Aggregate     bool `json:"aggregate"`
									ArgumentTypes []struct {
										Position int    `json:"position"`
										Type     string `json:"type"`
									} `json:"argument_types"`
									LocalizedName  string `json:"localized_name"`
									Name           string `json:"name"`
									ReturnDataType string `json:"return_data_type"`
								} `json:"function"`
								LocalizedName string `json:"localized_name"`
								SemanticType  string `json:"semantic_type"`
							} `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameters []string `json:"parameters"`
						} `json:"right_filter"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameters  []string `json:"parameters"`
					RightFilter struct {
						Argument struct {
							ArgumentFields []struct {
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								LocalizedName  string `json:"localized_name"`
								SemanticType   string `json:"semantic_type"`
							} `json:"argument_fields"`
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							Filter         string `json:"filter"`
							Function       struct {
								Aggregate     bool `json:"aggregate"`
								ArgumentTypes []struct {
									Position int    `json:"position"`
									Type     string `json:"type"`
								} `json:"argument_types"`
								LocalizedName  string `json:"localized_name"`
								Name           string `json:"name"`
								ReturnDataType string `json:"return_data_type"`
							} `json:"function"`
							LocalizedName string `json:"localized_name"`
							SemanticType  string `json:"semantic_type"`
						} `json:"argument"`
						LeftFilter struct {
							Argument struct {
								ArgumentFields []struct {
									ContextualType string `json:"contextual_type"`
									DataType       string `json:"data_type"`
									LocalizedName  string `json:"localized_name"`
									SemanticType   string `json:"semantic_type"`
								} `json:"argument_fields"`
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								Filter         string `json:"filter"`
								Function       struct {
									Aggregate     bool `json:"aggregate"`
									ArgumentTypes []struct {
										Position int    `json:"position"`
										Type     string `json:"type"`
									} `json:"argument_types"`
									LocalizedName  string `json:"localized_name"`
									Name           string `json:"name"`
									ReturnDataType string `json:"return_data_type"`
								} `json:"function"`
								LocalizedName string `json:"localized_name"`
								SemanticType  string `json:"semantic_type"`
							} `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameters []string `json:"parameters"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameters  []string `json:"parameters"`
						RightFilter struct {
							Argument struct {
								ArgumentFields []struct {
									ContextualType string `json:"contextual_type"`
									DataType       string `json:"data_type"`
									LocalizedName  string `json:"localized_name"`
									SemanticType   string `json:"semantic_type"`
								} `json:"argument_fields"`
								ContextualType string `json:"contextual_type"`
								DataType       string `json:"data_type"`
								Filter         string `json:"filter"`
								Function       struct {
									Aggregate     bool `json:"aggregate"`
									ArgumentTypes []struct {
										Position int    `json:"position"`
										Type     string `json:"type"`
									} `json:"argument_types"`
									LocalizedName  string `json:"localized_name"`
									Name           string `json:"name"`
									ReturnDataType string `json:"return_data_type"`
								} `json:"function"`
								LocalizedName string `json:"localized_name"`
								SemanticType  string `json:"semantic_type"`
							} `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameters []string `json:"parameters"`
						} `json:"right_filter"`
					} `json:"right_filter"`
				} `json:"filter"`
				Function struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"field"`
		} `json:"sorts"`
	} `json:"sample_query"`
}

// SchemasPaginatedResponse is the paginated response.
type SchemasPaginatedResponse struct {
	Total       int        `json:"total"`
	Min         int        `json:"min"`
	Max         int        `json:"max"`
	ListSchemas []*Schemas `json:"offense_types"`
}

// Field is a QRadar field
type Field struct {
	ArgumentFields struct {
		ContextualType string `json:"contextual_type"`
		DataType       string `json:"data_type"`
		LocalizedName  string `json:"localized_name"`
		SemanticType   string `json:"semantic_type"`
	} `json:"argument_fields"`
	ContextualType string `json:"contextual_type"`
	DataType       string `json:"data_type"`
	Filter         struct {
		Argument struct {
			ArgumentFields []struct {
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				LocalizedName  string `json:"localized_name"`
				SemanticType   string `json:"semantic_type"`
			} `json:"argument_fields"`
			ContextualType string `json:"contextual_type"`
			DataType       string `json:"data_type"`
			Filter         string `json:"filter"`
			Function       struct {
				Aggregate     bool `json:"aggregate"`
				ArgumentTypes []struct {
					Position int    `json:"position"`
					Type     string `json:"type"`
				} `json:"argument_types"`
				LocalizedName  string `json:"localized_name"`
				Name           string `json:"name"`
				ReturnDataType string `json:"return_data_type"`
			} `json:"function"`
			LocalizedName string `json:"localized_name"`
			SemanticType  string `json:"semantic_type"`
		} `json:"argument"`
		LeftFilter struct {
			Argument struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         string `json:"filter"`
				Function       struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"argument"`
			LeftFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters []string `json:"parameters"`
			} `json:"left_filter"`
			Operator struct {
				Connective    string `json:"connective"`
				DataType      string `json:"data_type"`
				LocalizedName string `json:"localized_name"`
				Name          string `json:"name"`
			} `json:"operator"`
			Parameters  []string `json:"parameters"`
			RightFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters []string `json:"parameters"`
			} `json:"right_filter"`
		} `json:"left_filter"`
		Operator struct {
			Connective    string `json:"connective"`
			DataType      string `json:"data_type"`
			LocalizedName string `json:"localized_name"`
			Name          string `json:"name"`
		} `json:"operator"`
		Parameters  []string `json:"parameters"`
		RightFilter struct {
			Argument struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         string `json:"filter"`
				Function       struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"argument"`
			LeftFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters []string `json:"parameters"`
			} `json:"left_filter"`
			Operator struct {
				Connective    string `json:"connective"`
				DataType      string `json:"data_type"`
				LocalizedName string `json:"localized_name"`
				Name          string `json:"name"`
			} `json:"operator"`
			Parameters  []string `json:"parameters"`
			RightFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameters []string `json:"parameters"`
			} `json:"right_filter"`
		} `json:"right_filter"`
	} `json:"filter"`
	Function struct {
		Aggregate     bool `json:"aggregate"`
		ArgumentTypes []struct {
			Position int    `json:"position"`
			Type     string `json:"type"`
		} `json:"argument_types"`
		LocalizedName  string `json:"localized_name"`
		Name           string `json:"name"`
		ReturnDataType string `json:"return_data_type"`
	} `json:"function"`
	LocalizedName string `json:"localized_name"`
	SemanticType  string `json:"semantic_type"`
}

// FieldsPaginatedResponse is the paginated response.
type FieldsPaginatedResponse struct {
	Total  int      `json:"total"`
	Min    int      `json:"min"`
	Max    int      `json:"max"`
	Fields []*Field `json:"offense_types"`
}

// Function is a QRadar function.
type Function struct {
	Aggregate     bool `json:"aggregate"`
	ArgumentTypes []struct {
		Position int    `json:"position"`
		Type     string `json:"type"`
	} `json:"argument_types"`
	LocalizedName  string `json:"localized_name"`
	Name           string `json:"name"`
	ReturnDataType string `json:"return_data_type"`
}

// FunctionsPaginatedResponse is the paginated response.
type FunctionsPaginatedResponse struct {
	Total     int         `json:"total"`
	Min       int         `json:"min"`
	Max       int         `json:"max"`
	Functions []*Function `json:"offense_types"`
}

// Operator is a QRadar operator.
type Operator struct {
	Connective    string `json:"connective"`
	DataType      string `json:"data_type"`
	LocalizedName string `json:"localized_name"`
	Name          string `json:"name"`
}

// OperatorsPaginatedResponse is the paginated response.
type OperatorsPaginatedResponse struct {
	Total     int         `json:"total"`
	Min       int         `json:"min"`
	Max       int         `json:"max"`
	Operators []*Operator `json:"offense_types"`
}

// Search is a QRadar search
type Search struct {
	Description string `json:"description"`
	Handle      string `json:"handle"`
	Header      struct {
		Columns []struct {
			ColumnName string `json:"column_name"`
			Field      struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         struct {
					Argument   string `json:"argument"`
					LeftFilter struct {
						Argument   string `json:"argument"`
						LeftFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter   string `json:"parameter"`
						RightFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"right_filter"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter   string `json:"parameter"`
					RightFilter struct {
						Argument   string `json:"argument"`
						LeftFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter   string `json:"parameter"`
						RightFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"right_filter"`
					} `json:"right_filter"`
				} `json:"filter"`
				Function struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"field"`
		} `json:"columns"`
	} `json:"header"`
	Query struct {
		Fields []struct {
			ArgumentFields []struct {
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				LocalizedName  string `json:"localized_name"`
				SemanticType   string `json:"semantic_type"`
			} `json:"argument_fields"`
			ContextualType string `json:"contextual_type"`
			DataType       string `json:"data_type"`
			Filter         struct {
				Argument   string `json:"argument"`
				LeftFilter struct {
					Argument   string `json:"argument"`
					LeftFilter struct {
						Argument string `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter string `json:"parameter"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter   string `json:"parameter"`
					RightFilter struct {
						Argument string `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter string `json:"parameter"`
					} `json:"right_filter"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameter   string `json:"parameter"`
				RightFilter struct {
					Argument   string `json:"argument"`
					LeftFilter struct {
						Argument string `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter string `json:"parameter"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter   string `json:"parameter"`
					RightFilter struct {
						Argument string `json:"argument"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter string `json:"parameter"`
					} `json:"right_filter"`
				} `json:"right_filter"`
			} `json:"filter"`
			Function struct {
				Aggregate     bool `json:"aggregate"`
				ArgumentTypes []struct {
					Position int    `json:"position"`
					Type     string `json:"type"`
				} `json:"argument_types"`
				LocalizedName  string `json:"localized_name"`
				Name           string `json:"name"`
				ReturnDataType string `json:"return_data_type"`
			} `json:"function"`
			LocalizedName string `json:"localized_name"`
			SemanticType  string `json:"semantic_type"`
		} `json:"fields"`
		Filters []struct {
			Argument struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         string `json:"filter"`
				Function       struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"argument"`
			LeftFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				LeftFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter string `json:"parameter"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameter   string `json:"parameter"`
				RightFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter string `json:"parameter"`
				} `json:"right_filter"`
			} `json:"left_filter"`
			Operator struct {
				Connective    string `json:"connective"`
				DataType      string `json:"data_type"`
				LocalizedName string `json:"localized_name"`
				Name          string `json:"name"`
			} `json:"operator"`
			Parameter   string `json:"parameter"`
			RightFilter struct {
				Argument struct {
					ArgumentFields []struct {
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						LocalizedName  string `json:"localized_name"`
						SemanticType   string `json:"semantic_type"`
					} `json:"argument_fields"`
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					Filter         string `json:"filter"`
					Function       struct {
						Aggregate     bool `json:"aggregate"`
						ArgumentTypes []struct {
							Position int    `json:"position"`
							Type     string `json:"type"`
						} `json:"argument_types"`
						LocalizedName  string `json:"localized_name"`
						Name           string `json:"name"`
						ReturnDataType string `json:"return_data_type"`
					} `json:"function"`
					LocalizedName string `json:"localized_name"`
					SemanticType  string `json:"semantic_type"`
				} `json:"argument"`
				LeftFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter string `json:"parameter"`
				} `json:"left_filter"`
				Operator struct {
					Connective    string `json:"connective"`
					DataType      string `json:"data_type"`
					LocalizedName string `json:"localized_name"`
					Name          string `json:"name"`
				} `json:"operator"`
				Parameter   string `json:"parameter"`
				RightFilter struct {
					Argument struct {
						ArgumentFields []struct {
							ContextualType string `json:"contextual_type"`
							DataType       string `json:"data_type"`
							LocalizedName  string `json:"localized_name"`
							SemanticType   string `json:"semantic_type"`
						} `json:"argument_fields"`
						ContextualType string `json:"contextual_type"`
						DataType       string `json:"data_type"`
						Filter         string `json:"filter"`
						Function       struct {
							Aggregate     bool `json:"aggregate"`
							ArgumentTypes []struct {
								Position int    `json:"position"`
								Type     string `json:"type"`
							} `json:"argument_types"`
							LocalizedName  string `json:"localized_name"`
							Name           string `json:"name"`
							ReturnDataType string `json:"return_data_type"`
						} `json:"function"`
						LocalizedName string `json:"localized_name"`
						SemanticType  string `json:"semantic_type"`
					} `json:"argument"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter string `json:"parameter"`
				} `json:"right_filter"`
			} `json:"right_filter"`
		} `json:"filters"`
		Range struct {
			Limit  int `json:"limit"`
			Offset int `json:"offset"`
		} `json:"range"`
		Sorts []struct {
			Direction string `json:"direction"`
			Field     struct {
				ArgumentFields []struct {
					ContextualType string `json:"contextual_type"`
					DataType       string `json:"data_type"`
					LocalizedName  string `json:"localized_name"`
					SemanticType   string `json:"semantic_type"`
				} `json:"argument_fields"`
				ContextualType string `json:"contextual_type"`
				DataType       string `json:"data_type"`
				Filter         struct {
					Argument   string `json:"argument"`
					LeftFilter struct {
						Argument   string `json:"argument"`
						LeftFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter   string `json:"parameter"`
						RightFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"right_filter"`
					} `json:"left_filter"`
					Operator struct {
						Connective    string `json:"connective"`
						DataType      string `json:"data_type"`
						LocalizedName string `json:"localized_name"`
						Name          string `json:"name"`
					} `json:"operator"`
					Parameter   string `json:"parameter"`
					RightFilter struct {
						Argument   string `json:"argument"`
						LeftFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"left_filter"`
						Operator struct {
							Connective    string `json:"connective"`
							DataType      string `json:"data_type"`
							LocalizedName string `json:"localized_name"`
							Name          string `json:"name"`
						} `json:"operator"`
						Parameter   string `json:"parameter"`
						RightFilter struct {
							Argument string `json:"argument"`
							Operator struct {
								Connective    string `json:"connective"`
								DataType      string `json:"data_type"`
								LocalizedName string `json:"localized_name"`
								Name          string `json:"name"`
							} `json:"operator"`
							Parameter string `json:"parameter"`
						} `json:"right_filter"`
					} `json:"right_filter"`
				} `json:"filter"`
				Function struct {
					Aggregate     bool `json:"aggregate"`
					ArgumentTypes []struct {
						Position int    `json:"position"`
						Type     string `json:"type"`
					} `json:"argument_types"`
					LocalizedName  string `json:"localized_name"`
					Name           string `json:"name"`
					ReturnDataType string `json:"return_data_type"`
				} `json:"function"`
				LocalizedName string `json:"localized_name"`
				SemanticType  string `json:"semantic_type"`
			} `json:"field"`
		} `json:"sorts"`
	} `json:"query"`
	Retention struct {
		CreationDate     int `json:"creation_date"`
		ExpiresAt        int `json:"expires_at"`
		LastAccessedDate int `json:"last_accessed_date"`
		RetainDuration   int `json:"retain_duration"`
	} `json:"retention"`
	SearchType string `json:"search_type"`
	Status     string `json:"status"`
}

// DynamicSearchesPaginatedResponse is the paginated response.
type DynamicSearchesPaginatedResponse struct {
	Total           int       `json:"total"`
	Min             int       `json:"min"`
	Max             int       `json:"max"`
	DynamicSearches []*Search `json:"offense_types"`
}

// PostedSearch is QRadar search
type PostedSearch struct {
	Query string `json:"query"`
}

// SearchResult is a QRadar search result
type SearchResult struct {
	Columns struct {
		String string `json:"String"`
	} `json:"columns"`
}

//------------------------------------------------------------------------------
// Functions
//------------------------------------------------------------------------------

// ListSchemas returns the schemas with given fields, filters.
func (endpoint *Endpoint) ListSchemas(ctx context.Context, fields, filter string, min, max int) (*SchemasPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/schemas", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &SchemasPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.ListSchemas)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// GetSchemas retrieves a Schemas .
func (endpoint *Endpoint) GetSchemas(ctx context.Context, name string, fields string) (*Schemas, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/schemas/"+name, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *Schemas

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// ListFields returns the fields with given fields, filters.
func (endpoint *Endpoint) ListFields(ctx context.Context, name, fields, filter string, min, max int) (*FieldsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/schemas/"+name+"/fields", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &FieldsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Fields)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListFunctions returns the functions with given fields, filters.
func (endpoint *Endpoint) ListFunctions(ctx context.Context, name, fields, filter string, min, max int) (*FunctionsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/schemas/"+name+"/functions", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &FunctionsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Functions)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListOperators returns the operators with given fields, filters.
func (endpoint *Endpoint) ListOperators(ctx context.Context, name, fields, filter string, min, max int) (*OperatorsPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/schemas/"+name+"/operators", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &OperatorsPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.Operators)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// ListDynamicSearches returns the searches with given fields, filters.
func (endpoint *Endpoint) ListDynamicSearches(ctx context.Context, fields, filter string, min, max int) (*DynamicSearchesPaginatedResponse, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}
	if filter != "" {
		options = append(options, WithParam("filter", filter))
	}
	options = append(options, WithHeader("Range", fmt.Sprintf("items=%d-%d", min, max)))

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/searches", options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Process the Content-Range
	min, max, total, err := parseContentRange(resp.Header.Get("Content-Range"))
	if err != nil {
		return nil, fmt.Errorf("error while parsing the content-range [%s]: %s", resp.Header.Get("Content-Range"), err)
	}

	// Prepare the response
	response := &DynamicSearchesPaginatedResponse{
		Total: total,
		Min:   min,
		Max:   max,
	}

	// Decode the response
	err = json.NewDecoder(resp.Body).Decode(&response.DynamicSearches)
	if err != nil {
		return nil, fmt.Errorf("error while decoding the response: %s", err)
	}

	return response, nil
}

// CreateDynamicSearch posts a search to be performed by the service.
func (endpoint *Endpoint) CreateDynamicSearch(ctx context.Context, data map[string]interface{}) (*PostedSearch, error) {

	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return nil, fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/dynamic_search/searches"

	// Create the data
	d, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("Error while marshalling the values : %s", err)
	}

	// Create the request
	req, err := http.NewRequest("POST", reqURL.String(), bytes.NewBuffer(d))
	if err != nil {
		return nil, fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error while doing the request : %s", err)
	}
	defer resp.Body.Close()

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *PostedSearch

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// GetDynamicSearch retrieves a search .
func (endpoint *Endpoint) GetDynamicSearch(ctx context.Context, handle, fields string) (*Search, error) {
	// Options
	options := []Option{}
	if fields != "" {
		options = append(options, WithParam("fields", fields))
	}

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/searches/"+handle, options...)
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *Search

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}

// DeleteDynamicSearch by ID
func (endpoint *Endpoint) DeleteDynamicSearch(ctx context.Context, handle string) error {
	// Prepare the URL
	var reqURL *url.URL
	reqURL, err := url.Parse(endpoint.client.BaseURL)
	if err != nil {
		return fmt.Errorf("Error while parsing the URL : %s", err)
	}
	reqURL.Path += "/dynamic_search/searches/" + handle

	// Create the request
	req, err := http.NewRequest("DELETE", reqURL.String(), nil)
	if err != nil {
		return fmt.Errorf("Error while creating the request : %s", err)
	}

	// Set HTTP headers
	req.Header.Set("SEC", endpoint.client.Token)
	req.Header.Set("Version", endpoint.client.Version)
	req.Header.Set("Content-Type", "application/json")

	// Do the request
	resp, err := endpoint.client.client.Do(req)
	if err != nil {
		return fmt.Errorf("Error while doing the request : %s", err)
	}

	// Check the status code
	if resp.StatusCode != 204 {
		return fmt.Errorf("Status code is %d : Error while reading the body", resp.StatusCode)
	}

	return nil

}

// GetDynamicSearchResult retrieves a search result.
func (endpoint *Endpoint) GetDynamicSearchResult(ctx context.Context, handle string) (*SearchResult, error) {

	// Do the request
	resp, err := endpoint.client.do(ctx, http.MethodGet, "/dynamic_search/searches/"+handle+"/results")
	if err != nil {
		return nil, fmt.Errorf("error while calling the endpoint: %s", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error with the status code: %d", resp.StatusCode)
	}

	// Read the respsonse
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error while reading the request : %s", err)
	}

	// Prepare the response
	var response *SearchResult

	// Unmarshal the response
	err = json.Unmarshal([]byte(body), &response)
	if err != nil {
		return nil, fmt.Errorf("Error while unmarshalling the response : %s. HTTP response is : %s", err, string(body))
	}

	return response, nil
}
