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
import { CountryForm } from "./form";

import { CountryItem } from "../applogic/model";
export default function CreateCountry() {
  const country = {} as CountryItem;
  return (
    <div>{country && <CountryForm country={country} editmode="create" />}</div>
  );
}
