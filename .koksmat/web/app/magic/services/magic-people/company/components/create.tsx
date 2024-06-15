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
import { CompanyForm } from "./form";

import { CompanyItem } from "../applogic/model";
export default function CreateCompany() {
  const company = {} as CompanyItem;
  return (
    <div>{company && <CompanyForm company={company} editmode="create" />}</div>
  );
}
