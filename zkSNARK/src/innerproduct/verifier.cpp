//
// Created by Hasini Gunasinghe on 2/8/18.
//

#include <libsnark/zk_proof_systems/ppzksnark/r1cs_ppzksnark/r1cs_ppzksnark.hpp>

namespace libsnark {
    template <typename ppT>
    const bool verifyProofOnR1CS(const r1cs_ppzksnark_verification_key<ppT> &vk,
                                 const r1cs_ppzksnark_primary_input<ppT> &primary_input,
                                 const r1cs_ppzksnark_proof<ppT> &proof){

        return r1cs_ppzksnark_verifier_strong_IC(vk, primary_input, proof);

    }
}