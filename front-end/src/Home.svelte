<script>
import { onMount } from "svelte";
    let authURL
    async function getAuthURL() {
        await fetch("http://127.0.0.1:3000/api/url").then((response) => {
            return response.json();
         }).then((data) => {
                authURL = data.url;
        });
    }
    onMount(getAuthURL)
</script>
<a href={authURL}>gogo</a>