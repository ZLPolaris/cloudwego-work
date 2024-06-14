package constants

const OldIdlContent = "namespace go student_demo\n\n//--------------------request & response--------------\nstruct College {\n    1: required string name(go.tag = 'json:\"name\"'),\n    2: string address(go.tag = 'json:\"address\"'),\n}\n\nstruct Student {\n    1: required i32 id(api.body='id'),\n    2: required string name(api.body='name'),\n    3: required College college(api.body='college'),\n    4: optional list<string> email(api.body='email'),\n}\n\nstruct RegisterResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n}\n\nstruct QueryReq {\n    1: required i32 id(api.query='id')\n}\n\n//----------------------service-------------------\nservice StudentService {\n    RegisterResp Register(1: Student student)(api.post = '/add-student-info')\n    Student Query(1: QueryReq req)(api.get = '/query')\n}"
const NewIdlContent = "namespace go student_demo\n\n//--------------------request & response--------------\nstruct College {\n    1: required string name(go.tag = 'json:\"name\"'),\n    2: string address(go.tag = 'json:\"address\"'),\n}\n\nstruct Student {\n    1: required i32 id(api.body='id'),\n    2: required string name(api.body='name'),\n    3: required College college(api.body='college'),\n    4: optional list<string> email(api.body='email'),\n    5: string gender(api.body='gender'),\n}\n\nstruct RegisterResp {\n    1: bool success(api.body='success'),\n    2: string message(api.body='message'),\n}\n\nstruct QueryReq {\n    1: required i32 id(api.query='id')\n}\n\n//----------------------service-------------------\nservice StudentService {\n    RegisterResp Register(1: Student student)(api.post = '/add-student-info')\n    Student Query(1: QueryReq req)(api.get = '/query')\n}"