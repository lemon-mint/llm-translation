<script lang="ts">
  import { TranslationService } from "$lib/proto/api/v1beta1/translate_connect";
  import {
    createPromiseClient,
    type PromiseClient,
    type Transport,
  } from "@connectrpc/connect";
  import { createConnectTransport } from "@connectrpc/connect-web";
  import { Language, LLMOptions, LLMOptions_Provider, TranslateRequest } from "$lib/proto/api/v1beta1/translate_pb";
  import { onMount } from "svelte";
  import "./styles.css";

  let transport: Transport;
  let client: PromiseClient<typeof TranslationService>;

  let inputText = "";
  let outputText = "";
  let outputLanguage = "Spanish";

  let apiKey = "";
  let apiType = "SERVER";
  let model = "";

  function resetAdvancedSettings() {
    apiType = "SERVER";
    apiKey = "";
    model = "";
    storeCredentials();
  }

  function storeCredentials() {
    localStorage.setItem("apiKey", apiKey);
    localStorage.setItem("apiType", apiType);
    localStorage.setItem("model", model);
  }

  function loadCredentials() {
    apiType = localStorage.getItem("apiType") || "SERVER";
    apiKey = localStorage.getItem("apiKey") || "";
    model = localStorage.getItem("model") || "";
    storeCredentials();
  }

  function showError(message: string) {
    outputText = message;
  }

  $: characterCount = inputText.length;

  const languageOptions = Object.keys(Language)
    .filter((key) => isNaN(Number(key)) && key !== "UNSPECIFIED")
    .map((key) => key.charAt(0).toUpperCase() + key.slice(1).toLowerCase());

  const languageMaps: Record<string, Language> = {}
  Object.entries(Language).forEach(([key, value]) => {
    languageMaps[key.charAt(0).toUpperCase() + key.slice(1).toLowerCase()] = value as Language;
  })

  async function handleTranslation() {
    storeCredentials();
    if (inputText.length === 0) {
      return;
    }

    let langCode: Language;
    if (!languageMaps[outputLanguage]) {
      outputLanguage = "Spanish";
      outputText = "Language not supported";
    }
    langCode = languageMaps[outputLanguage];

    const request: TranslateRequest = TranslateRequest.fromJson({
      text: inputText,
      targetLanguage: langCode,
    });

    if (apiType !== "SERVER") {
      let provider: LLMOptions_Provider = LLMOptions_Provider.SERVER
      switch (apiType) {
        case "AISTUDIO":
          provider = LLMOptions_Provider.AISTUDIO
          if (model === "") {
            model = "gemini-1.5-flash-latest"
          }
          break;
        case "ANTHROPIC":
          provider = LLMOptions_Provider.ANTHROPIC
          if (model === "") {
            model = "claude-3-5-sonnet-20240620"
          }
          break;
        case "OPENAI":
          provider = LLMOptions_Provider.OPENAI
          if (model === "") {
            model = "gpt-4o-mini"
          }
          break;
        case "OPENROUTER":
          provider = LLMOptions_Provider.OPENROUTER
          if (model === "") {
            model = "google/gemini-flash-1.5"
          }
          break;
        default:
        showError("Invalid API Type");
      }
      storeCredentials();

      request.options = LLMOptions.fromJson({
        provider: provider,
        credentials: {
          apiKey: apiKey,
        },
        model: model,
      });
    }

    try {
      const response = await client.translate(request);
      console.log(response.toJson());

      if (response.success) {
        outputText = response.translatedText;
      } else {
        outputText = "Translation Failed, please try again later";
      }
    } catch (e) {
      showError("Error: " + e);
      return;
    }
  }

  function toggleReverse() {
    [inputText, outputText] = [outputText, inputText];
  }

  function copyToClipboard(text: string) {
    if (navigator.clipboard) {
      navigator.clipboard.writeText(text);
    } else {
      // Fallback for browsers that don't support clipboard API
      const textArea = document.createElement("textarea");
      textArea.hidden = true;
      textArea.style.display = "none";
      textArea.value = text;
      document.body.appendChild(textArea);
      textArea.select();
      document.execCommand("copy");
      document.body.removeChild(textArea);
    }
  }

  onMount(() => {
    loadCredentials();

    transport = createConnectTransport({
      baseUrl: "/",
      useBinaryFormat: true,
    });
    client = createPromiseClient(TranslationService, transport);
    client
      .healthz({})
      .then((res) => {
        console.log(res.toJson());
      })
      .catch((err) => {
        outputText =
          "Failed to connect to server, please check your connection and try again later.\n\n" +
          err;
        console.log(err);
      });
  });
</script>

<svelte:head>
  <title>LLM Translation</title>
  <link rel="stylesheet" href="/styles.css" />
  <meta name="viewport" content="width=device-width, initial-scale=1" />
</svelte:head>

<main>
  <div class="container">
    <h1>LLM Translation</h1>
    <div class="translation-container">
      <div class="translation-box">
        <h2>Input Text</h2>
        <textarea bind:value={inputText} placeholder="Enter text to translate"
        ></textarea>
        <p class="char-count">Character count: {characterCount}</p>
        <button on:click={() => copyToClipboard(inputText)} class="copy-btn"
          >Copy</button
        >
      </div>
      <div class="translation-box">
        <h2>
          <select bind:value={outputLanguage}>
            {#each languageOptions as language}
              <option value={language}>{language}</option>
            {/each}
          </select>
        </h2>
        <textarea
          bind:value={outputText}
          placeholder="Translation will appear here"
          readonly
        ></textarea>
        <button on:click={() => copyToClipboard(outputText)} class="copy-btn"
          >Copy</button
        >
      </div>
    </div>
    <div class="translation-actions">
      <button on:click={handleTranslation} class="translate-btn">
        Translate
      </button>
      <button on:click={toggleReverse} class="reverse-btn"> Swap </button>
    </div>

    <details>
      <summary>Advanced Settings</summary>
      <div class="advanced-settings">
        <div class="setting">
          <label for="apiType">API Type:</label>
          <select id="apiType" bind:value={apiType}>
            <option value="SERVER">Server</option>
            <option value="AISTUDIO">AI Studio</option>
            <option value="OPENAI">OpenAI</option>
            <option value="ANTHROPIC">Anthropic</option>
            <option value="OPENROUTER">OpenRouter</option>
          </select>
        </div>
        <div class="setting">
          <label for="apiKey">API Key:</label>
          <input
            type="password"
            id="apiKey"
            bind:value={apiKey}
            placeholder="Enter API Key"
          />
        </div>
        <div class="setting">
          <label for="model">Model:</label>
          <input
            type="text"
            id="model"
            bind:value={model}
            placeholder="Enter Model Name"
          />
        </div>
        <button class="reset-btn" on:click={resetAdvancedSettings}>Reset</button
        >
      </div>
    </details>
  </div>
</main>
