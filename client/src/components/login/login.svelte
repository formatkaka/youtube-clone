<script>
	// @ts-nocheck
	import { validateEmail } from '@utils/validation';
	import * as services from '@services/auth';
	import { user } from '@stores/user';

	export let closePopup;
	let email = '',
		password = '',
		name = '';
	let isEmailErr, isPasswordErr, isNameErr, submitErr;
	let nameErr = 'Name should contain minimum 2 characters';

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
				<label class="text-left" for="name-input">Name</label>
				<input
					id="name-input"
					class=" border-black border-2 border-solid px-5 py-2 rounded"
					type="text"
					placeholder="Enter your name"
					required
					bind:value={name}
				/>
			</div>
		{/if}

		<div class="my-5 w-full flex flex-col">
			<label class="text-left" for="email-input">Email</label>
			<input
				required
				type="text"
				id="email-input"
				bind:value={email}
				placeholder="Enter your email"
				class={`border-black border-2 border-solid px-5 py-2 rounded ${
					isEmailErr ? 'border-rose-800	' : ''
				}`}
			/>
			{#if isEmailErr}
				<p class="text-left text-sm">Add proper email</p>
			{/if}
		</div>
		<div class="my-5 w-full flex flex-col">
			<label class="text-left" for="password-input">Password</label>
			<input
				required
				type="password"
				id="password-input"
				bind:value={password}
				placeholder="Enter your password"
				class={`border-black border-2 border-solid px-5 py-2 rounded ${
					isPasswordErr ? 'border-rose-800	' : ''
				}`}
			/>
			{#if isPasswordErr}
				<p class="text-left text-sm">Password should be minimum 8 characters</p>
			{/if}
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
