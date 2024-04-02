<script lang="ts">
    import TypeItem from "./type-item.svelte";
    import { Min, Max } from "../properties";
</script>

<TypeItem name="string">
    <svelte:fragment slot="props">
        <li>
            <Min />
        </li>
        <li>
            <Max />
        </li>
    </svelte:fragment>
</TypeItem>
