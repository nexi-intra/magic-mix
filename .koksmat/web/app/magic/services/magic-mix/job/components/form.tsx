    /* 
File have been automatically created. To prevent the file from getting overwritten
set the Front Matter property ´keep´ to ´true´ syntax for the code snippet
---
keep: false
---
*/ 
"use client";
import { useState,useEffect } from "react";
import {JobItem} from "../applogic/model";
import {JobSchema} from "../applogic/model";
import { zodResolver } from "@hookform/resolvers/zod"
import { useForm } from "react-hook-form"
import { z } from "zod"
import { Button } from "@/components/ui/button";
import { toast } from "@/components/ui/use-toast"
import {
    Form,
    FormControl,
    FormDescription,
    FormField,
    FormItem,
    FormLabel,
    FormMessage,
  } from "@/components/ui/form"
  import { Input } from "@/components/ui/input"
/* marsbar */



export function JobForm(props : {job: JobItem,editmode:"create"|"update"}) {
    const {job,editmode} = props;
    function onSubmit(data: z.infer<typeof JobSchema>) {
        toast({
          title: "You submitted the following values:",
          description: (
            <pre className="mt-2 w-[340px] rounded-md bg-slate-950 p-4">
              <code className="text-white">{JSON.stringify(data, null, 2)}</code>
            </pre>
          ),
        })
      }
    const form = useForm<z.infer<typeof JobSchema>>({
        resolver: zodResolver(JobSchema)
      })

      useEffect(() => {
        form.reset(job);
      }, [job]);
    return (
      <div>
      <Form {...form}>
      <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-8">
       
          
    {job && <div>
        {/* string */}<FormField
 control={form.control}
 name="name"
 render={({ field }) => (
   <FormItem>
     <FormLabel>Name</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* string */}<FormField
 control={form.control}
 name="description"
 render={({ field }) => (
   <FormItem>
     <FormLabel>Description</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* string */}<FormField
 control={form.control}
 name="status"
 render={({ field }) => (
   <FormItem>
     <FormLabel>status</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* datetime */}<FormField
 control={form.control}
 name="startat"
 render={({ field }) => (
   <FormItem>
     <FormLabel>startat</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* datetime */}<FormField
 control={form.control}
 name="startedAt"
 render={({ field }) => (
   <FormItem>
     <FormLabel>startedAt</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* datetime */}<FormField
 control={form.control}
 name="completedAt"
 render={({ field }) => (
   <FormItem>
     <FormLabel>completedAt</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* int */}<FormField
 control={form.control}
 name="maxduration"
 render={({ field }) => (
   <FormItem>
     <FormLabel>maxduration</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* string */}<FormField
 control={form.control}
 name="script"
 render={({ field }) => (
   <FormItem>
     <FormLabel>script</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>
    {/* json */}<FormField
 control={form.control}
 name="data"
 render={({ field }) => (
   <FormItem>
     <FormLabel>data</FormLabel>
     <FormControl>
       <Input placeholder="" {...field} />

     </FormControl>
     <FormDescription>
       
     </FormDescription>
     <FormMessage />
   </FormItem>
 )}
/>

    <div>
   
    </div>
    </div>}


      <Button  type="submit">{editmode === "create"?"Create":"Update"}</Button>
      </form>
     </Form>
      </div>
    );
  }
      
