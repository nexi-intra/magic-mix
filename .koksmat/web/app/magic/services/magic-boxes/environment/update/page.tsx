"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import EnvironmentUpdate from "@/services/magic-mix/endpoints/environment/update/webcomponent";
import { VsCodeEdittoolbar } from "@/app/magic/components/VsCodeEdittoolbar";

export default function TestEnvironmentUpdate() {
  return (
    <div>
      <VsCodeEdittoolbar
        filePath={"app/magic/services/magic-mix/environment/update/page.tsx"}
      />
      <EnvironmentUpdate />
    </div>
  );
}
