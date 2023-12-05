=== How To Use ToDoList EndPoint ===

Run The Server 

Using PostMan

GET
http://localhost:8080/tasks --> Send

POST
http://localhost:8080/tasks
Choose raw --> JSON
{
    "title": "masukkan_tugas_baru",
    "done": false
}
Send

UPDATE
http://localhost:8080/tasks/ID
Choose raw --> JSON
{
    "done": true
}
Send
NB: hanya dapat mengubah dari false ke true, sedangkan sebaliknya tidak bisa

DELETE
http://localhost:8080/tasks/ID
Send
