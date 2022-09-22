<script>
	// @ts-nocheck
	import { validateEmail } from '@utils/validation';
	import * as services from '@services/auth';
	import { user } from '@stores/user';
	import Input from './input.svelte';

	export let closePopup;
	let email = '',
		password = '',
		name = '';
	let isEmailErr, isPasswordErr, isNameErr, submitErr;

	const SIGN_IN = 'SIGN_IN';
	const SIGN_UP = 'SIGN_UP';
	let isLoading = false;

	let mode = SIGN_IN;

	function attemptToClosePopup(evt) {
		if (evt.target === evt.currentTarget) {
			closePopup();
		}
	}

	function toggleMode() {
		resetForm();
		mode = mode === SIGN_IN ? SIGN_UP : SIGN_IN;
	}

	function resetForm() {
		email = '';
		password = '';
		name = '';
		isEmailErr = false;
		isPasswordErr = false;
		isNameErr = false;
		submitErr = false;
	}

	async function signIn(evt) {
		evt.stopPropagation();
		isEmailErr = !validateEmail(email);
		isPasswordErr = password.length < 8;
		const formDataValidates = !isEmailErr && !isPasswordErr;

		if (formDataValidates) {
			isLoading = true;
			let data;
			// simulate api request
			try {
				data = await services.SignIn(email, password);
				closePopup();
				user.loginUser(data);
				isLoading = false;
			} catch (err) {
				// console.log('🚀 ~ file: login.svelte ~ line 57 ~ signIn ~ err', JSON.stringify(err));
				submitErr = err;
				isLoading = false;
			}
		}
	}

	async function signUp(evt) {
		evt.stopPropagation();
		isEmailErr = !validateEmail(email);
		isPasswordErr = password.length < 8;
		isNameErr = name.length < 2;

		const formDataValidates = !isEmailErr && !isPasswordErr && !isNameErr;

		if (formDataValidates) {
			isLoading = true;
			let data;
			// simulate api request
			try {
				data = await services.SignUp(email, password, name);
				user.loginUser(data);
				closePopup();
				isLoading = false;
			} catch (err) {
				submitErr = err;
				isLoading = false;
			}
			console.log('signup with : ', data);
		}
	}

	function handleInput(evt) {
		switch (evt.target.dataset.label) {
			case 'Name':
				name = evt.target.value;
				break;
			case 'Email':
				email = evt.target.value;
				break;
			case 'Password':
				password = evt.target.value;
				break;
		}
	}
</script>

<div
	class="fixed w-screen h-screen top-0 left-0 flex justify-center items-center bg-gray-500/[0.8]"
	on:click={attemptToClosePopup}
>
	<div
		class="w-full min-w-[300px] sm:w-full md:w-2/3 max-w-[500px] px-10 py-10 bg-white z-10 text-base rounded relative"
	>
		{#if isLoading}
			<div
				class="w-full h-full absolute left-0 top-0 bg-gray-500/[0.3] flex items-center justify-center"
			>
				<img src="./loaders/eclipse.svg" alt="loader" />
			</div>
		{/if}

		<img class="w-20" src="icons/yt-full.svg" alt="youtube logo" />
		{#if mode === SIGN_IN}
			<h2 class="text-left">Login to Youtube</h2>
		{:else}
			<h2 class="text-left bold">Welcome to Youtube</h2>
		{/if}

		{#if mode === SIGN_UP}
			<div class="my-5 w-full flex flex-col">
				<Input inpType="name" inpValue={name} isErr={isNameErr} {handleInput} />
			</div>
		{/if}

		<div class="my-5 w-full flex flex-col">
			<Input inpType="email" inpValue={email} isErr={isEmailErr} {handleInput} />
		</div>
		<div class="my-5 w-full flex flex-col">
			<Input inpType="password" inpValue={password} isErr={isPasswordErr} {handleInput} />
			{#if mode === SIGN_IN}
				<button class="self-end">Forgot password ?</button>
			{/if}
		</div>

		{#if mode === SIGN_IN}
			<button class="my-5 border-black border-2 border-solid px-5 py-2 rounded" on:click={signIn}
				>Login</button
			>

			<div class="flex mt-10 justify-center">
				<p>I'm a new user</p>
				<button
					class="text-rose-500 ml-2 border-0 border-b-2 border-rose-500 border-solid
				"
					on:click={toggleMode}>Sign up</button
				>
			</div>
		{:else}
			<button class="my-5 border-black  border-2 border-solid px-5 py-2 rounded" on:click={signUp}
				>SIGN UP</button
			>
			<div class="flex mt-10 justify-center">
				<p>Go back to</p>
				<button
					class="text-rose-500 ml-2 border-0 border-b-2 border-rose-500 border-solid
				"
					on:click={toggleMode}>Sign in</button
				>
			</div>
		{/if}

		{#if submitErr}
			<p class="text-rose-500">{submitErr}</p>
		{/if}
	</div>
</div>
