# Bedrock

Amazon Bedrock is a fully managed service that makes high-performing foundation models (FMs) from leading AI companies and Amazon available for your use through a unified API.
Claude's context window accepts up to 200,000 tokens (roughly 150,000 words, or over 500 pages of material).

## Claude 3 Sonnet

Claude's context window accepts up to 200,000 tokens (roughly 150,000 words, or over 500 pages of material).

Model ID

```
anthropic.claude-3-sonnet-20240229-v1:0:28k
anthropic.claude-3-sonnet-20240229-v1:0:200k
anthropic.claude-3-sonnet-20240229-v1:0
```

## Amazon Bedrock

以下是 Amazon Bedrock 支援的操作：

- 批量刪除評估作業（BatchDeleteEvaluationJob）
- 創建評估作業（CreateEvaluationJob）
- 創建防護欄（CreateGuardrail）
- 創建防護欄版本（CreateGuardrailVersion）
- 創建推理配置檔（CreateInferenceProfile）
- 創建市場模型端點（CreateMarketplaceModelEndpoint）
- 創建模型複製作業（CreateModelCopyJob）
- 創建模型定制作業（CreateModelCustomizationJob）
- 創建模型導入作業（CreateModelImportJob）
- 創建模型調用作業（CreateModelInvocationJob）
- 創建預配置模型吞吐量（CreateProvisionedModelThroughput）
- 刪除自定義模型（DeleteCustomModel）
- 刪除防護欄（DeleteGuardrail）
- 刪除導入的模型（DeleteImportedModel）
- 刪除推理配置檔（DeleteInferenceProfile）
- 刪除市場模型端點（DeleteMarketplaceModelEndpoint）
- 刪除模型調用日誌配置（DeleteModelInvocationLoggingConfiguration）
- 刪除預配置模型吞吐量（DeleteProvisionedModelThroughput）
- 注銷市場模型端點（DeregisterMarketplaceModelEndpoint）
- 獲取自定義模型（GetCustomModel）
- 獲取評估作業（GetEvaluationJob）
- 獲取基礎模型（GetFoundationModel）
- 獲取防護欄（GetGuardrail）
- 獲取導入的模型（GetImportedModel）
- 獲取推理配置檔（GetInferenceProfile）
- 獲取市場模型端點（GetMarketplaceModelEndpoint）
- 獲取模型複製作業（GetModelCopyJob）
- 獲取模型定制作業（GetModelCustomizationJob）
- 獲取模型導入作業（GetModelImportJob）
- 獲取模型調用作業（GetModelInvocationJob）
- 獲取模型調用日誌配置（GetModelInvocationLoggingConfiguration）
- 獲取提示路由器（GetPromptRouter）
- 獲取預配置模型吞吐量（GetProvisionedModelThroughput）
- 列出自定義模型（ListCustomModels）
- 列出評估作業（ListEvaluationJobs）
- 列出基礎模型（ListFoundationModels）
- 列出防護欄（ListGuardrails）
- 列出導入的模型（ListImportedModels）
- 列出推理配置檔（ListInferenceProfiles）
- 列出市場模型端點（ListMarketplaceModelEndpoints）
- 列出模型複製作業（ListModelCopyJobs）
- 列出模型定制作業（ListModelCustomizationJobs）
- 列出模型導入作業（ListModelImportJobs）
- 列出模型調用作業（ListModelInvocationJobs）
- 列出提示路由器（ListPromptRouters）
- 列出預配置模型吞吐量（ListProvisionedModelThroughputs）
- 列出資源標籤（ListTagsForResource）
- 設置模型調用日誌配置（PutModelInvocationLoggingConfiguration）
- 註冊市場模型端點（RegisterMarketplaceModelEndpoint）
- 停止評估作業（StopEvaluationJob）
- 停止模型定制作業（StopModelCustomizationJob）
- 停止模型調用作業（StopModelInvocationJob）
- 標籤資源（TagResource）
- 取消標籤資源（UntagResource）
- 更新防護欄（UpdateGuardrail）
- 更新市場模型端點（UpdateMarketplaceModelEndpoint）
- 更新預配置模型吞吐量（UpdateProvisionedModelThroughput）

## Amazon Bedrock Runtime

以下是 Amazon Bedrock Runtime 支援的操作：

- 應用防護欄（ApplyGuardrail）
- 對話（Converse）
- 對話流（ConverseStream）
- 獲取異步調用（GetAsyncInvoke）
- 調用模型（InvokeModel）
- 調用帶有響應流的模型（InvokeModelWithResponseStream）
- 列出異步調用（ListAsyncInvokes）
- 開始異步調用（StartAsyncInvoke）
