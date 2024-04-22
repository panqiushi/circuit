<script setup lang="ts">
import { SideMenu } from '#components';
</script>

<template>
    <div class="flex h-screen w-full">
        <div class="w-1/6">
            <SideMenu />
        </div>
        <div class="h-screen w-5/6 bg-slate-50">
            <h1>main right</h1>
        </div>
    </div>
</template>