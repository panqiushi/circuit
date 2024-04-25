<script setup lang="ts">
import { object, string, type InferType } from 'yup'
import type { FormSubmitEvent } from '#ui/types'

import { definePageMeta, useAuth } from '#imports'

// const { signIn, token, data, status, lastRefreshedAt } = useAuth()

const router = useRouter()
const gotoDashboard = () => router.push('/dashboard')
const gotoSignup = () => router.push('/signup')


// const cookie = useCookie("Set-Cookie")

const schema = object({
    email: string().email('Invalid email').required('Required'),
    password: string()
        .min(8, 'Must be at least 8 characters')
        .required('Required')
})

type Schema = InferType<typeof schema>

const state = reactive({
    email: undefined,
    password: undefined
})

definePageMeta({
    auth: {
        unauthenticatedOnly: true,
        navigateAuthenticatedTo: '/dashboard'
    }
})

const { signIn } = useAuth()
async function signInWithCredentials(event: FormSubmitEvent<Schema>) {
    try {
        await signIn(event.data, {callbackUrl: '/dashboard', external: false})
        console.log('Signed in')
        // gotoDashboard()
    } catch (error) {
        console.error(error)
    }
}

</script>

<template>
    <div class="flex justify-center items-center h-screen w-full">
        <UCard class="w-1/3">
            <h3 class="text-2xl font-semibold text-center mb-4">Login</h3>
            <span class="text-center block mb-4">
                Don't have an account? <a @click="gotoSignup" class="text-blue-500 cursor-pointer">Signup</a></span>
            <UForm :schema="schema" :state="state" class="space-y-4" @submit="signInWithCredentials">
                <UFormGroup label="Email" name="email">
                    <UInput v-model="state.email" />
                </UFormGroup>

                <UFormGroup label="Password" name="password">
                    <UInput v-model="state.password" type="password" />
                </UFormGroup>

                <UButton type="submit">
                    Submit
                </UButton>
            </UForm>
        </UCard>
    </div>
</template>