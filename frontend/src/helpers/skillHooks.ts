import { activeSkill, skills } from "@stores/skillStore";
import { AddSkill, GetSkills } from "@wails/main/App";
import { throwError } from "./errorHandler";

export async function getSkills(): Promise<void> {
  let res = await GetSkills();
  if (!res) {
    throwError("Failed to fetch skills");
  }
  skills.set(res);
  // TODO get rid of this (just for dev purposes to save mouse clicks)
  activeSkill.set(res[0]);
}

export async function addSkill(
  newSkillName: string,
  svgFileContent: string
): Promise<void> {
  let errorText: string;
  if (!newSkillName) {
    errorText = "Please enter a skill name";
    return;
  }
  try {
    await AddSkill(newSkillName, svgFileContent);
    getSkills();
  } catch (error) {
    throwError(`${errorText}: ${newSkillName}, ${svgFileContent}`, error);
    return error;
  }
}
