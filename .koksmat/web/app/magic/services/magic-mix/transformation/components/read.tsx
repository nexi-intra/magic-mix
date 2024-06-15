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
import {TransformationItem} from "../applogic/model";


/* yankiebar */

export default function ReadTransformation(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TransformationItem>(
      "magic-mix.transformation",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const transformation = readResult.data;
    return (
      <div>
          
    {transformation && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{transformation.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{transformation.description}</div>
    </div>    <div>
        <div className="font-bold" >input</div>
        <div>{transformation.input_id}</div>
    </div>    <div>
        <div className="font-bold" >output</div>
        <div>{transformation.output_id}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{transformation.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{transformation.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{transformation.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{transformation.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{transformation.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
