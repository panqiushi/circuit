<script setup lang="ts">
import { ref } from 'vue'

const router = useRouter()

onMounted(() => {
    if (localStorage.getItem('token')) {
        router.push('/dashboard')
    } else {
        router.push('/login')
    }
})
</script>

<template>
    <div>
        <h1>f root</h1>
    </div>
</template>