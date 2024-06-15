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
import { PersonForm } from "./form";

import { PersonItem } from "../applogic/model";
export default function CreatePerson() {
  const person = {} as PersonItem;
  return (
    <div>{person && <PersonForm person={person} editmode="create" />}</div>
  );
}
