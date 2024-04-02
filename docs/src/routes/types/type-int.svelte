<script lang="ts">
    import TypeItem from "./type-item.svelte";
    import { Min, Max, MultipleOf } from "../properties";
</script>

<TypeItem name="int">
    <svelte:fragment slot="props">
        <li>
            <Min />
        </li>
        <li>
            <Max />
        </li>
        <li>
            <MultipleOf />
        </li>
    </svelte:fragment>
</TypeItem>
