<script>
import Router from "svelte-spa-router";
import Sort from "./Sort.svelte";
import Home from "./Home.svelte";

	const routes = {
		"/": Home,
		"/sort": Sort ,
	}
</script>
<Router {routes} />

