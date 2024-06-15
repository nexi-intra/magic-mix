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
import {TransformerItem} from "../applogic/model";


/* yankiebar */

export default function ReadTransformer(props: { id: number }) {
    const { id } = props;
    const [transactionId, settransactionId] = useState(0);
    const readResult = useService<TransformerItem>(
      "magic-mix.transformer",
      ["read", id?.toString()],
      "",
      6000,
      transactionId.toString()
    );
    const transformer = readResult.data;
    return (
      <div>
          
    {transformer && <div>
        <div>
        <div className="font-bold" >Name</div>
        <div>{transformer.name}</div>
    </div>    <div>
        <div className="font-bold" >Description</div>
        <div>{transformer.description}</div>
    </div>    <div>
        <div className="font-bold" >code</div>
        <div>{transformer.code}</div>
    </div>
    <div>
        <div>
        <div className="font-bold" >id</div>
        <div>{transformer.id}</div>
    </div>
        <div>
        <div className="font-bold" >created_at</div>
        <div>{transformer.created_at}</div>
    </div>
        <div>
        <div className="font-bold" >created_by</div>
        <div>{transformer.created_by}</div>
    </div>
        <div>
        <div className="font-bold" >updated_at</div>
        <div>{transformer.updated_at}</div>
    </div>
        <div>
        <div className="font-bold" >updated_by</div>
        <div>{transformer.updated_by}</div>
    </div>
    </div>
    </div>}


     
      </div>
    );
  }
      
