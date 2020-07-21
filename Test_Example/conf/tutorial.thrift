namespace go tutorial

const string BS_DEPART = "DEPARTMENT"
const string BS_Data = "QUANLY"

struct SharedStruct {
  1: i32 key
  2: string value
}

service SharedService {
  SharedStruct getStruct(1: i32 key)
}
struct User{
	1: string User_id,
	2: string  Username,
	3: string Department_id,
}

struct Department{
    1: string Department_id,
    2: string Name,
}

service UserService{
    void ping(),
    list<User> GetAll(),
    string AddUser(1: User u),
    string Put(1: User u),
    string Delete(1: string id),
}

service DepartmentService{
    void ping(),
    list<Department> GetAll(),
    string AddDepartment(1: Department d),
    string Put(1: Department d),
    string Delete(1: string id),
    list<Department> GetUsers(1: string deid),
    list<Department> GetUsersSame(1: string userid),
}
