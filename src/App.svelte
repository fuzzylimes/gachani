<script>
	import Router from "svelte-spa-router";
	import Navbar from "./components/Navbar.svelte";
	import routes from "./routes";
</script>

<Navbar />

<Router {routes} />

<style>
	:global(html) {
		/* position: fixed; */
		/* background-color: #e9e9e979; */
		background-size: cover;
		/* background-repeat: repeat; */
	}
</style>
