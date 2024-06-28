<script lang="ts">
    import 'open-props/style'
	import "../app.css";

    import {
        Int,
        Float,
        Bool,
        String,
        Date,
        DateTime,
        Duration,
        Time,
        Object,
        Array,
        Enum,
    } from "./types";

    import {Ea} from './utils'
</script>

<h1>Canonical Model Specification 2024-04-01</h1>

<h2>Types</h2>

<ul>
    <li><Array /></li>
    <li><Bool /></li>
    <li><Date /></li>
    <li><DateTime /></li>
    <li><Duration /></li>
    <li><Enum /></li>
    <li><Float /></li>
    <li><Int /></li>
    <li><Object /></li>
    <li><String /></li>
    <li><Time /></li>
</ul>

<h2>Operations</h2>

<h2>Data</h2>

<h2>References</h2>
<ul>
    <li>
        <Ea url="https://json-schema.org" />
        <ul>
            <li><Ea url="https://json-schema.org/understanding-json-schema" /></li>
        </ul>
    </li>
    <li><Ea url="https://pkl-lang.org" /></li>
    <li><Ea url="https://graphql.org" /></li>
    <li><Ea url="https://avro.apache.org" /></li>
    <li><Ea url="https://datatracker.ietf.org/doc/html/rfc3339" /></li>
</ul>

<style>
    li {
        list-style: none;
        margin-block: 1rem;
        border-style: solid;
        border-width: 1px;
        border-color: lightblue;
        background-color: aliceblue;
        padding: .25rem;
    }
</style>
