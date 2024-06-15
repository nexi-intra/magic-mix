/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
"use client";
import { useService } from "@/app/koksmat/useservice";
import { useState } from "react";
import { GroupForm } from "./form";

import { GroupItem } from "../applogic/model";
export default function CreateGroup() {
  const group = {} as GroupItem;
  return <div>{group && <GroupForm group={group} editmode="create" />}</div>;
}
