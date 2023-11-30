<script lang="ts">
  import axios from 'axios';
  // import { LOCAL_BASE_URL } from '$env/static/public';
  import Footer from '../../components/Footer.svelte';

  let username: string = "";
  let password: string = "";
  let confirmPassword: string = "";
  let displayName: string = "";

  function register() {
    if (password.trim() !== confirmPassword.trim()) {
      alert("Passwords do not match!");
      return;
    }

    if(displayName.trim() === "") {
      displayName = username;
    }

    const registerForm = {
      username: username.trim(),
      password: password.trim(),
      display_name: displayName.trim()
    };
    console.log(registerForm);

    axios.post(`http://locahost:8011/api/users`, registerForm)
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
  <title>Fancy Quiz | Register</title>
</svelte:head>

<main>
  <header>
    <h1>Register</h1>
  </header>
  <section id="register-section">
    <form id="register-form">
      <div>
        <label for="username">Username *</label>
        <input type="text" name="username" id="username" bind:value={username} required>
      </div>
      <div>
        <label for="password">Password *</label>
        <input type="password" name="password" id="password" bind:value={password} required>
      </div>
      <div>
        <label for="confirm-password">Confirm Password *</label>
        <input type="password" name="confirm-password" id="confirm-password" bind:value={confirmPassword} required>
      </div>
      <div>
        <label for="display-name">Display Name</label>
        <input type="text" name="display-name" id="display-name" bind:value={displayName}>
      </div>
      <div>
        <button type="submit" on:click={register}>Register</button>
      </div>
    </form>
  </section>
  <Footer />
</main>

<style>
  header {
    text-align: center;
  }
</style>
