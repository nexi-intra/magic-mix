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
import { RelationForm } from "./form";

import { RelationItem } from "../applogic/model";
export default function CreateRelation() {
  const relation = {} as RelationItem;
  return (
    <div>
      {relation && <RelationForm relation={relation} editmode="create" />}
    </div>
  );
}
