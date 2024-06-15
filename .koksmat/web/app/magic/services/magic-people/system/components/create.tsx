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
import { SystemForm } from "./form";

import { SystemItem } from "../applogic/model";
export default function CreateSystem() {
  const system = {} as SystemItem;
  return (
    <div>{system && <SystemForm system={system} editmode="create" />}</div>
  );
}
