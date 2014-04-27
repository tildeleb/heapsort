// Copyright © 2014 Lawrence E. Bakst. All rights reserved.
// A simple HeapSort algorithm.
var HeapSort = (function () {
    function siftup(v, ni, ri, compare) {
        var t;
        while (ri < Math.floor((ni + 1) / 2)) {
            var ci = ri * 2 + 1;                        // calculate left child node
            if (ci < ni && compare(v[ci], v[ci+1])) {   // follow the largest node left or right
                ci++;
            }
            if (compare(v[ci], v[ri]))  { // invariant holds
                break
            }
            // invariant doesn't hold, swap child and root and descent to the next level of tree
            t=v[ri], v[ri]=v[ci], v[ci]=t;
            ri = ci;
        }
    }

    return function (v, compare) {
        var n = v.length -1, t;

        //console.log("HeapSort: v.length=%o", v.length)
        // build a heap with max element at index 0 by building a binary tree and repeatedly satisfying the invaraint from the bottom up
        for (var i = Math.floor(n/2); i >= 0; i--) {
            siftup(v, n, i, compare)
        }

        while (n > 0) {
            t=v[0], v[0]=v[n], v[n]=t;
            n--;
            siftup(v, n, 0, compare);
         }
        return v;
    };
}());

/*
function main() {
    var compare = function(a, b) {
        if (a < b) {
            return true
        } else {
            return false
        }
    }

    var recs = [
        "AsfAGHM5om  00000000000000000000000000000000  0000222200002222000022220000222200002222000000001111",
        "~sHd0jDv6X  00000000000000000000000000000001  77779999444488885555CCCC777755555555BBBB666644446666",
        "uI^EYm8s=|  00000000000000000000000000000002  CCCCFFFF777799995555FFFF11112222999988884444DDDDFFFF",
        "Q)JN)R9z-L  00000000000000000000000000000003  FFFF111100000000000066668888BBBB33333333AAAA1111CCCC",
        "o4FoBkqERn  00000000000000000000000000000004  7777AAAABBBBBBBB22224444444499995555BBBB11118888DDDD",
        "*}-Wz1;TD-  00000000000000000000000000000005  AAAA88883333BBBB888888884444777722227777999900002222",
        "0fssx}~[oB  00000000000000000000000000000006  FFFF999977774444AAAA7777EEEEDDDDAAAAAAAA99998888BBBB",
        "mz4VCN@a#\"  00000000000000000000000000000007  DDDDBBBB1111FFFF2222DDDDFFFFBBBBFFFF6666444477778888",
        "my+=5r7(N|  00000000000000000000000000000008  22226666CCCC66662222FFFF0000EEEE11118888444455559999",
        "5H\\z%qt{%  00000000000000000000000000000009  0000AAAA8888FFFF0000888800000000222255551111FFFFEEEE"
    ]

    var nums = [60, 94, 66, 44, 43, 68, 7, 16, 10, 30, 52, 81, 22, 38, 32];
    console.log(HeapSort(nums, compare));
    console.log(HeapSort(recs, compare));
}

main();
*/