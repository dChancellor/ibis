import {writable} from 'svelte/store'
import { application } from '../../wailsjs/go/models'

export const activeSkill = writable<application.Skill>()

export const skills = writable<application.Skill[]>([])