"use client";
/* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/
import { Button } from "@/components/ui/button";
/* spejderhagl */

export default function DeleteCompany(props: { id: number ,onDeleteConfirmed:()=>void}) {
    const { id,onDeleteConfirmed } = props;
return (
<div>
<Button variant="destructive" onClick={onDeleteConfirmed}>Delete</Button><Button variant="secondary">Cancel</Button>


</div>
);
}
    
