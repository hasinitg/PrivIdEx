/** @file
 *****************************************************************************
 * @author     This file is part of libsnark, developed by SCIPR Lab
 *             and contributors (see AUTHORS).
 * @copyright  MIT license (see LICENSE file)
 *****************************************************************************/

#ifndef SIMPLE_EXAMPLE_HPP_
#define SIMPLE_EXAMPLE_HPP_

#include <libff/common/default_types/ec_pp.hpp>
//#include <libsnark/gadgetlib2/pp.hpp>
#include "r1cs_examples.hpp"

namespace libsnark {

    r1cs_example<libff::Fr<libff::default_ec_pp>> gen_r1cs_example_from_gadgetlib2_protoboard(const size_t size);

} // libsnark

#include "simple_example.cpp"

#endif // SIMPLE_EXAMPLE_HPP_
