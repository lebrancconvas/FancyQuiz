<script lang="ts">
  import axios from 'axios';
  import Footer from '../../components/Footer.svelte';

  let userID = 1;
  let reportContent: string = '';

  function submitHandler() {
    let requestData = {
      userID: userID,
      content: reportContent
    };
    axios.post('http://localhost:8011/api/reports/', requestData)
      .then(response => {
        console.log(response);
        window.location.reload();
      })
      .catch(error => {
        console.error(error);
      });
  }
</script>

<svelte:head>
  <title>Fancy Quiz | Report Problem</title>
</svelte:head>

<main>
  <header>
    <h1>Report Problem</h1>
  </header>
  <section id="report-section">
    <div id="textarea-container">
      <textarea name="report" id="report" cols="30" rows="10" bind:value={reportContent}></textarea>
    </div>
    <div id="button-container">
      <button on:click={submitHandler}>Submit</button>
    </div>
  </section>
  <Footer />
</main>

<style>
  header {
    text-align: center;
  }

  #textarea-container {
    width: 80%;
    /* outline: 1px solid black; */
    margin: 0 auto;
  }

  textarea {
    width: 100%;
    height: 700px;
    border: 2px solid black;
    border-radius: 5px;
    font-size: 18px;
    padding: 5px;
    margin-left: auto;
    margin-right: auto;
  }

  #button-container {
    width: 0;
    margin: 0 auto;
  }

  button {
    border: none;
    border-radius: 5px;
    padding: 10px;
    font-size: 20px;
    background-color: greenyellow;
    cursor: pointer;
    margin-left: auto;
    margin-right: auto;
  }

  button:active {
    transform: scale(0.98);
  }
</style>
