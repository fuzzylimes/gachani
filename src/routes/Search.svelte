<script>
    import { onMount } from "svelte";
    import Spinner from "../components/Spinner.svelte";

    import { inview } from "svelte-inview";
    import Card from "../components/Card.svelte";

    let next = "";
    let hasNext = false;
    let results = [];
    let newResults = [];
    let isLoading = false;

    $: results = [...results, ...newResults];

    const fetchData = async () => {
        isLoading = true;
        const q = document.getElementById("query").value;
        const res = await fetch(`/api/search?q=${q}&${next}`);
        const resJson = await res.json();
        newResults = resJson.data.map((n) => n.node);
        if (resJson?.paging?.next) {
            const nextLink = resJson.paging.next;
            next = nextLink.split("?").slice(1).join("?");
            hasNext = true;
        } else {
            hasNext = false;
        }
        isLoading = false;
    };

    const handleChange = (e) => {
        if (e.detail.inView && hasNext) fetchData();
    };

    onMount(() => {
        document
            .getElementById("query")
            .addEventListener("keypress", function (e) {
                if (e.key === "Enter") {
                    fetchData();
                }
            });
    });
</script>

<main>
    <div class="container px-5">
        <div
            class="columns has-text-centered is-centered is-multiline is-variable is-10-mobile is-two-thirds-tablet is-half-desktop is-one-third-widescreen"
        >
            <div class="column is-full">
                <div class="section mt-6">
                    <h1 class="title is-1">Search For Anime</h1>
                    <div
                        class="columns is-centered is-10-mobile is-half-tablet mt-2"
                    >
                        <div class="column is-6">
                            <input
                                class="input"
                                id="query"
                                type="string"
                                placeholder="Search Anime by Name"
                            />
                        </div>
                        <div class="column is-narrow">
                            <button
                                class="button is-link"
                                type="button"
                                disabled={isLoading}
                                on:click={fetchData}>Search!</button
                            >
                        </div>
                    </div>
                </div>
            </div>

            {#each results as result}
                <div class="column is-5-tablet is-4-desktop">
                    <Card anime={result} />
                </div>
            {/each}
            <!-- This component is place RIGHT BELOW the last post fetched. handleChange will be triggered
            as soon as it appears on frame. We'll use this info to know when to fetch more data  -->
            <div use:inview={{}} on:change={handleChange} />
            {#if isLoading}
                <div class="column is-full">
                    <Spinner />
                </div>
            {/if}
        </div>
    </div>
</main>

<style>
</style>
