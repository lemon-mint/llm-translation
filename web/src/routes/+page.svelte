<script lang="ts">
  import { TranslationService } from "$lib/proto/api/v1beta1/translate_connect";
  import { createPromiseClient } from "@connectrpc/connect";
  import { createConnectTransport } from "@connectrpc/connect-web";
  import { onMount } from "svelte";
  import "./styles.css";

  let inputText = "";
  let outputText = "";
  let outputLanguage = "SPANISH";
  
  // Reactive statement for character count
  $: characterCount = inputText.length;

  enum Language {
    UNSPECIFIED = 0,
    ENGLISH = 1,
    SPANISH = 2,
    CHINESE = 3,
    KOREAN = 4,
    JAPANESE = 5,
    GERMAN = 6,
    RUSSIAN = 7,
    FRENCH = 8,
    DUTCH = 9,
    ITALIAN = 10,
    INDONESIAN = 11,
    PORTUGUESE = 12,
    TAIWANESE = 13,
  }
  const languageOptions = Object.keys(Language).filter(
    (key) => isNaN(Number(key)) && key !== "UNSPECIFIED"
  );

  onMount(async () => {
    const transport = createConnectTransport({
      baseUrl: "/",
      useBinaryFormat: true,
    });
    const client = createPromiseClient(TranslationService, transport);
    const rpcStatus = await client.healthz({});
    console.log(rpcStatus.toJson());
  });

  async function handleTranslation() {
    // Translation logic here
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
        <textarea bind:value={inputText} placeholder="Enter text to translate"></textarea>
        <p class="char-count">Character count: {characterCount}</p>
        <button on:click={() => copyToClipboard(inputText)} class="copy-btn">Copy</button>
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
        <button on:click={() => copyToClipboard(outputText)} class="copy-btn">Copy</button>
      </div>
    </div>
    <div class="translation-actions">
      <button on:click={handleTranslation} class="translate-btn">
        Translate
      </button>
      <button on:click={toggleReverse} class="reverse-btn">
        Swap
      </button>
    </div>
  </div>
</main>
