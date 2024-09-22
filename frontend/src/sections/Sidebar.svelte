<script lang="ts">
    import {AddSkill, GetSkills} from '../../wailsjs/go/main/App'
    import Button from "../components/Button.svelte";

    import {constants} from '../stores/backendConstants'
    import {activeSkill, skills} from '../stores/skillStore'

    import {main} from '../../wailsjs/go/models'

    let dialog: HTMLDialogElement;
    let newSkillName: string | null;
    let errorText: string;

    function showAddSkillModal(): void {
        dialog.showModal();
    }
    async function addSkill(): Promise<void> {
        try {
            await AddSkill(newSkillName)
            dialog.close();
            newSkillName = null;
            getSkills();
        } catch(error) {
            if(error = $constants.SkillExistsMessage) {
                errorText = 'This skill name already exists';
            } else {
                errorText = 'There was an error saving your skill';
                console.error(errorText + ':', error)
            }
        }
    }
    async function getSkills(): Promise<void> {
        let res = await GetSkills();
        if(!res) {
            console.warn('Hmm, no skills were fetched...')
        }
        skills.set(res)
    }
    function activateSkill(skill: main.Skill): void {
        activeSkill.set(skill)
    }
</script>

<svelte:window on:load={getSkills}/>

<section>
    <div class='skills-list'> 
        {#each $skills as skill}
            <div class={$activeSkill?.Name === skill.Name && 'active'}>
            <Button text={skill.Name} onClick={() => activateSkill(skill)} />
            </div>

        {/each}
    </div>

    <Button text={'+'} onClick={showAddSkillModal}/>
    <dialog bind:this={dialog}>
        {#if errorText} <p>{errorText}</p> {/if}
        <input name="New Skill" bind:value={newSkillName} placeholder="Name of the new skill"/>
        <Button text={'Add skill'} onClick={() => addSkill()} />
        <Button text={'Close'} onClick={() => dialog.close()} />
    </dialog>
</section>

<style>
    section {
        padding: 2rem;
    }
    h1{
        width: min-content;
        margin-inline: 2rem;
    }
    .divider {
        height: 2px;
        background-color:white;
        width: 100%;
    }
    .skills-list{
        display:flex;
        flex-flow: column;
    }
    .active {
        background-color: pink;
    }
</style>