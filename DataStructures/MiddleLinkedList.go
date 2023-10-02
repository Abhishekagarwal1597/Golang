/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
    var len int
    var updatedList, lenPointer *ListNode
    lenPointer = head

    // First Find the length of the linked list
    for lenPointer != nil{
        len++
        lenPointer = lenPointer.Next
    }
    
    //Iterate through the half of the length
    for i:=0; i<(len/2); i++{
        head = head.Next
    }
    
    //return the pointer
    updatedList = head
    return updatedList

}
